package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var raretechCmd = &cobra.Command{
	Use:   "raretech",
	Short: "RareTECH Logo Ascii Art",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(content))
	},
}

func init() {
	rootCmd.AddCommand(raretechCmd)
}
