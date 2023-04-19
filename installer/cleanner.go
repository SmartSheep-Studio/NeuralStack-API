package installer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func CleanupFolder(source string) {
	filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error when cleanning up installer cache: %w\n", err)
		} else if !info.IsDir() || !strings.HasSuffix(info.Name(), ".installer-cache") {
			return nil // Skip this folder if it isn't a installer cache
		} else {
			p, _ := os.Getwd()
			os.RemoveAll(filepath.Join(p, info.Name()))
		}

		return nil
	})
}
