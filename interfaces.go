package plugins

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Plugin to description the plugin information and what to do and need what
// Also define some configuration settings for loader and web-injector to use
type Plugin struct {
	// Plugin Manifest
	Name        string `json:"name"`
	PackageID   string `json:"pkg"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Author      string `json:"author"`
	Version     string `json:"version"`

	// Plugin Assets
	Assets *PluginAssets `json:"assets"`

	// Plugin Hooks
	Setup   func(p *Plugin, router gin.IRouter) `json:"-"`
	Migrate func(datasource *gorm.DB)           `json:"-"`
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
	Descriptions       string                 `json:"descriptions"`
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
