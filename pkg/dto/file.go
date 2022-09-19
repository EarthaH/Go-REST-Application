package file

import (
	"io/fs"
)

type File struct {
	FileName    string `json:"name"`
	FileContent string `json:"content"`
}

func ParseFiles(rawFiles []fs.FileInfo) []File {
	var files []File

	for _, f := range rawFiles {
		files = append(files, File{FileName: f.Name()})
	}

	return files
}
