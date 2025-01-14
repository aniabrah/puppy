package puppy

import (
	"fmt"

	"github.com/aniabrah/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
func From11() {
	fmt.Println("I am from versin 1.1.0")
}
