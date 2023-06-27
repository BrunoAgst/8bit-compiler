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

			str := fmt.Sprintf("%v:", r)
			exceptionHandling(str)
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
	case "hlt":
		writeFile(write, 0x00)
		getArgument(array, 3)
	case "nop":
		writeFile(write, 0x01)
		getArgument(array, 3)
	case "oti":
		writeFile(write, 0x02)
		getArgument(array, 3)
	case "lda":
		writeFile(write, 0x03)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
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
	case "sub":
		writeFile(write, 0x06)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "sta":
		writeFile(write, 0x07)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "and":
		writeFile(write, 0x08)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "orl":
		writeFile(write, 0x09)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "xor":
		writeFile(write, 0x0A)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "not":
		writeFile(write, 0x0B)
		getArgument(array, 3)
	case "gta":
		writeFile(write, 0x0C)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ifc":
		writeFile(write, 0x0D)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ifz":
		writeFile(write, 0x0E)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ife":
		writeFile(write, 0x0F)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "iti":
		writeFile(write, 0x10)
		getArgument(array, 3)
	case "shr":
		writeFile(write, 0x11)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "shl":
		writeFile(write, 0x12)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "cpa":
		writeFile(write, 0x13)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ada":
		writeFile(write, 0x14)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "sba":
		writeFile(write, 0x15)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ana":
		writeFile(write, 0x16)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "ora":
		writeFile(write, 0x17)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "xra":
		writeFile(write, 0x18)
		argument := verifyArgument((*array)[4:7])
		writeFile(write, argument)
		getArgument(array, 7)
	case "swa":
		writeFile(write, 0x19)
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

func exceptionHandling(errorCode string) {
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
