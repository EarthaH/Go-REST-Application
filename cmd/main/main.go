package main

import (
	"fmt"
	"os"

	//"icecream.com/chocolate/pkg/api"
	file "icecream.com/chocolate/pkg/dto"
	"icecream.com/chocolate/pkg/ls"
)

func main() {
	home, _ := os.UserHomeDir()
	//api.HandleRequest()
	rawFiles := ls.ListDirectory(home)
	file.SetFiles(rawFiles)
	files := file.GetFiles()
	for _, f := range files {
		fmt.Println(f)
	}
}
