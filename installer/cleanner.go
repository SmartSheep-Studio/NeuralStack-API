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
		} else if !info.IsDir() || (!strings.HasSuffix(info.Name(), ".installer-cache") && !strings.HasSuffix(info.Name(), ".plug.zip")) {
			return nil // Skip this folder if it isn't a installer cache or a install pack
		} else {
			os.RemoveAll(path)
		}

		return nil
	})
}
