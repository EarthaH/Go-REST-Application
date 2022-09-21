package file

import (
	"io/fs"
)

type File struct {
	FileName string `json:"name"`
}

type Line []struct {
	FileLine string `json:"line"`
}

func ParseFileNames(rawFiles []fs.FileInfo) []File {
	var files []File

	for _, f := range rawFiles {
		files = append(files, File{FileName: f.Name()})
	}

	return files
}
