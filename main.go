package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	defer func() {
		if r := recover(); r != nil {

			str := fmt.Sprintf("%v", r)
			exceptions(str)
			if _, err := os.Stat("./file.bin"); os.IsNotExist(err) {
				fmt.Println("[ERROR] - Compiled Error")
			} else {
				os.Remove("file.bin")
				fmt.Println("[ERROR] - Compiled Error")
			}
		}
	}()

	file, err := os.Open("./file.asm8")
	var code []string

	if err != nil {
		fmt.Println("[ERROR] - Read File", err)
		panic(0)
	}
	code = removeCommentsAndScape(file)

	defer file.Close()

	fileCreate, err := os.Create("file.bin")

	if err != nil {
		fmt.Println("[ERROR] - Create File", err)
		panic(0)
	}

	defer fileCreate.Close()

	codeLen := len(code)

	for {

		if codeLen == 0 {
			break
		}
		instruction := strings.Join(code[:3], "")
		switchInstructions(instruction, &code, fileCreate)

		codeLen = len(code)
	}

	if codeLen == 0 {
		fmt.Println("[OK] - Compiled Success")
	}

}

func switchInstructions(instruction string, array *[]string, write *os.File) {

	switch instruction {
	case "ldi":
		writeFile(write, 0x04)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "add":
		writeFile(write, 0x05)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "oti":
		writeFile(write, 0x02)
		getArgument(array, 3)
	case "hlt":
		writeFile(write, 0x00)
		getArgument(array, 3)
	default:
		panic(-2)
	}

}

func writeFile(file *os.File, data uint8) {
	binary.Write(file, binary.LittleEndian, data)
}

func verifyArgument(argument []string) uint8 {
	if argument[0] != "x" {
		panic(-1)
	}
	value := strings.Join(argument[1:], "")
	number, _ := strconv.ParseInt(value, 16, 64)
	return uint8(number)
}

func getArgument(code *[]string, number int) {
	slice := (*code)[number:]
	*code = slice
}

func exceptions(errorCode string) {
	switch errorCode {
	case "-1":
		fmt.Println("[ERROR] - Invalid Argument")
		break
	case "-2":
		fmt.Println("[ERROR] - Not Found")
		break
	default:
		fmt.Println("[ERROR] - Compiled Error")
	}

}

func removeCommentsAndScape(file *os.File) []string {
	var valid int
	var s []string

	reader := bufio.NewReader(file)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err != io.EOF {
				fmt.Println("[ERROR] - Read Byte")
				panic(0)
			}
			break
		}

		if b == ';' {
			valid = 1
		}

		if b == '\n' {
			valid = 0
		}

		if valid != 1 && b != ' ' && b != '\n' {
			s = append(s, string(b))
		}
	}

	return s
}
