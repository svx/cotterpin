/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"
	"errors"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type Todo struct {
	Name string
	Description string
}

// addReadmeCmd represents the addReadme command
var addReadmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Add a README",
	Long: `Add a README to a project.`,
	PreRun: toggleDebug, // This is for logging.
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Processing")
		addReadmeFile()
	},
}

func addReadmeFile() {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("project name must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Name",
		Validate: validate,
		Default:  "My-Cool-Project",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	vars := make(map[string]interface{})
	vars["Name"] = result

	// Check if file, exists, if yes fail with error message
	if _, err := os.Stat("README.gen.md"); err == nil {
		fmt.Printf("File already exists\nRun 'cotterpin add readme -f' to overwrite it")
	} else {
		// Parse the template
		tmpl, _ := template.ParseFiles("templates/readme.tmpl")

		// Create a new file
		file, _ := os.Create("README.gen.md")
		defer file.Close()

		// Apply the template to the vars map and write the result to file.
		tmpl.Execute(file, vars)
	}
}

func init() {
	addCmd.AddCommand(addReadmeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addReadmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addReadmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addReadmeCmd.Flags().StringP("name", "n", "", "Name of the project, for example My-Cool-Project")
}
