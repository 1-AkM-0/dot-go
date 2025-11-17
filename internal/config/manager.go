package config

import "fmt"

func (c *Config) Add(dotFile DotFile) error {
	if _, exists := c.DotFiles[dotFile.name]; exists {
		return fmt.Errorf("%s dotfile already exists", dotFile.name)
	}
	c.DotFiles[dotFile.name] = dotFile
	return nil
}
