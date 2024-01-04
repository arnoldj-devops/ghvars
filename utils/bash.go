package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteBashCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v", err)
	}
	return string(output), nil
}
