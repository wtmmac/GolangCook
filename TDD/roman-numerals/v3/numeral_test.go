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
		{
			description: "1 gets converted to I",
			args: args{
				Arabic: 1,
			},
			want:      "I",
			wantError: nil,
		},
		{
			description: "2 gets converted to II",
			args: args{
				Arabic: 2,
			},
			want:      "II",
			wantError: nil,
		},
		{
			description: "3 gets converted to III",
			args: args{
				Arabic: 3,
			},
			want:      "III",
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := ConverToRoman(tt.args.Arabic)
			assert.Equal(t, tt.want, got, "should be equal")
		})
	}
}

func ConverToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
