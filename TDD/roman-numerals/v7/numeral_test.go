package romanNumerals

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToRoman(t *testing.T) {
	type args struct {
		Arabic int
	}

	tests := []struct {
		description string
		args        args
		want        string
		wantError   error
	}{
		{description: "1 gets converted to I", args: args{Arabic: 1}, want: "I", wantError: nil},
		{description: "2 gets converted to II", args: args{Arabic: 2}, want: "II", wantError: nil},
		{description: "3 gets converted to III", args: args{Arabic: 3}, want: "III", wantError: nil},
		{description: "4 gets converted to IV", args: args{Arabic: 4}, want: "IV", wantError: nil},
		{description: "5 gets converted to V", args: args{Arabic: 5}, want: "V", wantError: nil},
		{"6 gets converted to VI", args{Arabic: 6}, "VI", nil},
		{"7 gets converted to VII", args{Arabic: 7}, "VII", nil},
		{"8 gets converted to VIII", args{Arabic: 8}, "VIII", nil},
		{"9 gets converted to IX", args{Arabic: 9}, "IX", nil},
		{"10 gets converted to X", args{Arabic: 10}, "X", nil},
		{"14 gets converted to XIV", args{Arabic: 14}, "XIV", nil},
		{"20 gets converted to XX", args{Arabic: 20}, "XX", nil},
		{"39 gets converted to XXXIX", args{Arabic: 39}, "XXXIX", nil},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := ConverToRoman(tt.args.Arabic)
			assert.Equal(t, tt.want, got, "should be equal")
		})
	}
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
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
