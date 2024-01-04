package cmd

import (
	"arnoldj-devops/ghvars/utils"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
		fmt.Printf("Error running gh variable list: %v\nOutput:\n%s\n", err, output)
		return fmt.Errorf("error running gh variable list: %v", err)
	}

	// Create or open the file for writing
	fileName := fmt.Sprintf("%s.env", environmentName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", fileName, err)
	}
	defer file.Close()

	// Parse the tab-separated values
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line into parts: variable name, value, and timestamp
		parts := strings.Split(line, "\t")
		if len(parts) >= 3 {
			variableName := parts[0]
			value := parts[1] // Only consider the second part as the value
			// Write the variable and value to the file
			fmt.Fprintf(file, "%s=%s\n", variableName, value)
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
