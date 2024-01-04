package cmd

import (
	"arnoldj-devops/ghvars/utils"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

func doctor() {

	// Check if gh cli is available
	ghversioncommand := "gh version | awk '{print $3}' | tr -d '\n'"
	ghversionout, err := utils.ExecuteBashCommand(ghversioncommand)
	if strings.Contains(string(ghversionout), "command not found") {
		log.Fatal("Please reinstall gh cli, https://cli.github.com/")
	}
	if err != nil {
		fmt.Println("Unexpected error:", err)
	}
	fmt.Printf("gh cli version: %s\n", ghversionout)

	// Check if user is authenticated in github
	ghusercommand := "gh config get user -h github.com"
	ghuserout, err := utils.ExecuteBashCommand(ghusercommand)
	if err != nil {
		fmt.Println("User not authenticated, Run Command: gh auth login")
		log.Fatal(err)
	}
	fmt.Printf("Authenticated user account: %s", ghuserout)

	//Print eveything is ok
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)
	boldGreen.Printf("Your system is ready to use ghvars")
}
