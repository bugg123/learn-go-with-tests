package numerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var romanNumerals = RomanNumerals{
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

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToRoman(input uint16) string {

	var results strings.Builder

	for _, numeral := range romanNumerals {
		for input >= numeral.Value {
			results.WriteString(numeral.Symbol)
			input -= numeral.Value
		}
	}
	return results.String()
}

func ConvertToArabic(input string) uint16 {
	var total uint16 = 0

	for i := 0; i < len(input); i++ {
		symbol := input[i]

		if CouldBeSubtractive(i, symbol, input) {
			if value := romanNumerals.ValueOf(symbol, input[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += romanNumerals.ValueOf(symbol)
			}
		} else {
			total += romanNumerals.ValueOf(symbol)
		}

	}
	return total
}

func CouldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
