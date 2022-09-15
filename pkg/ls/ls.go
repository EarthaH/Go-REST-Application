package ls

import (
	"io/fs"
	"io/ioutil"
)

func ListDirectory(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	return files, nil
}
