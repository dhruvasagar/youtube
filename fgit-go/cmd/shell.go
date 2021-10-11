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
	"io"
	"reflect"
	"strings"
	"unsafe"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Launches an interactive fgit shell",
	Long:  `Launches an interactive fgit shell allowing the user to continually invoke fgit commands`,
	Run: func(cmd *cobra.Command, args []string) {
		completer := readline.NewPrefixCompleter(
			readline.PcItem("help",
				readline.PcItem("init"),
				readline.PcItem("shell"),
				readline.PcItem("status"),
			),
			readline.PcItem("init",
				readline.PcItem("-q"),
			),
			readline.PcItem("shell"),
			readline.PcItem("status"),
		)
		rl, err := readline.NewEx(&readline.Config{
			Prompt:            "> ",
			InterruptPrompt:   "^C",
			EOFPrompt:         "exit",
			HistorySearchFold: true,
			AutoComplete:      completer,
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		for {
			line, err := rl.Readline()
			if err == readline.ErrInterrupt || err == io.EOF {
				break
			}

			line = strings.TrimSpace(line)
			switch {
			case strings.HasPrefix(line, "shell"):
				fmt.Println("Already in shell")
			case line == "exit" || line == "bye":
				goto exit
			default:
				executeCommand(line)
			}
		}
	exit:
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}

func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Value.Type() == "stringSlice" {
			// Unfortunately, flag.Value.Set() appends to original
			// slice, not resets it, so we retrieve pointer to the slice here
			// and set it to new empty slice manually
			value := reflect.ValueOf(flag.Value).Elem().FieldByName("value")
			ptr := (*[]string)(unsafe.Pointer(value.Pointer()))
			*ptr = make([]string, 0)
		}

		flag.Value.Set(flag.DefValue)
	})
	for _, cmd := range cmd.Commands() {
		resetFlags(cmd)
	}
}

func executeCommand(cmdName string) {
	resetFlags(rootCmd)

	args := strings.Fields(cmdName)
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("%v\n", err)
		rootCmd.Help()
	}
}
