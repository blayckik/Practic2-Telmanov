package main

import "fmt"

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

func calculate(items []InventoryItem) float64 {
	var totalWeight float64

	for _, item := range items {
		totalWeight += item.Weight
	}

	return totalWeight
}

func main() {

	inventory := []InventoryItem{
		{Name: "Меч", Weight: 3.5, IsQuestItem: false},
		{Name: "Щит", Weight: 5.0, IsQuestItem: false},
		{Name: "Зелье здоровья", Weight: 0.5, IsQuestItem: false},
		{Name: "Древний амулет", Weight: 0.2, IsQuestItem: true},
		{Name: "Зачарованная броня", Weight: 12.0, IsQuestItem: false},
	}

	total := calculate(inventory)

	fmt.Println("Предметы в инвентаре:")
	for _, item := range inventory {
		fmt.Printf("- %s (Вес: %.1f кг, Квестовый: %t)\n", item.Name, item.Weight, item.IsQuestItem)
	}

	fmt.Printf("\nОбщий вес инвентаря: %.1f кг\n", total)
}