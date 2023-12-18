package romanNumerals

import (
	"strings"
)

func ConverToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			potentialNumber := string([]byte{symbol, nextSymbol})

			value := allRomanNumerals.ValueOf(potentialNumber)

			if value != 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total += allRomanNumerals.ValueOf(string(symbol))
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol byte, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}

func ConverToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = romanNumerals{
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
