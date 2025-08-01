package v5

import "strings"

type RomanNumaral struct {
	Value  int
	Symbol string
}

var allRomanNumarals = []RomanNumaral{
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

// ConvertToRoman change arabic to roman
func ConvertToRoman(arbic int) string {

	var result strings.Builder

	for _, romanNumaral := range allRomanNumarals {
		for arbic >= romanNumaral.Value {
			arbic -= romanNumaral.Value
			result.WriteString(romanNumaral.Symbol)
		}
	}
	return result.String()
}
