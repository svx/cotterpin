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
	"embed"
    //"io/fs"
	"text/template"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

//go:embed templates/*
var files embed.FS

//go:embed templates/readme.tmpl
var tmplReadme []byte

// readerCmd represents the reader command
var readerCmd = &cobra.Command{
	Use:   "reader",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reader called")
		//templates, _ := fs.ReadDir(files, "templates")
		//for _, template := range templates {
		//	fmt.Printf("%q\n", template.Name())
		//}
		tmpl := template.New("test")
	tmpl, err := tmpl.Parse(string(tmplReadme))
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}

	type data struct {
		Name string
	}

	d := data{Name: "Jane"}
	err1 := tmpl.Execute(os.Stdout, d)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)
	}
	// Create a new file
	color.Green("Creating README")
	file, _ := os.Create("README.foo.md")
	defer file.Close()
	tmpl.Execute(file, d)
	},
}

func init() {
	rootCmd.AddCommand(readerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
