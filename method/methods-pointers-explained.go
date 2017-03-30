package main

import(
  "fmt"
  "math"
)

type Verter struct{
  X,Y float64
}

func Abs(v Verter) float64 {
  return math.Sqrt(v.X * v.X + v.Y*v.Y)
}

func Scala(v *Verter,f float64)  {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main()  {
  v:= Verter{3,4}
  Scala(&v,10)
  fmt.Println(Abs(v))
}
