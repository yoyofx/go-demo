package test

import (
	"hello/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MethodForStruct(t *testing.T) {
	employee := model.Employee{
		Name:     "maxzhang",
		Currency: "china",
		Salary:   1985,
	}
	var salary model.SalaryCalculator = employee
	display := salary.Display()
	assert.Equal(t, display, "maxzhang,china,1985")

}
