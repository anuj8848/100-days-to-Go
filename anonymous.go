package main
import "fmt"

func subint() func() int{
	i:=0
	return func() int{
		i--
		return i
		
	}
}

func main(){
	nxtint:=subint()

	fmt.Println(nxtint())
	fmt.Println(nxtint())

	fmt.Println(nxtint())

	nextint:=subint()
	fmt.Println((nextint()))

}