package main

import (
	"fmt"
	"os"

	"github.com/2xic-speedrun/brainfuck-interpreter/brainfuck"
)

func main() {
	argsWithProg := os.Args
	if 1 < len(argsWithProg) {
		arguments_without_program := os.Args[1:]
		program := arguments_without_program[0]
		fmt.Println(program)
		output := brainfuck.Run(program)

		fmt.Printf("Program output: %s\n", output)
	} else {
		fmt.Println("Not program input :/")
	}
}
