package handler

import (
	"Abhishek2010DevSingh/expenseTracker/utils"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func WriteCsv(description string, amount float64) {
	file := utils.GetFileToWrite("data.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Cannot get file info: %v", err)
	}

	if stat.Size() == 0 {
		_ = writer.Write([]string{"ID", "Date", "Description", "Amount"})
	}

	_, _ = file.Seek(0, os.SEEK_SET)

	nextID := utils.GetNextID(file)
	record := []string{
		strconv.Itoa(nextID),
		time.Now().Format("2006-01-02"),
		description,
		strconv.Itoa(int(amount)),
	}

	_ = writer.Write(record)

	fmt.Printf("Expense added successfully (ID: %v)", nextID)
}
