package romanNumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConverToRoman(1)

	want := "I"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ConverToRoman(i int) string {
	return "I"
}
