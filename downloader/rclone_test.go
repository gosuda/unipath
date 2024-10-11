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

func TestDownloadLocalDir(t *testing.T) {
	createTestDir(t, "a")
	createTestFile(t, "a/a.txt", []byte("hello"))

	defer cleanup("a", "b")

	err := DownloadRclone(context.Background(), "a", "b")
	require.NoError(t, err)
}

func TestDownloadHttp(t *testing.T) {
	defer cleanup("test")

	err := DownloadRclone(context.Background(), ":http,url=https://github.com/pion/webrtc/archive/refs/tags/v4.0.0.zip", "test")
	require.NoError(t, err)
}
