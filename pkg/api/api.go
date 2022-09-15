package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type File struct {
	FileName  string `json:Name`
	FileSize  int    `json:size`
	Attribute string `json:attribute`
}

var files []File

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	files = []File{
		File{FileName: "file1", FileSize: 12, Attribute: "Something"},
		File{FileName: "file2", FileSize: 0, Attribute: "More stuff"},
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/files", getFiles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getFiles")
	json.NewEncoder(w).Encode(files)
}
