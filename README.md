# Go Rune Demo

A simple demonstration project to understand runes in Go.

## What is a Rune in Go?

In Go, a `rune` is a built-in type that represents a single Unicode code point. It's an alias for the `int32` type, meaning it's a 32-bit integer that can represent any character in the Unicode standard.

While some languages represent characters with a simple byte (which can only handle 256 different characters), Go's `rune` type can represent any character from any language in the world, including emojis and special symbols.

## Why Use Runes?

A simple byte (like the `byte` type in Go, which is an alias for `uint8`) is not enough to represent all characters in the world. Some characters need more than one byte when encoded in UTF-8 (which is how Go stores strings internally). The `rune` type ensures you're dealing with a full character, regardless of how many bytes it takes up in the string.

## Code Example

The demo code in this project illustrates how runes work:

```go
package main

import "fmt"

func main() {
    // A rune literal using a simple English character
    var myRune1 rune = 'A'
    fmt.Printf("myRune1: %c (Unicode: %U)\n", myRune1, myRune1)

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
        fmt.Printf("At byte index %d, found rune: %c (Unicode: %U)\n",
                   index, character, character)
    }
}
```

## Explanation

1. We declare variables of type `rune` using single quotes `' '`.
2. We use `fmt.Printf` with the `%c` format specifier to print the character that the rune represents, and `%U` to print its Unicode code point.
3. When we loop through a string using `for range`, Go automatically gives us each character as a `rune`. The index represents the byte position where that rune starts in the string.
4. Notice how non-ASCII characters and emojis start at indices that are more than 1 byte apart, because they take up multiple bytes in the string's UTF-8 encoding.

## Running the Demo

To run this demo:

1. Make sure you have Go installed on your system
2. Clone this repository
3. Navigate to the project directory
4. Run the following command:

```
go run main.go
```

## Expected Output

```
myRune1: A (Unicode: U+0041)
myRune2: 世 (Unicode: U+4E16)
myRune3: 😊 (Unicode: U+1F60A)

Looping through a string with range:
At byte index 0, found rune: H (Unicode: U+0048)
At byte index 1, found rune: e (Unicode: U+0065)
At byte index 2, found rune: l (Unicode: U+006C)
At byte index 3, found rune: l (Unicode: U+006C)
At byte index 4, found rune: o (Unicode: U+006F)
At byte index 5, found rune:   (Unicode: U+0020)
At byte index 6, found rune: 😊 (Unicode: U+1F60A)
At byte index 10, found rune:   (Unicode: U+0020)
At byte index 11, found rune: 世 (Unicode: U+4E16)
At byte index 14, found rune: 界 (Unicode: U+754C)
```

Notice how the emoji (😊) takes 4 bytes and the Chinese characters (世, 界) each take 3 bytes in UTF-8 encoding.
