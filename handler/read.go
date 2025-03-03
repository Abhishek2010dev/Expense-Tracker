package handler

import (
	"Abhishek2010DevSingh/expenseTracker/models"
	"Abhishek2010DevSingh/expenseTracker/utils"
	"fmt"
)

func FindExpenseById(id int) (models.Expense, error) {
	recode, err := utils.Read("data.csv")
	if err != nil {
		return models.Expense{}, fmt.Errorf("Error reading file: %v", err)
	}

	for _, expense := range recode {
		if int(expense.ID) == id {
			return expense, nil
		}
	}

	return models.Expense{}, fmt.Errorf("Expense with ID %v not found", id)
}
