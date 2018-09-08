package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all the note titles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ALL THE LIST")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
