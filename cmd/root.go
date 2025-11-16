package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "dotgo",
	Short: "DotGo is a CLI that manage your dotfiles",
}

func Execute() error {
	return rootCmd.Execute()
}
