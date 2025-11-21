package config

import "fmt"

func (c *Config) Add(dotFile DotFile) error {
	if _, exists := c.DotFiles[dotFile.Name]; exists {
		return fmt.Errorf("%s dotfile already exists", dotFile.Name)
	}
	c.DotFiles[dotFile.Name] = dotFile
	return nil
}
