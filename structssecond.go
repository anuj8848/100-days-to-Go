package main
import "fmt"

type animal struct{
	name string
	legs int
	sound string
}

func newAnimal(name string,legs int, sound string) *animal{

	a:=animal{name: name,legs:legs,sound:sound}
	// a.legs=legs
	// a.sound=sound
	return &a
}

func pl(a ...interface{}){
	fmt.Println(a ...)
}

func main(){

	pl(animal{"cow",4,"vau vau"})

	// dog:=animal()
	// dog.name="dog"
	// dog.legs=4
	// dog.sound="vau vau"

	// pl(dog)

}