package main 
import "fmt"
import "math/rand/v2"

func print(a ...interface{}){
	fmt.Println(a ...)

}

func fill_array() [10]int{
	var a [10] int
	for i:=0;i<10;i++{
		a[i]=rand.IntN(100)
	}
	return a
}

}

func main(){
	arr:=fill_array()
	print(arr)



	
}