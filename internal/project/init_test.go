package project

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitProjectCreatesRequiredStructure(t *testing.T) {
	root := t.TempDir()

	result, err := InitProject(context.Background(), InitProjectRequest{Root: root})
	if err != nil {
		t.Fatalf("InitProject returned error: %v", err)
	}

	if result.Root != root {
		t.Fatalf("Root = %q, want %q", result.Root, root)
	}

	for _, dir := range defaultDirectories {
		assertDir(t, filepath.Join(root, dir))
	}
	for _, file := range defaultFiles {
		assertFile(t, filepath.Join(root, file.Path))
	}

	if len(result.CreatedDirs) != len(defaultDirectories) {
		t.Fatalf("CreatedDirs length = %d, want %d", len(result.CreatedDirs), len(defaultDirectories))
	}
	if len(result.CreatedFiles) != len(defaultFiles) {
		t.Fatalf("CreatedFiles length = %d, want %d", len(result.CreatedFiles), len(defaultFiles))
	}
}

func TestInitProjectPreservesExistingFilesAndReportsSkipped(t *testing.T) {
	root := t.TempDir()
	projectFile := filepath.Join(root, ".yllmcode", "project.md")

	if err := os.MkdirAll(filepath.Dir(projectFile), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(projectFile, []byte("custom project notes\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := InitProject(context.Background(), InitProjectRequest{Root: root})
	if err != nil {
		t.Fatalf("InitProject returned error: %v", err)
	}

	content, err := os.ReadFile(projectFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "custom project notes\n" {
		t.Fatalf("project.md was overwritten: %q", string(content))
	}

	if !contains(result.SkippedFiles, ".yllmcode/project.md") {
		t.Fatalf("SkippedFiles = %#v, want project.md", result.SkippedFiles)
	}
}

func TestInitProjectForceOverwritesDefaultFiles(t *testing.T) {
	root := t.TempDir()
	projectFile := filepath.Join(root, ".yllmcode", "project.md")

	if err := os.MkdirAll(filepath.Dir(projectFile), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(projectFile, []byte("custom project notes\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	result, err := InitProject(context.Background(), InitProjectRequest{
		Root:  root,
		Force: true,
	})
	if err != nil {
		t.Fatalf("InitProject returned error: %v", err)
	}

	content, err := os.ReadFile(projectFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) == "custom project notes\n" {
		t.Fatal("project.md was not overwritten with force")
	}
	if contains(result.SkippedFiles, ".yllmcode/project.md") {
		t.Fatalf("SkippedFiles = %#v, did not want project.md", result.SkippedFiles)
	}
}

func TestInitProjectRejectsInvalidRoot(t *testing.T) {
	root := filepath.Join(t.TempDir(), "missing")

	_, err := InitProject(context.Background(), InitProjectRequest{Root: root})
	if err == nil {
		t.Fatal("InitProject returned nil error for missing root")
	}
	if !strings.Contains(err.Error(), "inspect project root") {
		t.Fatalf("error = %q, want useful root inspection error", err)
	}
}

func TestFindProjectRootUsesNearestMarker(t *testing.T) {
	root := t.TempDir()
	nested := filepath.Join(root, "a", "b")

	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.test/root\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := FindProjectRoot(nested)
	if err != nil {
		t.Fatalf("FindProjectRoot returned error: %v", err)
	}
	if got != root {
		t.Fatalf("FindProjectRoot = %q, want %q", got, root)
	}
}

func assertDir(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat directory %q: %v", path, err)
	}
	if !info.IsDir() {
		t.Fatalf("%q is not a directory", path)
	}
}

func assertFile(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat file %q: %v", path, err)
	}
	if info.IsDir() {
		t.Fatalf("%q is a directory", path)
	}
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
