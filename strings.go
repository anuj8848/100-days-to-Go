package main
import (
	"fmt"
)

func main(){
	const s=";kdjfadlkfj"

	fmt.Println("len: ",len(s))

	for i:=0;i<len(s);i++{
		fmt.Println(s[i])
		fmt.Printf("%x",s[i])

	}
}