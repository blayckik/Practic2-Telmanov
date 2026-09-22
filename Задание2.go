package main

import "fmt"
func main(){
var a,b,c float64
fmt.Print("Введите вес основго багажа: ")
fmt.Scan(&a)
fmt.Print("Введите вес ручной клади: ")
fmt.Scan(&b)
fmt.Print("Введите вес доп.ручной клади: ")
fmt.Scan(&c)
result := a+b+c
fmt.Println(result)
}