package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	file "icecream.com/chocolate/pkg/dto"
	"icecream.com/chocolate/pkg/ls"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/files/", getFiles)
	r.HandleFunc("/files/new/{filename}", makeFile)
	r.HandleFunc("/files/replace/{oldname}/{newname}", renameFile)
	r.HandleFunc("/files/delete/{filename}", deleteFile)
	r.HandleFunc("/files/{filename}", readFile)
	r.HandleFunc("/files/{filename}/save", writeFile)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getFiles")

	rawFiles, err := ls.ListDirectory()

	if err != nil {
		log.Fatal(err)
	}

	files := file.ParseFileNames(rawFiles)

	json.NewEncoder(w).Encode(files)
}

func makeFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: makeFile")

	vars := mux.Vars(r)

	err := ls.CreatFile(vars["filename"])

	if err != nil {
		log.Fatal(err)
	}
}

func renameFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: renameFile")

	vars := mux.Vars(r)

	err := ls.RenameFile(vars["oldname"], vars["newname"])

	if err != nil {
		log.Fatal(err)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteFile")

	vars := mux.Vars(r)
	err := ls.DeleteFile(vars["filename"])

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: readFile")

	vars := mux.Vars(r)
	strlines, err := ls.ReadFile(vars["filename"])

	if err != nil {
		log.Fatal(err)
	}

	jsonlines, _ := json.Marshal(strlines)

	fmt.Println(string(jsonlines))
}

func writeFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: writeFile")

	var lines file.Line
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &lines)

	ls.WriteFile(vars["filename"], lines)
}
