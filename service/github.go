package service

import (
	"fmt"
	"mpbench/config"
	"mpbench/utils"
	"os"
	"os/exec"
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
