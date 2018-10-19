package main

import (
	"fmt"

	"github.com/Sach97/exported-types-and-methods-go/animals"
)

func main() {

	dog := animals.Dog{
		Name:         "Rex",
		BarkStrength: 3,
		Age:          7,
		Wouaf:        "wouaf",
	}

	wouaf := dog.Bark()

	fmt.Printf(dog)
	fmt.Printf(wouaf)
}
