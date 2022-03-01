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
	//"fmt"
	_ "embed"
	"log"
	"fmt"
	"os"
	"time"
	"text/template"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

//go:embed templates/dependabot.tmpl
var f string

// initRepoCmd represents the initRepo command
var initRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Initialize repository structure and setup",
	Long: `Initialize repository structure and setup.

	Creates .github directory with:
	- Issue and PR templates
	- Workflows and GitHub actions
	- Dependabot`,
	PreRun: toggleDebug, // This is for logging.
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("Initialize repo setup")
		time.Sleep(2 * time.Second)
		// getConfig()
		addGitHubDir()
		addDependabot()
	},
}

func addGitHubDir() {
	// Check if .git directory exists in current directory
	if _, err := os.Stat(".git"); err == nil {
		color.Green("Git directory exists")
		color.Yellow(">> Create .github directory")
		os.Mkdir(".github", 0o700)

	} else {
		prompt := promptui.Prompt{
			Label:     "It seems Git is not initialized yet, should I continue?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			color.Red("Ok, stopping for now")
			return
		}

		fmt.Printf("You choose %q\n", result)
		color.Yellow(">> Creating .github directory")
		os.Mkdir(".github", 0o700)
	}
}

func addDependabot() {
	if _, err := os.Stat(".github/dependabot.yml"); err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			color.Yellow(">> Checking for dependabot")
			//fmt.Println(f)
			tmpl := template.New("dependabot")
			tmpl, err := tmpl.Parse(string(f))
			if err != nil {
				log.Fatal("Error Parsing template: ", err)
				return
			}
			color.Yellow(">> Adding dependabot")
			err1 := tmpl.Execute(os.Stdout, f)
			if err1 != nil {
				log.Fatal("Error executing template: ", err1)
			}
			// Create a new file
			//color.Yellow("%v\n>> Adding dependabot")
			dependabotFile, _ := os.Create(".github/dependabot.yml")
			defer dependabotFile.Close()
			tmpl.Execute(dependabotFile, f)
			}
	}
}


// func addGitHubDir() {
// 	// Check if docs dir exists, if yes fail with error message
// 	if _, err := os.Stat(".github"); err == nil {
// 		color.Red("docs directory already exists")
// 	} else {
// 		color.Yellow(">> Create .github directory")
// 		os.Mkdir(".github", 0o700)
// 	}
// }

func init() {
	initCmd.AddCommand(initRepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initRepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initRepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
