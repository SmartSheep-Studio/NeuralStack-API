package installer

import (
	"fmt"
	"os"
	"path/filepath"
)

func InstallPlugin(source string, dst string) error {
	if err := CheckEnvironment(); err != nil {
		return err
	}

	workspace, manifest, err := DecodeInstallPack(source, filepath.Dir(source))
	if err != nil {
		return err
	} else {
		if err := CompliePlugin(manifest, workspace); err != nil {
			return err
		}
	}

	if stat, err := os.Stat(filepath.Join(workspace, manifest.Installer.Output)); stat.IsDir() || err != nil {
		return fmt.Errorf("complie failed, output doesn't exists: %s", err)
	} else {
		if err := os.Rename(filepath.Join(workspace, manifest.Installer.Output), filepath.Join(dst, manifest.Installer.Output)); err != nil {
			return err
		}
	}

	return nil
}
