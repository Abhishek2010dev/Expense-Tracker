package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func GetNextID(file *os.File) uint {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil || len(records) < 1 {
		return 1
	}

	var maxID uint = 0
	for i := 1; i < len(records); i++ {
		id, err := strconv.ParseUint(records[i][0], 10, 32)
		if err == nil && uint(id) > maxID {
			maxID = uint(id)
		}
	}

	return maxID
}
