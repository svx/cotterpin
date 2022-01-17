/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	_ "embed"
	"text/template"
	"os"
	"log"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

//go:embed templates/dockerfile.tmpl
var tmplDockerfile []byte

// addDockerfileCmd represents the addDockerfile command
var addDockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Add Dockerfile",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addDockerfile called")
		addDockerFile()
	},
}

func addDockerFile() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Alpine", "Debian", "Node"},
	}

	_, image, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", image)
	type input struct {
		Image string
	}
	if _, err := os.Stat("Dockerfile"); err == nil {
		fmt.Printf("File already exists")
	} else {
		// Create a new file
		tmpl := template.New("dockerfile")
		tmpl, err := tmpl.Parse(string(tmplDockerfile))
		if err != nil {
			log.Fatal("Error Parsing template: ", err)
			return
		}
		err1 := tmpl.Execute(os.Stdout, input{image})
		if err1 != nil {
			log.Fatal("Error using template: ", err1)
		}
		// Create a new file
		color.Yellow("Creating Dockerfile")
		file, _ := os.Create("Dockerfile")
		defer file.Close()
		tmpl.Execute(file, input{image})
	}

}

func init() {
	addCmd.AddCommand(addDockerfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addDockerfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addDockerfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
