package utils

import "github.com/charmbracelet/log"

func HandleFatalError(stage string, err error) {
	if err != nil {
		log.Fatalf("Error when %s: %w", stage, err)
	}
}
