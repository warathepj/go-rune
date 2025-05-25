package main

import "fmt"

func main() {

	// A rune literal using a simple English character
	var myRune1 rune = 'A'
	fmt.Printf("myRune1: %c (Unicode: %U)\n", myRune1, myRune1)
	// %c prints the character, %U prints the Unicode code point

	// A rune literal using a non-English character
	var myRune2 rune = '世'
	fmt.Printf("myRune2: %c (Unicode: %U)\n", myRune2, myRune2)

	// A rune literal using an emoji
	var myRune3 rune = '😊'
	fmt.Printf("myRune3: %c (Unicode: %U)\n", myRune3, myRune3)

	// When you loop over a string using 'range', you get runes
	myString := "Hello 😊 世界"
	fmt.Println("\nLooping through a string with range:")
	for index, character := range myString {
		fmt.Printf("At byte index %d, found rune: %c (Unicode: %U)\n", index, character, character)
	}
}
