package main

import (
	"os"

	compiler "github.com/BrunoAgst/8bit-compiler/compilerHandling"
)

func main() {

	errors := []string{}

	file, err := os.Open("./file.asm")
	if err != nil {
		panic(err)
	}

	code := compiler.RemoveCommentsAndScape(file, &errors)

	file.Close()

	fileCreate, err := os.Create("file.bin")
	if err != nil {
		panic(err)
	}

	defer fileCreate.Close()

	compiler.Execute(&code, fileCreate, &errors)
}
