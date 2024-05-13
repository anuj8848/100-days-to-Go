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

func pl(a ...interface{}){
	fmt.Println(a ...)
}

func main(){
	r:=square{length: 13}
	pl("area: ", r.area())
    pl("perim:", r.peri())
    rp := &r
    pl("area: ", rp.area())
    pl("perim:", rp.peri())
}