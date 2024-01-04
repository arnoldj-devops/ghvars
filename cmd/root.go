package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var (
	setDryRun bool
)
var rootCmd = &cobra.Command{
	Use:     "ghvars [sub]",
	Version: version,
	Short:   "ghvars",
}

// Command to check the prerequisite conditions are met or not
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "for troubleshooting",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		doctor()
	},
}

// Command to get the variables
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get variables",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		getVars()
	},
}

// Command to set the variables
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set variables",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		ghEnvironmentsList, err := ListEnvironments()
		
		if err != nil {
			fmt.Println("Error fetching environments:", err)
			os.Exit(1)
		}
		if (env != "all") && (env != "none") {
			if !envValid(ghEnvironmentsList, env) {
				fmt.Printf("%s is an invalid environment.\n", env)
				os.Exit(1)
			}
		}
		setVars(env)
	},
}

func Execute() {
	err := doctorCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd, getCmd, setCmd)
	setCmd.Flags().BoolVar(&setDryRun, "dryrun", false, "Show changes without actually setting variables")
	setCmd.PersistentFlags().StringP("env", "e", "none", "Github Environment")
}

// Function to check if a string is present in a []string variable
func envValid(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
