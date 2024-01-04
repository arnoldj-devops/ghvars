package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var ghEnvironments []string

func setVars(env string) {
	if env == "none" {
		ghEnvironments = strings.Fields(promptGithubEnvironments())
	} else if env == "all" {
		ghEnvironments, _ = ListEnvironments()
	} else {
		ghEnvironments = strings.Fields(env)
	}
	fmt.Printf("%s", ghEnvironments)
	for _, ghenv := range ghEnvironments {
		green := color.New(color.FgGreen)
		boldGreen := green.Add(color.Bold)
		boldGreen.Printf("Setting variables for environment: %s\n", ghenv)

		err := SetGitHubVariables(ghenv, setDryRun)
		if err != nil {
			fmt.Printf("Error setting variables for environment %s: %v\n", ghenv, err)
		}
	}
}

// SetGitHubVariables reads variables from the specified environment file and sets them as GitHub environment variables
func SetGitHubVariables(environmentName string, dryRun bool) error {
	if dryRun {
		fmt.Println("Dry Run - Changes to be made:")
		return nil
	}
	fileName := fmt.Sprintf("%s.env", environmentName)

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", fileName, err)
	}
	defer file.Close()

	// Read variables from the file and set them as GitHub environment variables
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			// Run the gh variable set command for each variable
			variableSetCommand := exec.Command("gh", "variable", "set", key, "--body", value, "-e", environmentName)
			// fmt.Printf("%s", variableSetCommand)
			output, err := variableSetCommand.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error running gh variable set for %s: %v\nOutput: %s", key, err, output)
			}

			fmt.Printf("GitHub environment variable %s set to %s\n", key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", fileName, err)
	}

	return nil
}
