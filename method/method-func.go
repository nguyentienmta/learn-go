package main
import(
  "fmt"
  "math"
)

type Verter struct{
  X,Y float64
}

func Abs(v Verter) float64{
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
  v:= Verter{3,4}
  fmt.Println(Abs(v))
}
