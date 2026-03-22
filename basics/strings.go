package main

import("fmt")

func main(){
	book := "Clean code"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])


	// Strings are immutable
	// book[0] = 112


	// Slice
	fmt.Println(book[4:8])

	fmt.Println(book[:4])

	fmt.Println(book[4:])

	fmt.Println("Uncle bob's book is " + book[0:])


	fmt.Println("It was ½ price!")


	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`


	fmt.Println(poem)
}