package romanNumerals

import (
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
	type args struct {
		username string
	}

	tests := []struct {
		description string
		args        args
		want        string
		wantError   error
	}{
		{
			description: "",
			args:        args{},
			want:        "",
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

		})
	}
}

func ConverToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
