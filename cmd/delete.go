package cmd

import (
	"fmt"
	"log"

	"github.com/1-AkM-0/dot-go/internal/config"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [file/directory]",
	Short: "Delete a file or a directory from DotGo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.Load()
		if err != nil {
			log.Fatal("Error loading JSON config")
			return
		}

		name := args[0]

		if err := c.Delete(name); err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("%s dotfile deleted \n", name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
