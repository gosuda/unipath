package downloader

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestFile(t *testing.T, path string, content []byte) {
	err := os.WriteFile(path, content, 0644)
	require.NoError(t, err)
}

func createTestDir(t *testing.T, dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	require.NoError(t, err)
}

func cleanup(paths ...string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}
func TestDownloadLocalFile(t *testing.T) {
	createTestFile(t, "a.txt", []byte("hello"))

	defer cleanup("a.txt", "b.txt")

	err := DownloadLocal(context.Background(), "a.txt", "b.txt")
	require.NoError(t, err)
}

func TestDownloadLocalDir(t *testing.T) {
	createTestDir(t, "a")
	createTestFile(t, "a/a.txt", []byte("hello"))

	defer cleanup("a", "b")

	err := DownloadLocal(context.Background(), "a/a.txt", "b/b.txt")
	require.NoError(t, err)
}
