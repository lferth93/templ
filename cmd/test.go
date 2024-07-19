/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:        "test",
	Short:      "Tests a template file.",
	Long:       `Test if a template file is correct and can be parsed.`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"template"},
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Println("test called with", args, cmd.Flags().Changed("html"), cmd.Flags().Changed("json"))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
