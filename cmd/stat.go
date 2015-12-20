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

	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

type statistic struct {
	begin, end                                                                                                                             time.Time
	absFilePath                                                                                                                            string
	uniquePayees, uniqueAccounts, postings, unclearedPostings, daysSinceLastPosts, postsLastSevenDays, postsLastThirtyDays, postsThisMonth int
	postingsPerDay                                                                                                                         float64
}

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "A brief description of your command",
	Long: `Provides summary information about all the postings matching report-query It provides information such as:
	- Time range of all matching postings
	- Unique payees
	- Unique accounts
	- Postings total
	- Uncleared postings
	- Days since last posting
	- More...
`,
	Run: func(cmd *cobra.Command, args []string) {
		if LedgerFile != "" {
			absPath, _ := filepath.Abs(LedgerFile)
			printStat(readStat(absPath))
		} else {
			cmd.Help()
		}
	},
}

var strict, pedantic, checkPayees, immediate bool

func init() {
	RootCmd.AddCommand(statCmd)
	statCmd.Flags().BoolVar(&strict, "strict", false, "Accounts, tags or commodities not previously declared will cause warnings.")
	statCmd.Flags().BoolVar(&pedantic, "pedantic", false, "Accounts, tags or commodities not previously declared will cause errors.")
	statCmd.Flags().BoolVar(&checkPayees, "check-payees", false, "Enable strict and pedantic checking for payees as well as accounts, commodities and tags. This only works in conjunction with --strict or --pedantic.")
	statCmd.Flags().BoolVar(&immediate, "immediate", false, "Instruct ledger to evaluate calculations immediately rather than lazily.")
}

func readStat(absPath string) statistic {
	stats := statistic{time.Now(), time.Now(), absPath, 0, 0, 0, 0, 0, 0, 0, 0, 0.0}
	reader, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		defer reader.Close()
	}
	return stats
}

var template = `Time period: %s to %s (%1.f days)

  Files these postings came from:
	%s

  Unique payees:             %8d
  Unique accounts:           %8d

  Number of postings:        %8d (%2.2f per day)
  Uncleared postings:        %8d

  Days since last post:      %8d
  Posts in last 7 days:      %8d
  Posts in last 30 days:     %8d
  Posts seen this month:     %8d
`

func printStat(st statistic) {
	fmt.Printf(template, st.begin.Format("2006/01/02"), st.end.Format("2006/01/02"), st.end.Sub(st.begin).Hours()/24, st.absFilePath, st.uniquePayees, st.uniqueAccounts, st.postings, st.postingsPerDay, st.unclearedPostings, st.daysSinceLastPosts, st.postsLastSevenDays, st.postsLastThirtyDays, st.postsThisMonth)
}
