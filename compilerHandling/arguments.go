package compilerHandling

import (
	"os"
	"strconv"
	"strings"
)

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

func ProcessInstruction(array *[]string, file *os.File, errors *[]string) {
	instruction := strings.Join((*array)[:3], "")
	switchInstructions(instruction, array, file, errors)

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
