package main

import (
	"icecream.com/chocolate/pkg/api"
	"icecream.com/chocolate/pkg/logger"
	"icecream.com/chocolate/pkg/ls"
)

func main() {
	logger.Init()
	err := ls.CreateHomeDirectory()
	if err != nil {
		logger.FatalError(err)
	}

	api.HandleRequest()
}
