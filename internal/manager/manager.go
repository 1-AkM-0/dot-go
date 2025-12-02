package manager

import (
	"fmt"
	"io"
	"os"

	"github.com/1-AkM-0/dot-go/internal/config"
)

type Manager struct {
	Out io.Writer
}

func (m *Manager) List(c *config.Config) {
	for _, df := range c.DotFiles {
		status := "OK"
		if _, err := os.Stat(df.TargetPath); os.IsNotExist(err) {
			status = "MISSING"
		}
		fmt.Fprintf(m.Out, "[%s] %s -> %s \n", status, df.Name, df.TargetPath)
	}
}
