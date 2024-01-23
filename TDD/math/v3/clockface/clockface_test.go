package clockface

import (
	"math"
	"testing"
	"time"
)

func Test_secondsInRadians(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		args args
		want float64
	}{
		{args{simpleTime(0, 0, 30)}, math.Pi},
		{args{t: simpleTime(0, 0, 0)}, 0},
		{args{t: simpleTime(0, 0, 45)}, math.Pi / 2 * 3},
		{args{t: simpleTime(0, 0, 7)}, math.Pi / 30 * 7},
	}
	for _, tt := range tests {
		t.Run(testName(tt.args.t), func(t *testing.T) {
			if got := secondsInRadians(tt.args.t); got != tt.want {
				t.Errorf("secondsInRadians() = %v, want %v", got, tt.want)
			}
		})
	}
}
func testName(t time.Time) string {
	return t.Format("15:04:05")
}
func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}
