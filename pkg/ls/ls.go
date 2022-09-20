package ls

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	file "icecream.com/chocolate/pkg/dto"
)

// projectDirectory - for testing. TODO: remove/change for container
var projectDirectory = "./"
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

	fmt.Printf("Successfully created Home directory %s.\n", homeDirectory)

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

func ReadFile(fileName string) ([]string, error) {
	var txtlines []string
	path := filepath.Join(homeDirectory, fileName)
	txtfile, err := os.Open(path)

	if err != nil {
		return txtlines, err
	}

	scanner := bufio.NewScanner(txtfile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	txtfile.Close()
	return txtlines, nil
}

func WriteFile(fileName string, txtlines file.Line) error {
	path := filepath.Join(homeDirectory, fileName)
	txtfile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(txtfile)

	for _, data := range txtlines {
		_, _ = datawriter.WriteString(fmt.Sprintln(data.FileLine))
	}

	datawriter.Flush()
	txtfile.Close()
	return nil
}
