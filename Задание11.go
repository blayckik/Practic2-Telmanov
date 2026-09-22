package main

import "fmt"

type Product struct {
	Name     string
	Category string
	Price    float64
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
	var filtered []Product

	for _, p := range products {
		if p.Price < maxPrice && p.Category == category {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func main() {

	catalog := []Product{
		{Name: "Хлеб", Category: "Еда", Price: 50.0},
		{Name: "Молоко", Category: "Еда", Price: 90.0},
		{Name: "Ноутбук", Category: "Техника", Price: 60000.0},
		{Name: "Компьютерная мышь", Category: "Техника", Price: 1500.0},
	}

	cheapFood := filterProducts(catalog, 100.0, "Еда")
	fmt.Println("Отфильтрованные товары:")
	for _, p := range cheapFood {
		fmt.Printf("Товар: %s | Категория: %s | Цена: %.2f\n", p.Name, p.Category, p.Price)
	}
}