package installer

import (
	"fmt"
	"os/exec"
	"runtime"
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

func CompliePlugin(source string, scripts []string) error {
	for _, script := range scripts {
		cmd := exec.Command(script)
		cmd.Dir = source
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
