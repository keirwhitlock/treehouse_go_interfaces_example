package main

import "pets/pets"

type FourLegged interface {
	Walk()
	Sit()
}

func Demo(animal FourLegged) {
	animal.Walk()
	animal.Sit()
}

func main() {
	dog := pets.Dog{"Fido", "Terrier"}
	cat := pets.Cat{"Fluffy", "Siamese"}
	Demo(dog)
	Demo(cat)
}
