package cmd

import (
	"arnoldj-devops/ghvars/utils"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Struct to unmarshal JSON
var response struct {
	Environments []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"environments"`
}

func getVars() {
	ghEnvironments, err := ListEnvironments()
	if err != nil {
		fmt.Println("Error fetching environments:", err)
		return
	}
	// fmt.Println("Environment names are:", ghEnvironments)

	for _, env := range ghEnvironments {
		err := listVariables(env)
		if err != nil {
			fmt.Printf("Error fetching variables for environment %s: %v\n", env, err)
		}
	}
}

func listVariables(environmentName string) error {
	// Run the gh variable list command for the specified environment
	variableListCommand := exec.Command("gh", "variable", "list", "-e", environmentName)
	output, err := variableListCommand.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running gh variable list: %v", err)
	}

	// Remove timestamp from each line in the output
	lines := strings.Split(string(output), "\n")
	var filteredOutput []string
	for _, line := range lines {
		// Assuming timestamp is in the format "YYYY-MM-DDTHH:MM:SSZ"
		lineWithoutTimestamp := regexp.MustCompile(`\s+\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`).ReplaceAllString(line, "")
		filteredOutput = append(filteredOutput, lineWithoutTimestamp)
	}

	// Create or open the file for writing
	fileName := fmt.Sprintf("%s.env", environmentName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", fileName, err)
	}
	defer file.Close()

	// Write variables and values to the file in the format VARIABLE=VALUE
	for _, line := range filteredOutput {
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			fmt.Fprintf(file, "%s=%s\n", parts[0], parts[1])
		}
	}

	fmt.Printf("Variables for environment %s written to file %s\n", environmentName, fileName)
	return nil
}

func ListEnvironments() ([]string, error) {

	listEnvironmentsCommand := "gh api repos/{owner}/{repo}/environments"
	listEnvironmentsExec, err := utils.ExecuteBashCommand(listEnvironmentsCommand)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Unmarshal JSON into the struct
	err = json.Unmarshal([]byte(listEnvironmentsExec), &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	var environmentNames []string
	for _, env := range response.Environments {
		environmentNames = append(environmentNames, env.Name)
	}

	return environmentNames, nil
}
