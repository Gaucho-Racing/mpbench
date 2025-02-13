package service

import (
	"fmt"
	"mpbench/config"
	"mpbench/utils"
	"os"
	"os/exec"
)

func CheckoutCommit(commit string) string {
	utils.SugarLogger.Infof("Checking out commit %s", commit)

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("mpbench-build-%s-*", commit))
	if err != nil {
		utils.SugarLogger.Error("failed to create temp directory", err)
		return ""
	}

	// Clone the repository
	cmd := exec.Command("git", "clone", config.MapacheRepo, tmpDir)
	if err := cmd.Run(); err != nil {
		utils.SugarLogger.Error("failed to clone repository", err)
		return ""
	}

	// Checkout specific tag/commit if provided
	if commit != "" {
		cmd = exec.Command("git", "-C", tmpDir, "checkout", commit)
		if err := cmd.Run(); err != nil {
			utils.SugarLogger.Error("failed to checkout tag", err)
			return ""
		}
	}

	utils.SugarLogger.Infof("Successfully checked out commit %s", commit)
	utils.SugarLogger.Infof("Repository checked out to: %s", tmpDir)
	return tmpDir
}
