package cmds

import "time"

type Expense struct {
	id          int
	description string
	amount      float64
	date        time.Time
}

func New(description string, amount float64) Expense {
	expense := Expense{
		id:          0,
		description: description,
		amount:      amount,
		date:        time.Now(),
	}
	return expense
}
