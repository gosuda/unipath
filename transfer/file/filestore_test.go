package file

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gosuda.org/unipath/unipath"
)

func TestClient_ReadWrite(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("failed to create tmp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testFilePath := "testfile.txt"
	testContent := "Hello, World!"

	client := NewClient(tmpDir)

	t.Run("Write", func(t *testing.T) {
		toPath := &unipath.UniPath{
			Protocol: unipath.Local,
			Path:     testFilePath,
		}
		reader := io.NopCloser(strings.NewReader(testContent))

		err := client.Write(context.Background(), toPath, reader)
		if err != nil {
			t.Fatalf("Write failed: %v", err)
		}

		if _, err := os.Stat(filepath.Join(tmpDir, testFilePath)); os.IsNotExist(err) {
			t.Fatalf("file %s does not exist", testFilePath)
		}
	})

	t.Run("Read", func(t *testing.T) {
		fromPath := &unipath.UniPath{
			Protocol: unipath.Local,
			Path:     testFilePath,
		}
		reader, err := client.Read(context.Background(), fromPath)
		if err != nil {
			t.Fatalf("Read failed: %v", err)
		}
		defer reader.Close()

		buf := new(strings.Builder)
		_, err = io.Copy(buf, reader)
		if err != nil {
			t.Fatalf("failed to copy data from reader: %v", err)
		}

		if buf.String() != testContent {
			t.Errorf("expected %s but got %s", testContent, buf.String())
		}
	})
}
