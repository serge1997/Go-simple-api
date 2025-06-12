package services

import (
	"errors"
	"log"
	"os"
)

var ErroToOpenLogFile error = errors.New("erro when trying top open log file")

func openLoFile() (*os.File, error) {
	file, err := os.OpenFile("storage/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, ErroToOpenLogFile
	}
	return file, nil
}
func WriteFatal(content string) {
	file, err := openLoFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Fatal("Fatal error ocurred: ", content)
}
func Write(content string) {
	file, err := openLoFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print(content)
}
