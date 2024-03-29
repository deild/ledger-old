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

	"bufio"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print transactions in a format readable by ledger",
	Long: `Prints out the full transactions of any matching postings using the same format as they would appear in a data file.
This can be used to extract subsets from a Ledger file to transfer to other files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
		if LedgerFile != "" {
			ledgerFileReader, err := os.Open(LedgerFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				defer ledgerFileReader.Close()
				// read and print file
				scan := bufio.NewScanner(ledgerFileReader)
				scan.Split(bufio.ScanLines)
				for scan.Scan() {
					line := scan.Text()
					switch {
					case strings.HasPrefix(line, ";"): //ingored
					case strings.HasPrefix(line, "#"): //ingored
					case strings.HasPrefix(line, "%"): //ingored
					case strings.HasPrefix(line, "|"): //ingored
					case strings.HasPrefix(line, "*"): //ingored
					default:
						fmt.Println(line)
					}
				}
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(printCmd)
}
