package main

import (
	compiler "github.com/BrunoAgst/8bit-compiler/compilerHandling"
	"github.com/BrunoAgst/8bit-compiler/errorsHandling"

	"fmt"
	"os"
	"strings"
)

func main() {

	var code []string
	errors := []string{}

	file, err := os.Open("./file.asm")
	if err != nil {
		panic(err)
	}

	code = compiler.RemoveCommentsAndScape(file, &errors)

	file.Close()

	fileCreate, err := os.Create("file.bin")
	if err != nil {
		panic(err)
	}

	defer fileCreate.Close()

	codeLen := len(code)
	errorsLen := 0

	for {

		if errorsLen > 0 {
			errorsHandling.PrintError(&errors, errorsLen)
			break
		}

		if codeLen == 0 {
			break
		}
		instruction := strings.Join(code[:3], "")
		compiler.SwitchInstructions(instruction, &code, fileCreate, &errors)

		codeLen = len(code)
		errorsLen = len(errors)

	}
	if errorsLen == 0 {
		fmt.Println("[OK] - Compiled Success")
	}
}
