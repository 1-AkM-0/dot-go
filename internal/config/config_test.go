package config

import (
	"path/filepath"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("must fail adding the same dotfile", func(t *testing.T) {
		c := configTestConfig(t)

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
		c := configTestConfig(t)

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

func TestDelete(t *testing.T) {
	t.Run("must delete a dotfile", func(t *testing.T) {
		c := configTestConfig(t)

		dotfile := DotFile{
			Name:       "nvim",
			Source:     "$HOME/.config/dotgo",
			TargetPath: "$HOME/.config/",
		}

		if err := c.Add(dotfile); err != nil {
			t.Errorf("Add fail(), dotfile was not added")
		}

		if err := c.Delete(dotfile.Name); err != nil {
			t.Errorf("Delete() fail, return a unexpected error %v", err)
		}
		if _, exists := c.DotFiles[dotfile.Name]; exists {
			t.Errorf("Delete() fail, dotfile %s was not deleted", dotfile.Name)
		}
	})
	t.Run("must fail deleting a dotfile that does not exists", func(t *testing.T) {
		c := configTestConfig(t)

		dotfileName := "dotfile"

		if err := c.Delete(dotfileName); err == nil {
			t.Errorf("Delete() fail, return a unexpected error %v", err)
		}
	})
}

func configTestConfig(t testing.TB) *Config {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "config_test.json")

	c := NewConfig()
	c.FilePath = tmpFile

	return c
}
