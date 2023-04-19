package installer

import (
	"github.com/charmbracelet/log"
	"io/fs"
	"path/filepath"
	"strings"
)

func ScanFolder(dst string) ([]string, error) {
	var packs []string
	return packs, filepath.Walk(dst, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Error when scanning plugins: %w\n", err)
		} else if info.IsDir() || !strings.HasSuffix(info.Name(), ".plug.zip") {
			return nil // Skip this file if it is not a plugin install pack
		}

		packs = append(packs, path)
		return nil
	})
}
