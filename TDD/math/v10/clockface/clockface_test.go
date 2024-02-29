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

func TestSecondHandVector(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		args  args
		point Point
	}{
		{args{simpleTime(0, 0, 30)}, Point{0, -1}},
		{args{simpleTime(0, 0, 45)}, Point{-1, 0}},
	}
	for _, tt := range tests {
		t.Run(testName(tt.args.t), func(t *testing.T) {
			if got := secondHandPoint(tt.args.t); !roughlyEqualPoint(got, tt.point) {
				t.Errorf("secondHandPoint() = %v, want %v", got, tt.point)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point,but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
func testName(t time.Time) string {
	return t.Format("15:04:05")
}
func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}
