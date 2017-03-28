package main
import(
  "fmt"
  "math"
)

type Verter struct{
  X,Y float64
}

func (v Verter) Abs() float64{
  return math.Sqrt(v.X * v.X + v.Y*v.Y)
}

func main(){
  v := Verter{3,4}
  fmt.Println(v.Abs())
}
