package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

// func SecondHand(t time.Time) Point {
// 	p := secondHandPoint(t)
// 	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
// 	p = Point{p.X, -p.Y}                                      // flip
// 	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
// 	return p
// }

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(time time.Time) Point {
	angle := secondsInRadians(time)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}
