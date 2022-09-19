package ls

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// projectDirectory - for testing. TODO: remove/change for container
var projectDirectory = "/Users/eartha/SideProjects/FileManager/Go-REST-Application"
var homeDirectory = filepath.Join(projectDirectory, "Icecream")

func ListDirectory() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(homeDirectory)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func CreatHomeDirectory() error {
	_, err := os.Stat(homeDirectory)

	if os.IsNotExist(err) {
		errDir := os.Mkdir(homeDirectory, 0755)
		if errDir != nil {
			return errDir
		}
	}

	fmt.Printf("Successfully created Home directory %s.", homeDirectory)

	return nil
}

func CreatFile(fileName string) error {
	if homeDirectory == "" {
		return errors.New("Home Directory not set")
	}

	newFile, err := os.Create(filepath.Join(homeDirectory, fileName))

	if err != nil {
		return err
	}

	fmt.Printf("New file %s created in %s.\n", newFile.Name(), homeDirectory)
	newFile.Close()

	return nil
}

func RenameFile(oldName string, newName string) error {
	err := os.Rename(filepath.Join(homeDirectory, oldName), filepath.Join(homeDirectory, newName))

	return err
}

func DeleteFile(fileName string) error {
	err := os.Remove(filepath.Join(homeDirectory, fileName))

	return err
}
