package file

import (
	"io/fs"
	"time"
)

type File struct {
	FileName    string      `json:"name"`
	FileSize    int64       `json:"size"`
	Mode        fs.FileMode `json:"mode"`
	FileModTime time.Time   `json:"time"`
	FileType    bool        `json:"directory"`
}

func ParseFiles(rawFiles []fs.FileInfo) []File {
	var files []File

	for _, f := range rawFiles {
		files = append(files, File{FileName: f.Name(), FileSize: f.Size(), Mode: f.Mode(), FileModTime: f.ModTime(), FileType: f.IsDir()})
	}

	return files
}
