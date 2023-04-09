package utils

import (
	"embed"
	"github.com/gin-contrib/static"
	"io/fs"
	"net/http"
)

type EmbedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e EmbedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return EmbedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
