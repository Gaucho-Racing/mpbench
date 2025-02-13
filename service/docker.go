package service

import (
	"fmt"
	"mpbench/utils"
	"os"
	"os/exec"
)

func BuildDockerImage(commit string, directory string, service string) string {
	utils.SugarLogger.Infof("Building Docker image for %s at commit %s", service, commit)
	cmd := exec.Command("docker", "build", "-t", fmt.Sprintf("gauchoracing/mp_%s:%s", service, commit), "--push", fmt.Sprintf("%s/%s", directory, service))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		utils.SugarLogger.Error("failed to build docker image", err)
		return ""
	}
	utils.SugarLogger.Infof("Successfully built and pushed Docker image for %s at commit %s", service, commit)
	if err := os.RemoveAll(directory); err != nil {
		utils.SugarLogger.Error("failed to remove temp directory", err)
	}
	return fmt.Sprintf("gauchoracing/mp_%s:%s", service, commit)
}
