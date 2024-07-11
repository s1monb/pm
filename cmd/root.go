/*
Copyright © 2024 SIMON BJØRNØY <SIMON.BJORNOY@GMAIL.COM>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/s1monb/proj-mgmt/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "proj-mgmt",
	Short: "A project management tool",
}

var projectCmd = &cobra.Command{
	Use:     "project",
	Short:   "Manage projects",
	Aliases: []string{"p"},
}

var projectListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all projects",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		project.List()
	},
}

var projectNewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new project",
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		project.New()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initConfigFile()

	// PROJECT-COMMANDS
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectListCmd)
	projectCmd.AddCommand(projectNewCmd)
}

func initConfigFile() {
	viper.SetConfigName(".pm")   // name of config file (without extension)
	viper.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME") // call multiple times to add many search paths
	viper.AddConfigPath(".")     // optionally look for config in the working directory
	err := viper.ReadInConfig()  // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
