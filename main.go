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

	var code []string
	errors := []string{}

	file, err := os.Open("./file.asm")
	if err != nil {
		fmt.Println("[ERROR] - Read File", err)
		panic(0)
	}

	code = removeCommentsAndScape(file, &errors)

	file.Close()

	fileCreate, err := os.Create("file.bin")
	if err != nil {
		fmt.Println("[ERROR] - Create File", err)
		panic(0)
	}

	defer fileCreate.Close()

	codeLen := len(code)
	errorsLen := 0

	for {

		if errorsLen > 0 {
			printError(&errors, errorsLen)
			break
		}

		if codeLen == 0 {
			break
		}
		instruction := strings.Join(code[:3], "")
		switchInstructions(instruction, &code, fileCreate, &errors)

		codeLen = len(code)
		errorsLen = len(errors)

	}
	if errorsLen == 0 {
		fmt.Println("[OK] - Compiled Success")
	}
}

func printError(errors *[]string, errorsLen int) {
	for i := 0; i < errorsLen; i++ {
		exceptionHandling((*errors)[i])
	}
	fmt.Println("[ERROR] - Compilation Error")
}

func switchInstructions(instruction string, array *[]string, write *os.File, errors *[]string) {

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
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ldi":
		writeFile(write, 0x04)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "add":
		writeFile(write, 0x05)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "sub":
		writeFile(write, 0x06)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "sta":
		writeFile(write, 0x07)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "and":
		writeFile(write, 0x08)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "orl":
		writeFile(write, 0x09)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "xor":
		writeFile(write, 0x0A)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "not":
		writeFile(write, 0x0B)
		getArgument(array, 3)
	case "gta":
		writeFile(write, 0x0C)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ifc":
		writeFile(write, 0x0D)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ifz":
		writeFile(write, 0x0E)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ife":
		writeFile(write, 0x0F)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "iti":
		writeFile(write, 0x10)
		getArgument(array, 3)
	case "shr":
		writeFile(write, 0x11)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "shl":
		writeFile(write, 0x12)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "cpa":
		writeFile(write, 0x13)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ada":
		writeFile(write, 0x14)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "sba":
		writeFile(write, 0x15)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ana":
		writeFile(write, 0x16)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ora":
		writeFile(write, 0x17)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "xra":
		writeFile(write, 0x18)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "swa":
		writeFile(write, 0x19)
		getArgument(array, 3)
	case "puh":
		writeFile(write, 0x1a)
		getArgument(array, 3)
	case "pop":
		writeFile(write, 0x1b)
		getArgument(array, 3)
	case "csr":
		writeFile(write, 0x1c)
		argument := verifyArgument((*array)[4:7], errors)
		writeFile(write, argument)
		getArgument(array, 7)
	case "ret":
		writeFile(write, 0x1d)
		getArgument(array, 3)
	default:
		*errors = append(*errors, "-2")
	}

}

func writeFile(file *os.File, data uint8) {
	binary.Write(file, binary.LittleEndian, data)
}

func verifyArgument(argument []string, errors *[]string) uint8 {
	if argument[0] != "x" {
		*errors = append(*errors, "-1")
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
	case "0":
		fmt.Println("[ERROR] - Read Byte")
	case "-1":
		fmt.Println("[ERROR] - Invalid Argument")
		break
	case "-2":
		fmt.Println("[ERROR] - Instruction Not Found")
		break
	default:
		fmt.Println("[ERROR] - Compiled Error")
	}
}

func removeCommentsAndScape(file *os.File, errors *[]string) []string {
	var valid int
	var s []string

	reader := bufio.NewReader(file)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err != io.EOF {
				*errors = append(*errors, "0")
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
