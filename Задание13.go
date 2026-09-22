package main

import "fmt"

func main() {
	expenses := make(map[string]float64)

	expenses["Еда"] = 15000.0
	expenses["Транспорт"] = 5000.0
	expenses["Развлечения"] = 3000.0

	expenses["Еда"] += 2000.0

	var total float64

	fmt.Println("Траты по категориям:")
	for category, amount := range expenses {
		fmt.Printf("- %s: %.2f руб.\n", category, amount)
		total += amount
	}

	fmt.Printf("\nИтоговая сумма по всем категориям: %.2f руб.\n", total)
}