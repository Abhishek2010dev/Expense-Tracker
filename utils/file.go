package utils

import (
	"log"
	"os"
)

// For writing to file
func GetFileToWrite(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal("Cannot open or create csv file: %s, error: %v", filename, err)
	}
	return file
}
