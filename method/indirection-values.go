package main

import (
	"fmt"
	"math"
)

type Verter struct {
	X, Y float64
}

func AbsFunc(v Verter) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Verter) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Verter{3, 4}
	fmt.Println(AbsFunc(v))
	fmt.Println(v.Abs())

	p := &Verter{4, 3}
	fmt.Println(AbsFunc(*p))
	fmt.Println(p.Abs())
}
