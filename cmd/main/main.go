package main

import (
	"log"

	"icecream.com/chocolate/pkg/api"
	"icecream.com/chocolate/pkg/ls"
)

func main() {
	err := ls.CreatHomeDirectory()
	if err != nil {
		log.Fatalln(err)
	}

	api.HandleRequest()
}
