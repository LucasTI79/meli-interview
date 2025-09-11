package helpers

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func ProjectRoot() string {

	if root := os.Getenv("PROJECT_ROOT"); root != "" {
		return root
	}

	if cwd, err := os.Getwd(); err == nil {
		return cwd
	}

	if exe, err := os.Executable(); err == nil {
		return filepath.Dir(exe)
	}

	log.Fatal("unable to determine project root")
	return ""
}

func PathInRoot(relPath string) string {
	return filepath.Join(ProjectRoot(), relPath)
}

func EnsureDir(relPath string) error {
	path := PathInRoot(relPath)
	return os.MkdirAll(path, os.ModePerm)
}

func CreateFile(relPath string) (*os.File, error) {
	dir := filepath.Dir(PathInRoot(relPath))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(PathInRoot(relPath))
}

func SaveJSON(relPath string, data any) error {
	f, err := CreateFile(relPath)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
