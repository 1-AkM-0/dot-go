package main

import (
	"log"

	"github.com/1-AkM-0/dot-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
