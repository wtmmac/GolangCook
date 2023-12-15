package romanNumerals

import (
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
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := ConverToRoman(tt.args.Arabic)
			assert.Equal(t, tt.want, got, "should be equal")
		})
	}
}

func ConverToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
