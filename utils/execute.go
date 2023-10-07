package utils

import (
	"fmt"
	"math"
)

func Execute(code string) {
	var tape [30000]uint8
	var loopMemory []int
	var iterator, pointer int

	for iterator < len(code) {
		if string(code[iterator]) == "+" {
			tape[pointer]++
		} else if string(code[iterator]) == "-" {
			tape[pointer]--
		} else if string(code[iterator]) == ">" {
			if pointer == math.MaxUint8 {
				pointer = 0
			} else {
				pointer++
			}
		} else if string(code[iterator]) == "<" {
			if pointer == 0 {
				pointer = 255
			} else {
				pointer--

			}
		} else if string(code[iterator]) == "." {
			fmt.Print(string(tape[pointer]))
		} else if string(code[iterator]) == "," {
			var sc string
			fmt.Scanln(&sc)
			tape[pointer] = sc[0]
		} else if string(code[iterator]) == "[" {
			loopMemory = append(loopMemory, iterator)
		} else if string(code[iterator]) == "]" {
			if tape[pointer] != 0 {
				iterator = loopMemory[len(loopMemory)-1]
				continue
			}

			if len(loopMemory) == 1 {
				loopMemory = []int{}
			} else {
				loopMemory = loopMemory[:len(loopMemory)-1]
			}

		}
		iterator++
	}
}
