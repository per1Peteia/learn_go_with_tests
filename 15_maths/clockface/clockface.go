package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	p = Point{p.X * 90, p.Y * 90}   // scale
	p = Point{p.X, -p.Y}            // flip (to account for 0/0 top left)
	p = Point{p.X + 150, p.Y + 150} // translate (to account for origin being 150/150)

	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
	// return float64(t.Second()) * (math.Pi / 30)  this will not pass, floats are weird like that
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
