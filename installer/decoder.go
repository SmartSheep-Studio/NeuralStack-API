package installer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"time"
)

func randomID(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

func DecodeInstallPack(source string, dst string) (string, api.PluginManifest, error) {
	workspace := filepath.Join(dst, fmt.Sprintf("%s%s", randomID(18), ".installer-cache"))

	// Extract zip into temporary folder
	extractZip(source, workspace)

	// Read `manifest.json`
	var manifest api.PluginManifest
	rawManifest, err := os.ReadFile(filepath.Join(workspace, "manifest.json"))
	if err != nil {
		return workspace, api.PluginManifest{}, err
	} else {
		if err := json.Unmarshal(rawManifest, &manifest); err != nil {
			return workspace, api.PluginManifest{}, err
		}
	}

	return workspace, manifest, nil
}
