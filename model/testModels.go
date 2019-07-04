package model

import (
	"fmt"
)

type SalaryCalculator interface {
	Display() string
}

type Employee struct {
	Name     string
	Salary   int
	Currency string
}

func (e Employee) Display() string {
	display := fmt.Sprintf("%s,%s,%d", e.Name, e.Currency, e.Salary)
	return display
}
