package logger

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var logfile = "out.log"

func Init() {
	file, _ := os.Create(logfile)
	logger = log.New(file, "File Manager: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("Application Started.")
}

func Info(message string) {
	infomessage := fmt.Sprintf("INFO: %s", message)
	logger.Println(infomessage)
}

func Warning(message string) {
	warnmessage := fmt.Sprintf("WARNING: %s", message)
	logger.Println(warnmessage)
}

func Error(err error) {
	errmessage := fmt.Sprintf("ERROR: %s", err)
	logger.Println(errmessage)
}

func FatalError(err error) {
	errmessage := fmt.Sprintf("FATAL ERROR: %s", err)
	logger.Fatalln(errmessage)
}
