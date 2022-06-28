package brainfuck

import "fmt"

func Run(program string) string {
	output := ""

	instruction_pointer := 0
	data_pointer := 0

	data := make([]byte, 10)
	bytes := []byte(program)

	for instruction_pointer < len(bytes) {
		char := string(bytes[instruction_pointer])

		if char == ">" {
			data_pointer++
		} else if char == "<" {
			data_pointer--
		} else if char == "+" {
			data[data_pointer] = data[data_pointer] + 1
		} else if char == "-" {
			data[data_pointer] = data[data_pointer] - 1
		} else if char == "." {
			output += string(data[data_pointer])
		} else if char == "," {
			fmt.Println("Enter input char (we only copy one char... don't try to cheat): ")
			var second string
			fmt.Scanln(&second)
			data[data_pointer] = second[0]
		} else if char == "[" {
			skip := data[data_pointer] == 0
			if skip {
				instruction_pointer = move_tape(
					bytes,
					"]",
					instruction_pointer,
					1,
				)
				continue
			}
		} else if char == "]" {
			rewind := data[data_pointer] != 0
			if rewind {
				instruction_pointer = move_tape(
					bytes,
					"[",
					instruction_pointer,
					-1,
				)
				continue
			}
		} else {
			panic("Unknown char ??? ")
		}
		instruction_pointer++
	}

	return output
}

func move_tape(data []byte, target string, index int, direction int) int {
	for 0 <= index && index < len(data) && string(data[index]) != target {
		index += direction
	}
	return index
}
