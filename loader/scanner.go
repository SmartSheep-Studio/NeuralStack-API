package loader

import (
	"github.com/charmbracelet/log"
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanFolder(dst string) ([]string, error) {
	var plugins []string
	return plugins, filepath.Walk(dst, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error when scanning plugins: %w", err)
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".plug.so") {
			return nil // Skip this file if it is not a plugin
		}

		plugins = append(plugins, path)
		return nil
	})
}
