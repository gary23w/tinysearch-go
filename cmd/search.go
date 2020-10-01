/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewSearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Searches the dictionary for the specified term",
		Run:   Search,
	}
}

func init() {
	searchCmd := NewSearchCmd()
	rootCmd.AddCommand(searchCmd)
}

func Search(cmd *cobra.Command, args []string) {
	fmt.Println("search called")
	log.Debug("Search subcommand run with log level: ", Verbose, "\n")
}
