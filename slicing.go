package main 
import(
	"fmt"
	"slices"
)

func pl(a ...interface{}){
	fmt.Println(a ...)
}

func main(){
	var s[]string
	pl("uninitialized slice:", s==nil, len(s)==0)


	s=make([]string, 3)
	pl("emp: ", s,"len: ", len(s), "caps: ",cap(s))


	s[0]="a"
	s[1]="b"
	s[2]="c"
	pl("emp: ", s,"len: ", len(s), "caps: ",cap(s))
}