package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	file "icecream.com/chocolate/pkg/dto"
	"icecream.com/chocolate/pkg/ls"
)

func homePage(c *gin.Context) {
	c.Writer.WriteString("Welcome to this Page.")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/files/", getFiles)
	r.GET("/files/new/:filename", makeFile)
	r.GET("/files/replace/:oldname/:newname", renameFile)
	r.GET("/files/delete/:filename", deleteFile)
	r.GET("/files/:filename", readFile)
	r.POST("/files/:filename/save", writeFile)

	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	r.Run("localhost:8080")
}

func getFiles(c *gin.Context) {
	fmt.Println("Endpoint Hit: getFiles")

	rawFiles, err := ls.ListDirectory()

	if err != nil {
		log.Fatal(err)
	}

	files := file.ParseFileNames(rawFiles)

	c.IndentedJSON(http.StatusOK, files)
}

func makeFile(c *gin.Context) {
	fmt.Println("Endpoint Hit: makeFile")

	filename := c.Param("filename")

	err := ls.CreatFile(filename)

	if err != nil {
		log.Fatal(err)
	}
}

func renameFile(c *gin.Context) {
	fmt.Println("Endpoint Hit: renameFile")

	oldname := c.Param("oldname")
	newname := c.Param("newname")

	err := ls.RenameFile(oldname, newname)

	if err != nil {
		log.Fatal(err)
	}
}

func deleteFile(c *gin.Context) {
	fmt.Println("Endpoint Hit: deleteFile")

	filename := c.Param("filename")
	err := ls.DeleteFile(filename)

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(c *gin.Context) {
	fmt.Println("Endpoint Hit: readFile")

	filename := c.Param("filename")
	strlines, err := ls.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, strlines)
}

func writeFile(c *gin.Context) {
	fmt.Println("Endpoint Hit: writeFile")

	var lines file.Line
	filename := c.Param("filename")
	err := c.BindJSON(&lines)

	if err != nil {
		log.Fatal(err)
	}

	ls.WriteFile(filename, lines)
}
