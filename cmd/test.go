/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	htemplate "html/template"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:        "test template [template ...]",
	Short:      "Tests a template file.",
	Long:       `Test if a template file is correct and can be parsed.`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"template"},
	Run: func(cmd *cobra.Command, args []string) {
		html, _ := cmd.Flags().GetBool(`html`)
		err := error(nil)
		if html {
			_, err = htemplate.ParseFiles(args...)

		} else {
			_, err = template.ParseFiles(args...)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "All template files seems to be correct.")
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
