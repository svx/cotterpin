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
	//"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	//getter "github.com/hashicorp/go-getter"
	"github.com/spf13/cobra"
)

// initDocsCmd represents the initDocs command
var initDocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Setup linting for documentation",
	Long: `Enables the following linting tools:
- alex for prose
- markdown-lint for Markdown
- markdown-link-checker
`,
	PreRun: toggleDebug, // This is for logging.
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initDocs called")
		color.Yellow("Downloading config files")
		//getConfig()
		addDocsDir()
	},
}

func addDocsDir() {
	// Check if docs direxists, if yes fail with error message
	if _, err := os.Stat("docs"); err == nil {
		color.Red("docs directory already exists")
	} else {
		color.Green("Creating docs directory")
		os.Mkdir("docs", 0700)
	}
}

// func getConfig() {
// 	client := &getter.Client{
// 		Ctx: context.Background(),
// 		// Define the destination to where the directory will be stored.
// 		// This will create the directory if it doesnt exist
// 		Dst: ".doc_config-test",
// 		Dir: true,
// 		// the repository with a subdirectory I would like to clone only
// 		Src:  "github.com/svx/cotterpin.git//.doc_config",
// 		Mode: getter.ClientModeDir,
// 		// define the type of detectors go getter should use, in this case only github is needed
// 		Detectors: []getter.Detector{
// 			&getter.GitHubDetector{},
// 		},
// 		// provide the getter needed to download the files
// 		Getters: map[string]getter.Getter{
// 			"git": &getter.GitGetter{},
// 		},
// 	}
// 	// download the files
// 	if err := client.Get(); err != nil {
// 		fmt.Fprintf(os.Stderr, "Error getting path %s: %v", client.Src, err)
// 		os.Exit(1)
// 	}
// 	// now you should check your temp directory for the files to see if they exist
// }

func init() {
	initCmd.AddCommand(initDocsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initDocsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initDocsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
