package test

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("hello world")

	a = 5
	b = 12

	c = a + b

	fmt.Println(c)

	var i = 0

	for i = 0; i < 5; i++ {
		if d := i % 2; d == 0 {
			fmt.Print(i)
			fmt.Printf(" = Genap\n")
		} else {
			fmt.Print(i)
			fmt.Printf(" = Ganjil\n")
		}
	}

	var lastname string = "vello"
	var firstname string = "calviolia"

	fmt.Printf("halo %s %s \n", firstname, lastname)

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:3]

	fmt.Println(newFruits)

}
