package v7

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

func ConvertToArabic(roman string) int {
	total := 0

	for _, numaral := range allRomanNumarals {
		for strings.HasPrefix(roman, numaral.Symbol) {
			total += numaral.Value
			roman = strings.TrimPrefix(roman, numaral.Symbol)
		}
	}

	return total
}
