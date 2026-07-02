package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type InitProjectRequest struct {
	Root  string
	Force bool
}

type InitProjectResult struct {
	Root         string
	CreatedDirs  []string
	CreatedFiles []string
	SkippedFiles []string
}

func InitProject(ctx context.Context, req InitProjectRequest) (*InitProjectResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	root, err := FindProjectRoot(req.Root)
	if err != nil {
		return nil, err
	}
	if err := validateRoot(root); err != nil {
		return nil, err
	}

	createdDirs, err := CreateDefaultDirectories(root)
	if err != nil {
		return nil, err
	}

	createdFiles, skippedFiles, err := CreateDefaultFiles(root, req.Force)
	if err != nil {
		return nil, err
	}

	return &InitProjectResult{
		Root:         root,
		CreatedDirs:  createdDirs,
		CreatedFiles: createdFiles,
		SkippedFiles: skippedFiles,
	}, nil
}

func CreateDefaultDirectories(root string) ([]string, error) {
	var created []string
	for _, dir := range defaultDirectories {
		path, err := safeJoin(root, dir)
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(path); err == nil {
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("inspect directory %q: %w", path, err)
		}

		if err := os.MkdirAll(path, 0o755); err != nil {
			return nil, fmt.Errorf("create directory %q: %w", path, err)
		}
		created = append(created, dir)
	}

	return created, nil
}

func CreateDefaultFiles(root string, force bool) (created []string, skipped []string, err error) {
	for _, file := range defaultFiles {
		path, err := safeJoin(root, file.Path)
		if err != nil {
			return nil, nil, err
		}

		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return nil, nil, fmt.Errorf("create parent directory for %q: %w", path, err)
		}

		if _, err := os.Stat(path); err == nil && !force {
			skipped = append(skipped, file.Path)
			continue
		} else if err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, nil, fmt.Errorf("inspect file %q: %w", path, err)
		}

		if err := os.WriteFile(path, []byte(file.Content), 0o644); err != nil {
			return nil, nil, fmt.Errorf("write file %q: %w", path, err)
		}
		created = append(created, file.Path)
	}

	return created, skipped, nil
}

func validateRoot(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("inspect project root %q: %w", root, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("project root %q is not a directory", root)
	}
	return nil
}

func safeJoin(root, name string) (string, error) {
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve root %q: %w", root, err)
	}
	path := filepath.Join(rootAbs, name)
	rel, err := filepath.Rel(rootAbs, path)
	if err != nil {
		return "", fmt.Errorf("validate path %q: %w", path, err)
	}
	if rel == ".." || rel == "." || len(rel) >= 3 && rel[:3] == "../" {
		return "", fmt.Errorf("path %q escapes project root %q", path, rootAbs)
	}
	return path, nil
}
