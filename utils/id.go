package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

func GetNextID(file *os.File) int {
	_, _ = file.Seek(0, os.SEEK_SET)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil || len(records) < 1 {
		return 1
	}

	var maxID int = 0
	for i := 1; i < len(records); i++ {
		id, err := strconv.Atoi(records[i][0])
		if err == nil && id > maxID {
			maxID = id
		}
	}

	return maxID + 1
}
