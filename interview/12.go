package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(1994))
}
func intToRoman(num int) string {
	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman []string
	for _, pair := range valueSymbols {
		value, symbol := pair.value, pair.symbol
		for num >= value {
			num -= value
			roman = append(roman, symbol)
		}
		if num == 0 {
			break
		}
	}
	return strings.Join(roman, "")
}
