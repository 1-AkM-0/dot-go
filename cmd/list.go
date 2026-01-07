package cmd

import (
	"log"

	"github.com/1-AkM-0/dot-go/internal/config"
	"github.com/1-AkM-0/dot-go/internal/manager"
	"github.com/spf13/cobra"
)

var m manager.Manager

var listCmd = &cobra.Command{
	Use:   "list [DotFiles]",
	Short: "List all DotFiles managed",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.Load()
		m.Out = cmd.OutOrStdout()
		if err != nil {
			log.Fatal("Error loading JSON config")
		}
		m.List(c)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
