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

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Git repository or reinitialize an existing one",
	Long: `This command creates an empty Git repository - basically a .git directory with subdirectories for objects, refs/heads, refs/tags, and template files.
	An initial branch without any commits will be created (see the --initial-branch option below for its name).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		quiet, _ := cmd.Flags().GetBool("quiet")
		if quiet {
			fmt.Println("enabled flag quiet")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("quiet", "q", false, "be quiet")
}
