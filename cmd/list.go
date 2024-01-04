package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func promptGithubEnvironments() string {
	ghEnvironments, err := ListEnvironments()

	prompt := promptui.Select{
		Label: "Select the Github Environment",
		Items: ghEnvironments,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}
