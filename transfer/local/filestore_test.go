package local

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gosuda.org/unipath/unipath"
)

func TestClient_ReadWrite(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("failed to create tmp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	client := NewClient(tmpDir)

	tests := []struct {
		name     string
		filePath string
		content  string
	}{
		{
			name:     "simple file write and read",
			filePath: "testfile.txt",
			content:  "Hello, World!",
		},
		{
			name:     "nested directory write and read",
			filePath: "subdir/inner.txt",
			content:  "Nested content!",
		},
		{
			name:     "empty file write and read",
			filePath: "empty.txt",
			content:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toPath := &unipath.UniPath{
				Protocol: unipath.Local,
				Path:     tt.filePath,
			}
			reader := io.NopCloser(strings.NewReader(tt.content))

			// Write
			err := client.Write(context.Background(), toPath, reader)
			require.NoError(t, err, "Write failed")

			fullPath := filepath.Join(tmpDir, tt.filePath)

			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				t.Fatalf("file %s does not exist", fullPath)
			}

			// Read
			fromPath := &unipath.UniPath{
				Protocol: unipath.Local,
				Path:     tt.filePath,
			}
			reader, err = client.Read(context.Background(), fromPath)
			require.NoError(t, err, "Read failed")

			defer reader.Close()

			buf := new(strings.Builder)
			_, err = io.Copy(buf, reader)
			require.NoError(t, err, "failed to copy data from reader")

			require.Equal(t, tt.content, buf.String(), "content mismatch")
		})
	}
}
