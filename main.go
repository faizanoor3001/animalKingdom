package main

import (
	"fmt"

	"github.com/faizanoor3001/puppy"
)

func main() {
	fmt.Println("this is the main package!")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.WhenGrownUp())
	fmt.Println(puppy.WhenGrownUp() + "Bhau !! Bhau!!!")
	fmt.Println(puppy.WhenGrownUp() + "Meow !! Oopss !! Am I a cat ?")
}
