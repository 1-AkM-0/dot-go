package manager

import (
	"fmt"
	"os"

	"github.com/1-AkM-0/dot-go/internal/config"
)

type Manager struct{}

func (m *Manager) List(c *config.Config) {
	for _, df := range c.DotFiles {
		status := "OK"
		if _, err := os.Stat(df.TargetPath); os.IsNotExist(err) {
			status = "MISSING"
		}
		fmt.Fprintf(os.Stdout, "[%s] %s -> %s \n", status, df.Name, df.TargetPath)
	}
}
