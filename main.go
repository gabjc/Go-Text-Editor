package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Making a text editor without any framework for a GUI

	//Things to figure out:
	// 1. How to make the data structure for efficient access and edits to a text file
	// ROPE: https://en.wikipedia.org/wiki/Rope_(data_structure)
	// Gap Buffer: https://en.wikipedia.org/wiki/Gap_buffer
	// Piece Table: https://en.wikipedia.org/wiki/Piece_table
	// 2. How to make the text editor look
	// 3. Text cursor behavior
	// 4. Undo/Redo
	// Memento:  https://en.wikipedia.org/wiki/Memento_pattern
	// Command: https://en.wikipedia.org/wiki/Command_pattern

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
