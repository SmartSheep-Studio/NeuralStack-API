package installer

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"repo.smartsheep.studio/smartsheep/neuralstack-api/filesystem"
)

func CheckEnvironment() error {
	if runtime.GOOS == "windows" {
		return fmt.Errorf("unsupported platform of plugin loader: %s", runtime.GOOS)
	}

	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("missing go: %w", err)
	} else if _, err := exec.LookPath("gcc"); err != nil {
		return fmt.Errorf("missing gcc: %w", err)
	}

	return nil
}

func CompliePlugin(manifest api.PluginManifest, source string) error {
	log.Infof("Starting complie plugin %s v%s", manifest.Name, manifest.Version)
	for _, script := range manifest.Installer.Scripts {
		slice := strings.Split(script, " ")
		cmd := exec.Command(slice[0], slice[1:]...)
		cmd.Dir = filepath.Join(filesystem.GetAbsRoot(), source)
		if output, err := cmd.Output(); err != nil {
			log.Errorf(string(output))
			return err
		} else {
			log.Infof(string(output))
		}
	}
	log.Infof("Success complie plugin %s v%s", manifest.Name, manifest.Version)

	return nil
}
