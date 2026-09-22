package main

import (
	"fmt"
	"strconv"
)

const (
	bin = "bin"
	dec = "dec"
	hex = "hex"
)


func convertBase(numStr string, fromBase string, toBase string) string {
	var baseFrom int
	switch fromBase {
	case bin:
		baseFrom = 2
	case dec:
		baseFrom = 10
	case hex:
		baseFrom = 16
	}

	val, _ := strconv.ParseInt(numStr, baseFrom, 64)

	var baseTo int
	switch toBase {
	case bin:
		baseTo = 2
	case dec:
		baseTo = 10
	case hex:
		baseTo = 16
	}

	return strconv.FormatInt(val, baseTo)
}

func main() {
	var num int64 = 255
	
	binStr := strconv.FormatInt(num, 2)
	hexStr := strconv.FormatInt(num, 16)

	fmt.Printf("Десятичное: %d\n", num)
	fmt.Printf("Двоичное: %s\n", binStr)
	fmt.Printf("Шестнадцатеричное: %s\n", hexStr)
}
