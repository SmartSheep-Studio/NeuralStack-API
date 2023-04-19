package installer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func CleanupFolder(source string) {
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatalf("Error when cleanning up installer cache: %w", err)
	}

	for _, info := range entries {
		if strings.HasSuffix(info.Name(), ".installer-cache") && info.IsDir() {
			os.RemoveAll(filepath.Join(source, info.Name()))
		} else if strings.HasSuffix(info.Name(), ".plug.zip") {
			os.Remove(filepath.Join(source, info.Name()))
		}
	}
}
