package installer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func extractZip(source, dst string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {

		return err
	}
	defer reader.Close()

	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		err := extractZipItem(f, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractZipItem(f *zip.File, dst string) error {
	fp := filepath.Join(dst, f.Name)
	if !strings.HasPrefix(fp, filepath.Clean(dst)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", fp)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(fp, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
		return err
	}

	fs, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer fs.Close()

	zf, err := f.Open()
	if err != nil {
		return err
	}
	defer zf.Close()

	if _, err := io.Copy(fs, zf); err != nil {
		return err
	}
	return nil
}
