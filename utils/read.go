package utils

import (
	"Abhishek2010DevSingh/expenseTracker/models"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func Read(filename string) ([]models.Expense, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []models.Expense
	for i, line := range lines {
		if i == 0 {
			continue
		}

		id, _ := strconv.ParseUint(line[0], 10, 32)
		date, _ := time.Parse("2006-01-02", line[1])
		amount, _ := strconv.ParseFloat(line[3], 64)

		records = append(records, models.Expense{
			ID:          uint(id),
			Date:        date,
			Description: line[2],
			Amount:      amount,
		})
	}

	return records, nil
}
