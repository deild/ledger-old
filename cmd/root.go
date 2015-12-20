// Copyright © 2015 Samuel Marcaille <smur@free.fr>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string = "0.1.0"

var LedgerFile string
var printVersion bool

const versionMessage = `Ledger %s, the command-line accounting tool

Copyright (c) 2015, Samuel Marcaille.  All rights reserved.

This program is made available under the terms of the MIT License.
See LICENSE file included with the distribution for details and disclaimer.
`

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ledger",
	Short: "Command-line, double-entry account reporting tool",
	Long: `ledger is a command-line accounting tool based on the power and completeness of double-entry accounting.
It is only a reporting tool, which means it never modifies your data files,
but it does offers a large selection of reports, and different ways to customize them to your liking.`,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Printf(versionMessage, Version)
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&LedgerFile, "file", "f", "", "Read FILE as a ledger file.")

	RootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Print version information and exit.")

}

