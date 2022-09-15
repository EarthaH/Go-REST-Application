package ls

import (
	"io/fs"
	"io/ioutil"
	"log"
)

func ListDirectory(path string) []fs.FileInfo {
	// if !fs.ValidPath(path) {
	// 	message := fmt.Sprintf("Error: file path \"%s\" is not a valid path.", path)
	// 	log.Fatal(message)
	// }

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
