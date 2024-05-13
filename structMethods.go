package main 
import "fmt"

type square struct{
	length int
}
func (r *square) area() int{
	return r.length*r.length
}

func (r square) peri() int{
	return 4*r.length
}

func main(){
	r:=square{length: 13}
	fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.peri())
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.peri())
}