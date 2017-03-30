// functions with a pointer argument must take a pointer
//methods with pointer receivers take either a value or a pointer as the receiver when they are called

package main

import "fmt"

// new type
type Verter struct {
	X, Y float64
}

func ScaleFunc(v *Verter, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Verter) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Verter{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Verter{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
