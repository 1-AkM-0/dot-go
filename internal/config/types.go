package config

import "time"

type DotFile struct {
	id           int       `json:"id"`
	name         string    `json:"name"`
	source       string    `json:"source"`
	targetPath   string    `json:"target_path"`
	addedAt      time.Time `json:"added_at"`
	lastModified time.Time `json:"last_modified"`
}

type Config struct {
	DotFiles map[string]DotFile `json:"dotfiles"`
}

func NewConfig() *Config {
	return &Config{
		DotFiles: make(map[string]DotFile),
	}
}
