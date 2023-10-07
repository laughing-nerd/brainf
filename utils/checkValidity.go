package utils

import "os"

func CheckValidity() string {
	var filename string
	// Check if a file is provided or not
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		panic("No filename provided...")
	}
	extension := filename[len(filename)-3:]

	// Check if the file extension is .bf or not
	if extension != ".bf" {
		panic("Wrong file provided. Filenames should have a .bf extension")
	}

	// Check if the file can be read or not
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return string(file)
}
