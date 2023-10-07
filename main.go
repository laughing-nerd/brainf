package main

import (
	utils "github.com/laughing-nerd/brainf/utils"
)

func main() {
	fileContent := utils.CheckValidity()
	code := utils.Construct(fileContent)
	utils.Execute(code)
}
