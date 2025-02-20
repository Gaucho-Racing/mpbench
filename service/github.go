package service

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"mpbench/config"
	"mpbench/model"
	"mpbench/utils"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CheckoutCommit(commit string) (string, error) {
	utils.SugarLogger.Infof("Checking out commit %s", commit)

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("mpbench-build-%s-*", commit))
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Clone the repository
	cmd := exec.Command("git", "clone", fmt.Sprintf("https://github.com/%s", config.MapacheRepo), tmpDir)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to clone repository: %w", err)
	}

	// Checkout specific tag/commit if provided
	if commit != "" {
		cmd = exec.Command("git", "-C", tmpDir, "checkout", commit)
		if err := cmd.Run(); err != nil {
			return "", fmt.Errorf("failed to checkout tag: %w", err)
		}
	}

	utils.SugarLogger.Infof("Successfully checked out commit %s", commit)
	utils.SugarLogger.Infof("Repository checked out to: %s", tmpDir)
	return tmpDir, nil
}

func GetGithubAppJWT() (string, error) {
	// Read private key from environment variable or file
	privateKeyPEM := os.Getenv("GITHUB_PRIVATE_KEY")
	if privateKeyPEM == "" {
		return "", fmt.Errorf("GITHUB_PRIVATE_KEY environment variable is not set")
	}

	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// Create the JWT claims
	now := time.Now()
	claims := jwt.MapClaims{
		"iat": now.Add(-60 * time.Second).Unix(), // Issued 60 seconds in the past
		"exp": now.Add(10 * time.Minute).Unix(),  // Expires in 10 minutes
		"iss": os.Getenv("GITHUB_APP_ID"),        // GitHub App ID
	}

	// Create token with claims and sign with RS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}

func GetGithubAppInstallationToken() (string, error) {
	jwt, err := GetGithubAppJWT()
	if err != nil {
		return "", fmt.Errorf("failed to get JWT: %w", err)
	}
	utils.SugarLogger.Infof("JWT: %s", jwt)

	type InstallationTokenResponse struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expires_at"`
	}
	url := fmt.Sprintf("https://api.github.com/app/installations/%s/access_tokens", config.GithubAppInstallationID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create installation token: %s", resp.Status)
	}

	var result InstallationTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Token, nil
}

func CreateCheckRun(commit string) (int, error) {
	token, err := GetGithubAppInstallationToken()
	if err != nil {
		return 0, fmt.Errorf("failed to get installation token: %w", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/check-runs", config.MapacheRepo)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	payload := model.CheckRunPayload{
		Name:       "mpbench / unit",
		HeadSHA:    commit,
		Status:     "queued",
		ExternalID: "1",
		Output: struct {
			Title   string `json:"title,omitempty"`
			Summary string `json:"summary,omitempty"`
			Text    string `json:"text,omitempty"`
		}{
			Title:   "MPBench Unit Tests",
			Summary: "queued",
			Text:    "",
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal check run payload: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to create check run: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("failed to create check run: %s", string(body))
	}

	var result model.CheckRunPayload
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	checkRunID := result.ID
	if checkRunID == 0 {
		return 0, fmt.Errorf("failed to get check run ID from response")
	}

	return checkRunID, nil
}

func UpdateCheckRun(checkRunID int, payload model.CheckRunPayload) error {
	token, err := GetGithubAppInstallationToken()
	if err != nil {
		return fmt.Errorf("failed to get installation token: %w", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/check-runs/%d", config.MapacheRepo, checkRunID)
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal check run payload: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to update check run: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update check run: %s", string(body))
	}

	return nil
}

func UpdateCheckRunInProgress(runID string) {
	run := GetRunByID(runID)
	if run.GithubCheckRunID != 0 {
		err := UpdateCheckRun(run.GithubCheckRunID, model.CheckRunPayload{
			Status:     "in_progress",
			ExternalID: run.ID,
			DetailsURL: fmt.Sprintf("https://mpbench.gauchoracing.com/runs/%s", run.ID),
			Output: struct {
				Title   string `json:"title,omitempty"`
				Summary string `json:"summary,omitempty"`
				Text    string `json:"text,omitempty"`
			}{
				Title:   fmt.Sprintf("MPBench %s Unit Tests", run.Service),
				Summary: fmt.Sprintf("Run ID: %s\nRunning tests...", run.ID),
				Text:    "",
			},
		})
		if err != nil {
			utils.SugarLogger.Error("Failed to update check run", err)
		}
	}
}

func GenerateCheckRunConclusion(runID string) {
	run := GetRunByID(runID)
	passed := make([]model.RunTest, 0)
	partial := make([]model.RunTest, 0)
	failed := make([]model.RunTest, 0)
	success := true
	for _, test := range run.RunTests {
		if test.Status == "passed" {
			passed = append(passed, test)
		} else {
			numPassed := 0
			for _, result := range test.RunTestResults {
				if result.Status == "passed" {
					numPassed++
				}
			}
			if numPassed == len(run.RunTests) {
				// all signals passed
				passed = append(passed, test)
			} else if numPassed > 0 {
				// some signals passed
				partial = append(partial, test)
			} else {
				// all signals failed
				failed = append(failed, test)
			}
		}
	}
	if len(partial) > 0 {
		success = false
	}

	allTests := run.RunTests
	sort.Slice(allTests, func(i, j int) bool {
		return allTests[i].Name < allTests[j].Name
	})
	sort.Slice(passed, func(i, j int) bool {
		return passed[i].Name < passed[j].Name
	})
	sort.Slice(partial, func(i, j int) bool {
		return partial[i].Name < partial[j].Name
	})
	sort.Slice(failed, func(i, j int) bool {
		return failed[i].Name < failed[j].Name
	})

	textBuffer := bytes.NewBufferString("")
	textBuffer.WriteString("# Summary\n\n")
	textBuffer.WriteString("| ID | Name | Status | Progress |\n")
	textBuffer.WriteString("|----|------|--------|----------|\n")
	for _, test := range allTests {
		numPassed := 0
		total := len(test.RunTestResults)
		for _, result := range test.RunTestResults {
			if result.Status == "passed" {
				numPassed++
			}
		}
		var status string
		if test.Status == "passed" {
			status = "✅ PASS"
		} else if numPassed == 0 {
			status = "❌ FAIL"
		} else {
			status = "⚠️ PARTIAL"
		}
		parts := strings.SplitN(test.Name, " ", 2) // Split on first space only
		textBuffer.WriteString(fmt.Sprintf("%s | %s | %s | %d/%d (%d%%)\n",
			parts[0], // hex ID (e.g., "0x003")
			parts[1], // actual name
			status,
			numPassed,
			total,
			(numPassed*100)/total))
	}

	if len(partial) > 0 {
		textBuffer.WriteString("\n# Partially Passed Tests\n\n")
		textBuffer.WriteString(RunTestsToResultString(partial))
	}
	if len(failed) > 0 {
		textBuffer.WriteString("\n# Failed Tests\n\n")
		textBuffer.WriteString(RunTestsToResultString(failed))
	}
	if len(passed) > 0 {
		textBuffer.WriteString("\n# Passed Tests\n\n")
		textBuffer.WriteString(RunTestsToResultString(passed))
	}

	if success {
		payload := model.CheckRunPayload{
			Name:       run.Name,
			Status:     "completed",
			Conclusion: "success",
			Output: struct {
				Title   string `json:"title,omitempty"`
				Summary string `json:"summary,omitempty"`
				Text    string `json:"text,omitempty"`
			}{
				Title:   fmt.Sprintf("MPBench %s Unit Tests", run.Service),
				Summary: fmt.Sprintf("Run ID: %s\n✅ %d tests passed\n⚠️ %d tests partially passed\n❌ %d tests failed\n\n**Note:** Failed tests are treated as unimplemented, allowing the check run as a whole to pass. Always double check failing tests for any unintentional misses.", run.ID, len(passed), len(partial), len(failed)),
				Text:    textBuffer.String(),
			},
		}
		UpdateCheckRun(run.GithubCheckRunID, payload)
	} else {
		payload := model.CheckRunPayload{
			Name:       run.Name,
			Status:     "completed",
			Conclusion: "failure",
			Output: struct {
				Title   string `json:"title,omitempty"`
				Summary string `json:"summary,omitempty"`
				Text    string `json:"text,omitempty"`
			}{
				Title:   fmt.Sprintf("MPBench %s Unit Tests", run.Service),
				Summary: fmt.Sprintf("Run ID: %s\n✅ %d tests passed\n⚠️ %d tests partially passed\n❌ %d tests failed\n\n**Note:** Failed tests are treated as unimplemented, allowing the check run as a whole to pass. Always double check failing tests for any unintentional misses.", run.ID, len(passed), len(partial), len(failed)),
				Text:    textBuffer.String(),
			},
		}
		UpdateCheckRun(run.GithubCheckRunID, payload)
	}
}

func RunTestsToResultString(run_test []model.RunTest) string {
	textBuffer := bytes.NewBufferString("")
	for _, test := range run_test {
		textBuffer.WriteString(fmt.Sprintf("### %s\n\n", test.Name))
		textBuffer.WriteString(fmt.Sprintf("Input Data: `%s`\n", test.Data))
		textBuffer.WriteString("\n| Signal Name | Decoded Value | Expected Value | Status |\n")
		textBuffer.WriteString("|------------|---------------|----------------|---------|\n")
		for _, result := range test.RunTestResults {
			if result.Status == "failed" {
				textBuffer.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					result.SignalName,
					result.Value,
					result.Expected,
					"❌"))
			} else {
				textBuffer.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					result.SignalName,
					result.Value,
					result.Expected,
					"✅"))
			}
		}
	}
	return textBuffer.String()
}
