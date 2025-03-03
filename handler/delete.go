package handler

import (
	"encoding/csv"
	"os"
	"strconv"
)

func DeleteRecord(filePath string, id int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var updatedRecords [][]string
	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		recordID, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		if recordID == id {
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}

	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(updatedRecords)
}
