package expense

import "time"

type Expense struct {
    expenseId  int
    userId     int
    amount     int
    categoryId int
    date       time.Time
    note       string
}
