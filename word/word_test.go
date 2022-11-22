package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty string", args{""}, true},
		{"aa", args{"aa"}, true},
		{"ab", args{"ab"}, false},
		{"奶牛产牛奶", args{"奶牛产牛奶"}, true},
		{"A man, a plan, a canal: Panama", args{"A man, a plan, a canal: Panama"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
