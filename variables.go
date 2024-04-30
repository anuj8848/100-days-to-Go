package main


import (
	"math"

	"fmt"
)


var printo ="printer.Println"

func printer(a ...interface{}){
	fmt.Println(a ...)
}

func main(){
	var a,s,d,e int=1,2,3,int(math.Round(math.Pi))
	printer(a,s,d,e)

	game :="catandmo" // variable declared and initialized at same time

	printer(game)


}