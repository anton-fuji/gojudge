package cmd

import (
	"fmt"
	"os"

	"github.com/anton-fuji/gojudge/internal/judge"
	"github.com/spf13/cobra"
)

// ans -l 用
var results []*judge.Result

var checkCmd = &cobra.Command{
	Use:   "check <file.go>",
	Short: "Check code against a problem",
	Long:  "Check a Go file against a problem's test cases and return AC/WA",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		problemID, _ := cmd.Flags().GetString("problem")
		result, err := judge.CheckSolution(filename, false, problemID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		results = append(results, result)
		if result.Passed {
			fmt.Printf("✅ AC - %s\n", result.Problem.Title)
		} else {
			fmt.Printf("❌ WA - %s\n", result.Problem.Title)
		}
	},
}

func init() {
	checkCmd.Flags().StringP("problem", "p", "1", "Problem ID")
}
