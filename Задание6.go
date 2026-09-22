package main

import "fmt"

func UniqueTags(posts [][]string) map[string]bool {
	uniqueTags := make(map[string]bool)

	for _, post := range posts {
		for _, tag := range post {
			uniqueTags[tag] = true
		}
	}

	return uniqueTags
}

func main() {
	posts := [][]string{
		{"go", "backend"},
		{"git", "go", "tools"},
	}

	result := UniqueTags(posts)

	fmt.Println("Уникальные теги:")
	for tag := range result {
		fmt.Println(tag)
	}
}