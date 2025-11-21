package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/1-AkM-0/dot-go/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file/directory]",
	Short: "Add a file or a directory to DotGo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.Load()
		if err != nil {
			log.Fatal("Error loading JSON config")
		}

		home, err := os.UserConfigDir()
		if err != nil {
			log.Fatal("Error looking for .config/")
		}

		name := args[0]

		targetPath := filepath.Join(home, "/", name)
		sourcePath := filepath.Join(home, "dot-go")

		dotFile := config.DotFile{
			Name:         name,
			Id:           0,
			TargetPath:   targetPath,
			Source:       sourcePath,
			AddedAt:      time.Now(),
			LastModified: time.Now(),
		}

		if err := c.Add(dotFile); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
