/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	htemplate "html/template"
	"io"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type templater interface {
	Execute(wr io.Writer, data any) error
}

// evalCmd represents the eval command
var evalCmd = &cobra.Command{
	Use:   "eval template [template ...] data",
	Short: "Evaluates a list of template file with data.",
	Long: `Evaluates a list of template files with data and 
	prints the result to the standard output. 

	Only evaluates the first template file but it can 
	use templates defined in the other passed files.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		html, _ := cmd.Flags().GetBool("html")
		jsonF, _ := cmd.Flags().GetBool("json")

		var tmpl templater
		var err error

		if html {
			tmpl, err = htemplate.ParseFiles(args[:len(args)-1]...)
		} else {
			tmpl, err = template.ParseFiles(args[:len(args)-1]...)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		data := map[string]interface{}{}
		var content []byte
		if jsonF {
			content, err = os.ReadFile(args[len(args)-1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		} else {
			content = []byte(args[len(args)-1])
		}
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(evalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
