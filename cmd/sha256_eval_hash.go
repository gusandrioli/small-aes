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
	"github.com/gusandrioli/small-aes/sha256"
	"github.com/spf13/cobra"
)

// sha256EvalHashCmd represents the sha256EvalHash command
var sha256EvalHashCmd = &cobra.Command{
	Use:   "sha256EvalHash",
	Short: "Returns bool if input text equals hash",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sha256.EvaluateHash(args)
	},
}

func init() {
	rootCmd.AddCommand(sha256EvalHashCmd)
}
