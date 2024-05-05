package main

import "fmt"

func main(){

	var a [10]int
	fmt.Println("emp: ",a)

	a[3]=3454
	fmt.Println(a)

	second_arrray :=[4]int{1,2,3,4}
	fmt.Println(second_arrray)

	third_array :=[...]int{1,2,3,4,5,6,7}
	fmt.Println("thrid array: ",third_array)

	fourth_array:=[...]int{2,5:56,439,11:34}
	fmt.Println("fourth array: ",fourth_array)

}