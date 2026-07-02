package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot(startPath string) (string, error) {
	if startPath == "" {
		startPath = "."
	}

	abs, err := filepath.Abs(startPath)
	if err != nil {
		return "", fmt.Errorf("resolve project root: %w", err)
	}

	info, err := os.Stat(abs)
	if err != nil {
		return "", fmt.Errorf("inspect project root %q: %w", abs, err)
	}
	if !info.IsDir() {
		abs = filepath.Dir(abs)
	}

	for dir := abs; ; dir = filepath.Dir(dir) {
		if exists(filepath.Join(dir, ".git")) || exists(filepath.Join(dir, "go.mod")) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return abs, nil
		}
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
