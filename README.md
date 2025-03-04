# 💰 Expense Tracker CLI

A simple expense tracker CLI tool built using Go and Cobra.

## 📦 Installation

Clone the repository and navigate into the project directory:

```sh
git clone https://github.com/yourusername/expense-tracker.git
cd expense-tracker
```

Build the project:

```sh
go build -o expense-tracker
```

Or run it directly without building:

```sh
go run .
```

## 🚀 Usage

### ➕ Add an expense
```sh
go run . add --description "Lunch" --amount 20
```
Output:
```
✅ Expense added successfully (ID: 1)
```

```sh
go run . add --description "Dinner" --amount 10
```
Output:
```
✅ Expense added successfully (ID: 2)
```

### 📋 List all expenses
```sh
go run . list
```
Output:
```
📊 ID  Date       Description  Amount
1   2024-08-06  Lunch        $20
2   2024-08-06  Dinner       $10
```

### 📈 View expense summary
```sh
go run . summary
```
Output:
```
💵 Total expenses: $30
```

### ❌ Delete an expense
```sh
go run . delete --id 2
```
Output:
```
🗑️ Expense deleted successfully
```

```sh
go run . summary
```
Output:
```
💵 Total expenses: $20
```

### 📅 Monthly summary
```sh
go run . summary --month 8
```
Output:
```
🗓️ Total expenses for August: $20
```

## 📚 Reference
For more ideas and features, check out the [Expense Tracker Roadmap](https://roadmap.sh/projects/expense-tracker).

## 📜 License
📝 MIT License

