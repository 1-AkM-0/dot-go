package config

import "time"

type DotFile struct {
	Name         string    `json:"name"`
	Source       string    `json:"source"`
	TargetPath   string    `json:"target_path"`
	AddedAt      time.Time `json:"added_at"`
	LastModified time.Time `json:"last_modified"`
}

type Config struct {
	DotFiles map[string]DotFile `json:"dotfiles"`
}

func NewConfig() *Config {
	return &Config{
		DotFiles: make(map[string]DotFile),
	}
}
