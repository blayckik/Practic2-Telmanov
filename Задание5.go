package main

import (
	"fmt"
	"strings"
)

func validateUser(name string, age int, email string) error {
	if name == "" || len(name) >= 50 {
		return fmt.Errorf("некорректное имя")
	}
	
	if age < 18 || age > 100 {
		return fmt.Errorf("некорректный возраст")
	}

	if strings.Index(email, "@") == -1 {
		return fmt.Errorf("email должен содержать символ @")
	}

	return nil
}

func main() {
	err := validateUser("Егор", 18, "egor@test.com")
	if err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Пользователь успешно прошёл валидацию")
	}
}