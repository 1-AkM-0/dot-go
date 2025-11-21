package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (c *Config) Add(dotFile DotFile) error {
	if _, exists := c.DotFiles[dotFile.Name]; exists {
		return fmt.Errorf("%s dotfile already exists", dotFile.Name)
	}
	c.DotFiles[dotFile.Name] = dotFile

	return c.Save()
}

func Load() (*Config, error) {
	home, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, "dot-go", "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return NewConfig(), nil
		}
		return nil, err
	}

	cfg := NewConfig()
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, err
}

func (c *Config) Save() error {
	home, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, "dot-go", "config.json")

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}
