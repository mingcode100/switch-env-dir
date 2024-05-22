/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jvm",
	Short: "Java Version Manager",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "list all jdk",
	Long:  `list all jdk`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		FindJDKs()
	},
}

var jdkVersion string
var cmdUse = &cobra.Command{
	Use:   "use",
	Short: "use jdk",
	Long:  `use jdk`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if jdkVersion == "" {
			_ = cmd.Help()
			os.Exit(0)
		}
		UseJDK(jdkVersion)
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
	cmdUse.Flags().StringVarP(&jdkVersion, "version", "v", "", "jdk version")

	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdUse)
}
