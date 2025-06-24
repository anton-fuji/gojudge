package cmd

import (
	"fmt"

	"github.com/anton-fuji/gojudge/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var raretechCmd = &cobra.Command{
	Use:   "raretech",
	Short: "RareTECH Logo Ascii Art",
	Run: func(cmd *cobra.Command, args []string) {
		selectType()
	},
}

func init() {
	rootCmd.AddCommand(raretechCmd)
}

func selectType() {
	prompt := promptui.Select{
		Label:     "What AsciiArt do you want to display??",
		Items:     []string{"raretech", "raretech-logo"},
		CursorPos: 0,
	}

	idx, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose no.%d %q\n", idx+1, result)

	switch result {
	case "raretech":
		utils.PrintAAFromTxt("raretech.txt")
	case "raretech-logo":
		utils.PrintAAFromTxt("raretech-logo.txt")
	}
}
