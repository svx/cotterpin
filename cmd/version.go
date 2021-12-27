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
	"runtime"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

// Default build-time variable.
// These values are overridden via ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long: `Show version of cotterpin.
For example:

Version: dev
Build Date: unknown
Git Commit: none
OS: linux
Arch: amd64
Go Version: go1.17

Please provide above output for support`,
	Run: func(cmd *cobra.Command, args []string) {
		println(color.Yellow + "Version and build Info\n" + color.Reset)
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", date)
		fmt.Printf("Git Commit: %s\n", commit)
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Arch: %s\n", runtime.GOARCH)
		fmt.Printf("Go Version: %s\n\n", runtime.Version())

		fmt.Printf("Please provide above output for support\n\n")

		//println(color.Yellow + "Tech support: support@foobar\n" + color.Reset)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
