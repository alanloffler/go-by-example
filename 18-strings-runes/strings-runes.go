
// A Go string is a read-only slice of bytes. A string is a containers of text encoded in UTF-8.
// In Go, a character is called a rune - it’s an integer that represents a Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "Alan"

	// Strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("len:", len(s))

	//Indexing into a string produces the raw byte values at each index
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Count how many runes are in the string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Same previous iteration with utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRunInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// Values enclosed in single quotes are rune literals
func examineRune(r rune) {
	if r == 'a' {
		fmt.Println("found a")
	} else if r == 'A' {
		fmt.Println("found A")
	}
}
