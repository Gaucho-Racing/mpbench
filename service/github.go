package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mpbench/config"
	"mpbench/utils"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func CheckoutCommit(commit string) (string, error) {
	utils.SugarLogger.Infof("Checking out commit %s", commit)

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("mpbench-build-%s-*", commit))
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Clone the repository
	cmd := exec.Command("git", "clone", config.MapacheRepo, tmpDir)
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

func CreateCheckRun() (string, error) {
	type CheckRunPayload struct {
		Name       string `json:"name"`
		HeadSHA    string `json:"head_sha"`
		Status     string `json:"status"`
		ExternalID string `json:"external_id"`
		StartedAt  string `json:"started_at"`
		Output     struct {
			Title   string `json:"title"`
			Summary string `json:"summary"`
			Text    string `json:"text"`
		} `json:"output"`
	}

	payload := CheckRunPayload{
		Name:       "mpbench",
		HeadSHA:    commit,
		Status:     "in_progress",
		ExternalID: "1",
		StartedAt:  time.Now().Format(time.RFC3339),
		Output: struct {
			Title   string `json:"title"`
			Summary string `json:"summary"`
			Text    string `json:"text"`
		}{
			Title:   "MPBench Performance Report",
			Summary: "",
			Text:    "",
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal check run payload: %w", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/check-runs", config.MapacheRepo)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+config.GithubToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

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

	checkRunID, ok := result["id"].(float64)
	if !ok {
		return "", fmt.Errorf("failed to get check run ID from response")
	}

	return fmt.Sprintf("%d", int(checkRunID)), nil
}
