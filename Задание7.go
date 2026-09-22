package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func calculate(employees []Employee) (float64, float64) {
	if len(employees) == 0 {
		return 0, 0
	}

	var total float64
	for _, emp := range employees {
		total += emp.Salary
	}

	avg := total / float64(len(employees))
	return total, avg
}

func main() {
	staff := []Employee{
		{ID: 1, Name: "Егор", Position: "Разработчик", Salary: 120000.0},
		{ID: 2, Name: "Диитрий", Position: "Аналитик", Salary: 95000.0},
		{ID: 3, Name: "Ярослав", Position: "Тестировщик", Salary: 85000.0},
	}

	total, avg := calculate(staff)

	fmt.Printf("Общий фонд оплаты труда: %.2f\n", total)
	fmt.Printf("Средняя зарплата: %.2f\n", avg)
}