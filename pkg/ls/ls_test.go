package ls

import (
	"os"
	"testing"
)

func TestListDirectory(t *testing.T) {
	path, _ := os.UserHomeDir()

	files, err := ListDirectory(path)
	if files == nil || err != nil {
		t.Fatalf("Given home directory path \"ls.ListDirectory\" did not return files.")
	}
}

func TestListDirectoryError(t *testing.T) {
	path := "/no/such/path"

	files, err := ListDirectory(path)
	if files != nil || err == nil {
		t.Fatalf("Given wrong directory path \"ls.ListDirectory\" did not return an error.")
	}
}
