package main
import "fmt"

func subint() func() int{
	func() int{
		func() int{
			func()int{
				func() int{
				 i:=0
					return func() int{
						return i--
					}
				}
			}
		}	
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