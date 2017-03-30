// You can declare methods with pointer receivers.
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
package main

import(
  "fmt"
  "math"
)
type Verter struct{
  X,Y float64
}

func (v Verter) Abs() float64{
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Verter) Scala( f float64){
  v.X = v.X * 10
  v.Y = v.Y * 10
}

func main()  {
  v:= Verter{3,4}
  v.Scala(10)
  fmt.Println(v.Abs())
}
