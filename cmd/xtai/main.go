package main

import (
	"os"

	"github.com/x3ns/xtai-cli/internal/delivery"
)

func main() {
	if err := delivery.Execute(); err != nil {
		os.Exit(1)
	}
}

