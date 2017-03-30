package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type Verter struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Verter) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	f := MyFloat(-10)
	fmt.Println(f)
	v := Verter{3, 4}
	fmt.Println(v.Abs())
}
