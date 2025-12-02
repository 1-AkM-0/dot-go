package manager

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/1-AkM-0/dot-go/internal/config"
)

var m = new(Manager)

func TestList(t *testing.T) {
	c := configTestConfig(t)
	buffer := bytes.Buffer{}
	m = &Manager{
		Out: &buffer,
	}
	dotfile := config.DotFile{
		Name:       "nvim",
		Source:     "$HOME/.config/dotgo",
		TargetPath: t.TempDir(),
	}

	if err := c.Add(dotfile); err != nil {
		t.Errorf("Add fail(), dotfile was not added")
	}

	m.List(c)

	output := buffer.String()
	if !strings.Contains(output, "nvim") {
		t.Errorf("Expect find 'nvim', but got %s", output)
	}
}

func configTestConfig(t testing.TB) *config.Config {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "config_test.json")

	c := config.NewConfig()
	c.FilePath = tmpFile

	return c
}
