package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func GetNextID(file *os.File) int {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil || len(records) < 1 {
		return 1
	}

	var maxID int = 0
	for i := 1; i < len(records); i++ {
		id, err := strconv.Atoi(records[i][0])
		if err == nil && id > maxID {
			maxID = (id)
		}
	}

	return maxID
}

func GetFileToWrite(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal("Cannot open or create csv file: %s, error: %v", filename, err)
	}
	return file
}
