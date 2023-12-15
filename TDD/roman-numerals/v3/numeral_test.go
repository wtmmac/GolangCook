package romanNumerals

import (
	"errors"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gets converted to I", func(t *testing.T) {
		got := ConverToRoman(1)

		want := "I"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("2 gets converted to II", func(t *testing.T) {
		got := ConverToRoman(2)

		want := "II"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func TestConvertToRoman(t *testing.T) {
	tests := []struct {
		description string
		wantError   error
	}{
		{
			description: "",
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := ConvertToRoman()
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}
		})
	}
}

func ConverToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
