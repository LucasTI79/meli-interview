package helpers_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasti79/meli-interview/pkg/helpers"
	"github.com/stretchr/testify/require"
)

func TestProjectRoot(t *testing.T) {
	// Backup original env
	orig := os.Getenv("PROJECT_ROOT")
	defer os.Setenv("PROJECT_ROOT", orig)

	// Test with PROJECT_ROOT env
	os.Setenv("PROJECT_ROOT", "/tmp/project")
	require.Equal(t, "/tmp/project", helpers.ProjectRoot())

	// Test without PROJECT_ROOT (falls back to cwd)
	os.Unsetenv("PROJECT_ROOT")
	root := helpers.ProjectRoot()
	require.NotEmpty(t, root)
}

func TestPathInRoot(t *testing.T) {
	os.Setenv("PROJECT_ROOT", "/tmp/project")
	defer os.Unsetenv("PROJECT_ROOT")

	path := helpers.PathInRoot("subdir/file.txt")
	require.Equal(t, filepath.Join("/tmp/project", "subdir/file.txt"), path)
}

func TestEnsureDir(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("PROJECT_ROOT", tmp)
	defer os.Unsetenv("PROJECT_ROOT")

	err := helpers.EnsureDir("nested/dir")
	require.NoError(t, err)

	info, err := os.Stat(filepath.Join(tmp, "nested/dir"))
	require.NoError(t, err)
	require.True(t, info.IsDir())
}

func TestCreateFile(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("PROJECT_ROOT", tmp)
	defer os.Unsetenv("PROJECT_ROOT")

	f, err := helpers.CreateFile("dir/file.txt")
	require.NoError(t, err)
	require.NotNil(t, f)
	defer f.Close()

	// Check file exists
	info, err := os.Stat(filepath.Join(tmp, "dir/file.txt"))
	require.NoError(t, err)
	require.False(t, info.IsDir())
}

func TestSaveJSON(t *testing.T) {
	tmp := t.TempDir()
	os.Setenv("PROJECT_ROOT", tmp)
	defer os.Unsetenv("PROJECT_ROOT")

	data := map[string]string{"foo": "bar"}
	err := helpers.SaveJSON("dir/data.json", data)
	require.NoError(t, err)

	// Read back and verify
	content, err := os.ReadFile(filepath.Join(tmp, "dir/data.json"))
	require.NoError(t, err)
	require.Contains(t, string(content), `"foo": "bar"`)
}
