package config

import (
	"fmt"
	"os"
)

func AddFile(filePath string) error {
	fullpath := "~/.config/" + filePath
	fmt.Println(fullpath)
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		return fmt.Errorf("file not found :%s", filePath)
	}
	fmt.Printf("File added: %s", filePath)
	return nil
}
