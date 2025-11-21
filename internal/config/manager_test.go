package config

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("must fail adding the same dotfile", func(t *testing.T) {
		c := NewConfig()
		dotfile1 := DotFile{
			Name:       "nvim",
			Source:     "$HOME/.config/dotgo",
			TargetPath: "$HOME/.config/",
		}

		dotfile2 := DotFile{
			Name:       "nvim",
			Source:     "$HOME/.config/dotgo",
			TargetPath: "$HOME/.config/",
		}

		if err := c.Add(dotfile1); err != nil {
			t.Errorf("Add() fail, return a unexpected error %v", err)
		}
		if err := c.Add(dotfile2); err == nil {
			t.Errorf("Add() fail, return a unexpected error %v", err)
		}
	})
	t.Run("must add a dotfile", func(t *testing.T) {
		c := NewConfig()
		dotfile := DotFile{
			Name:       "nvim",
			Source:     "$HOME/.config/dotgo",
			TargetPath: "$HOME/.config/",
		}

		if err := c.Add(dotfile); err != nil {
			t.Errorf("Add() fail, return a unexpected error %v", err)
		}
		if _, exists := c.DotFiles[dotfile.Name]; !exists {
			t.Errorf("Add() fail, dotfile nvim was not added")
		}
	})
}
