package plugins

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Plugin to description the plugin information and what to do and need what
// Also define some configuration settings for loader and web-injector to use
type Plugin struct {
	// Manifest
	// Recommend use embed.FS and json.Unmarshal to embed manifest.json
	Manifest PluginManifest `json:"manifest"`

	// Plugin Assets
	Assets *PluginAssets `json:"assets"`

	// Plugin Hooks
	Setup   func(p *Plugin, router gin.IRouter) `json:"-"`
	Migrate func(datasource *gorm.DB)           `json:"-"`
}

type PluginManifest struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Package      string   `json:"package"`
	Repository   string   `json:"repository"`
	Description  string   `json:"description"`
	Authors      []string `json:"authors"`
	Dependencies []string `json:"dependencies"`

	Installer PluginInstaller `json:"installer"`
}

type PluginInstaller struct {
	Scripts []string `json:"scripts"`
	Output  string   `json:"output"`
}

type PluginAssets struct {
	Locale map[string]any `json:"-"`

	Apps []WebApp `json:"-"`
}

type WebApp struct {
	RootFile           string                 `json:"index"`
	InjectApp          string                 `json:"inject"`
	DisplayOnDesktop   bool                   `json:"shortcut"`
	DisplayOnLaunchpad bool                   `json:"display"`
	ID                 string                 `json:"id"`
	Icon               string                 `json:"icon"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	WindowOptions      WebWindowOptions       `json:"window"`
	Uri                string                 `json:"uri"`
	Authorized         bool                   `json:"authorized"`
	Perms              []string               `json:"permissions"`
	Assets             static.ServeFileSystem `json:"-"`
	PackageID          string                 `json:"pkg"`
}

type WebWindowOptions struct {
	X         int            `json:"x"`
	Y         int            `json:"y"`
	Width     int            `json:"w"`
	Height    int            `json:"h"`
	MinWidth  int            `json:"minW"`
	MinHeight int            `json:"minH"`
	MaxWidth  int            `json:"maxW"`
	MaxHeight int            `json:"maxH"`
	Resizable bool           `json:"resizable"`
	Closeable bool           `json:"closeable"`
	Meta      map[string]any `json:"meta"`
}
