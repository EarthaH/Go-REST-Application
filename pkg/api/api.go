package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	file "icecream.com/chocolate/pkg/dto"
	"icecream.com/chocolate/pkg/ls"
)

type ResponseBody struct {
	Path string `json:"path"`
}

var responseBody ResponseBody

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/files", getFiles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getFiles")

	body, err := ioutil.ReadAll(r.Body)

	if body == nil || err != nil {
		log.Fatal("Error: Requires path to list files.", fmt.Sprintf("Request Error is: %s", err))
	}

	json.Unmarshal(body, &responseBody)

	rawFiles, err := ls.ListDirectory(responseBody.Path)

	if err != nil {
		log.Fatal(err)
	}

	file.SetFiles(rawFiles)

	files := file.GetFiles()

	json.NewEncoder(w).Encode(files)
}
