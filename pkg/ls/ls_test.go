package ls

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

var dirname = "Icecream"
var filename = "test.txt"
var filename2 = "test2.txt"

func TestWorkingDirectoryCreated(t *testing.T) {
	reset("")

	if err := CreateHomeDirectory(); err != nil {
		t.Failed()
	}
}

func TestListEmptyDirectory(t *testing.T) {
	reset("")
	CreateHomeDirectory()

	if files, err := ListDirectory(); files != nil || err != nil {
		t.Failed()
	}
	reset("")
}

func TestListDirectory(t *testing.T) {
	reset("")
	CreateHomeDirectory()
	CreateFile(filename)
	CreateFile(filename2)

	files, err := ListDirectory()
	if files == nil || err != nil {
		t.Failed()
	}
	if files[0].Name() != filename || files[1].Name() != filename2 {
		t.Failed()
	}
	reset("")
}

func TestCreateFile(t *testing.T) {
	reset("")
	CreateHomeDirectory()

	if err := CreateFile(filename); err != nil {
		t.Failed()
	}
	if _, err := os.Stat(filepath.Join(dirname, filename)); os.IsNotExist(err) {
		t.Failed()
	}
	reset("")
}

func TestCreateFileFail(t *testing.T) {
	reset("")
	CreateHomeDirectory()
	_, _ = os.Create(filepath.Join(dirname, filename))

	if err := CreateFile(filename); err == nil {
		t.Failed()
	}
	reset("")
}

func reset(fname string) {
	if fname == "" {
		if err := os.RemoveAll(dirname); err != nil {
			log.Println(err)
		}
	} else {
		if err := os.Remove(filepath.Join(dirname, fname)); err != nil {
			log.Fatalln(err)
		}
	}
}
