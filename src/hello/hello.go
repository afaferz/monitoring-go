package main

import (
	"fmt"
)

func main() {
	name := "Laurinha"
	// age := 23
	fmt.Println("Hello srta.", name, "Dedezinho te ama")

	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scanf("%d", &command)
	fmt.Println("The command has selected was:", command)
}
