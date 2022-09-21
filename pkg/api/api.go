package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	file "icecream.com/chocolate/pkg/dto"
	"icecream.com/chocolate/pkg/logger"
	"icecream.com/chocolate/pkg/ls"
)

func homePage(c *gin.Context) {
	c.Writer.WriteString("Welcome to this Page.")
	logger.Info("Endpoint Hit: homePage")
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

	r.Run("localhost:8080")
}

func getFiles(c *gin.Context) {
	logger.Info("Endpoint Hit: getFiles")

	rawFiles, err := ls.ListDirectory()

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	files := file.ParseFileNames(rawFiles)

	c.IndentedJSON(http.StatusOK, files)
}

func makeFile(c *gin.Context) {
	logger.Info("Endpoint Hit: makeFile")

	filename := c.Param("filename")

	err := ls.CreateFile(filename)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

func renameFile(c *gin.Context) {
	logger.Info("Endpoint Hit: renameFile")

	oldname := c.Param("oldname")
	newname := c.Param("newname")

	err := ls.RenameFile(oldname, newname)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

func deleteFile(c *gin.Context) {
	logger.Info("Endpoint Hit: deleteFile")

	filename := c.Param("filename")
	err := ls.DeleteFile(filename)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}
}

func readFile(c *gin.Context) {
	logger.Info("Endpoint Hit: readFile")

	filename := c.Param("filename")
	strlines, err := ls.ReadFile(filename)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.IndentedJSON(http.StatusOK, strlines)
}

func writeFile(c *gin.Context) {
	logger.Info("Endpoint Hit: writeFile")

	var lines file.Line
	filename := c.Param("filename")
	err := c.BindJSON(&lines)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	ls.WriteFile(filename, lines)
}
