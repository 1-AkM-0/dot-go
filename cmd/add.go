package cmd

import (
	"fmt"

	"github.com/1-AkM-0/dot-go/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file/directory]",
	Short: "Add a file or a directory to DotGo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.AddFile(args[0]); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
