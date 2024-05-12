package main
import "fmt"

type animal struct{
	name string
	legs int
	sound string
}

func newAnimal(name string,legs int, sound string) *animal{

	a:=animal{name: name}
	a.legs=legs
	a.sound=sound
	return &a
}

func pl(a ...interface{}){
	fmt.Println(a ...)
}

func main(){

	pl(animal{"cow",4,"baaaa"})

	dog:=animal{"dog",4,"vau vau"}

	pl(dog.legs)


}