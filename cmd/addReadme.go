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
	//"errors"
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
	"time"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// https://www.javaguides.net/2021/05/go-golang-read-input-from-user-or.html
// http://www.inanzzz.com/index.php/post/c4ul/a-terminal-cli-application-accepting-plain-and-secret-input-from-the-user-in-golang

//go:embed templates/readme.tmpl
var tmplReadme []byte

// addReadmeCmd represents the addReadme command
var addReadmeCmd = &cobra.Command{
	Use:    "readme",
	Short:  "Add a README",
	Long:   `Add a README to a project.`,
	PreRun: toggleDebug, // This is for logging.
	Run: func(cmd *cobra.Command, args []string) {
		forceReadme, _ := cmd.Flags().GetBool("force")
		if forceReadme {
			backupReadme()
			removeReadme()
			addReadme()
		} else {
			addReadme()
			github()
		}
	},
}

func addReadme() {
	// Check if file, exists, if yes fail with error message
	if _, err := os.Stat("README.md"); err == nil {
		fmt.Printf("File already exists, if you really want to overwrite, use --force")
	} else {
		// Create a new file
		tmpl := template.New("readme")
		tmpl, err := tmpl.Parse(string(tmplReadme))
		if err != nil {
			log.Fatal("Error Parsing template: ", err)
			return
		}
		reader := bufio.NewReader(os.Stdin)

		color.Yellow("Enter project name:")

		type input struct {
			Name string
		}
		name, _ := reader.ReadString('\n')

		err1 := tmpl.Execute(os.Stdout, input{name})
		if err1 != nil {
			log.Fatal("Error using template: ", err1)
		}
		// Create a new file
		color.Yellow("Creating README")
		file, _ := os.Create("README.md")
		defer file.Close()
		tmpl.Execute(file, input{name})
	}
}

func backupReadme() {
	// Open original file (README.md)
	original, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	// Create new file (README.md-bak)
	color.Yellow("Creating backup of README")
	currentTime := time.Now()
	new, err := os.Create("README-" + currentTime.Format("01-02-2006"))
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	// This will copy
	// bytesWritten, err := io.Copy(new, original)
	_, err = io.Copy(new, original)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Bytes Written: %d\n", bytesWritten)
}

func removeReadme() {
	err := os.Remove("README.md")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("file deleted")
}

func github() {
	// check if .github/workflows/repo-qa.yml already exists
	//if not copy files over from templates
	dirname := "github"
	if !strings.HasPrefix(dirname, ".") {
		dirname = "." + dirname
	}


	err := os.Mkdir(dirname, 0700)
	if err != nil {
		log.Fatal(err)
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
	addReadmeCmd.Flags().BoolP("force", "f", false, "Backup existent README.md and overwrite it")
}
