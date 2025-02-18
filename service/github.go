package service

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"mpbench/config"
	"mpbench/utils"
	"net/http"
	"os"
	"os/exec"
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

func CreateCheckRun(commit string) (string, error) {
	token, err := GetGithubAppInstallationToken()
	if err != nil {
		return "", fmt.Errorf("failed to get installation token: %w", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/check-runs", config.MapacheRepo)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	type CheckRunPayload struct {
		Name        string `json:"name,omitempty"`
		HeadSHA     string `json:"head_sha,omitempty"`
		Status      string `json:"status,omitempty"`
		Conclusion  string `json:"conclusion,omitempty"`
		ExternalID  string `json:"external_id,omitempty"`
		DetailsURL  string `json:"details_url,omitempty"`
		StartedAt   string `json:"started_at,omitempty"`
		CompletedAt string `json:"completed_at,omitempty"`
		Output      struct {
			Title   string `json:"title,omitempty"`
			Summary string `json:"summary,omitempty"`
			Text    string `json:"text,omitempty"`
		} `json:"output,omitempty"`
	}

	payload := CheckRunPayload{
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
		return "", fmt.Errorf("failed to marshal check run payload: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to create check run: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to create check run: %s", string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	checkRunID, ok := result["id"].(int)
	if !ok {
		return "", fmt.Errorf("failed to get check run ID from response")
	}

	return fmt.Sprintf("%d", checkRunID), nil
}
