package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPath(filename string) (string, error) {
	var baseDir string

	switch runtime.GOOS {
	case "windows":
		baseDir = os.Getenv("APPDATA")
	case "darwin":
		baseDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	default:
		baseDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	cliStorageDir := filepath.Join(baseDir, "ant")
	err := os.MkdirAll(cliStorageDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(cliStorageDir, filename), nil
}
