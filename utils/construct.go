package utils

import "unicode"

func Construct(fileContent string) string {
	var loopChecker int
	var code string
	for _, v := range fileContent {
		if v != ' ' && v != '\n' && unicode.IsLetter(v) == false {
			if v == '[' {
				loopChecker += 1
			}
			if v == ']' {
				loopChecker -= 1
			}
			code = code + string(v)
		}
	}
	if loopChecker != 0 {
		panic("Couldn't validate loops")
	}
	return code
}
