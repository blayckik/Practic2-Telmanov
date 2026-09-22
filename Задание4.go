package main

import "fmt"

func countVotes(votes []string) {
	var anna int
	var boris int
	var viktor int

	total := len(votes)

	for _, name := range votes {
		if name == "Анна" {
			anna++
		} else if name == "Борис" {
			boris++
		} else if name == "Виктор" {
			viktor++
		}
	}

	anna2 := (float64(anna) / float64(total)) * 100
	boris2 := (float64(boris) / float64(total)) * 100
	viktor2 := (float64(viktor) / float64(total)) * 100

	fmt.Printf("Всего голосов: %d\n", total)
	fmt.Printf("Анна: %d (%.2f%%)\n", anna, anna2)
	fmt.Printf("Борис: %d (%.2f%%)\n", boris, boris2)
	fmt.Printf("Виктор: %d (%.2f%%)\n", viktor, viktor2)
}

func main() {
	votes := []string{"Анна", "Борис", "Анна", "Виктор", "Анна", "Борис", "Виктор", "Анна"}
	countVotes(votes)
}