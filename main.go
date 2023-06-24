package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main(){

	file, err := os.Open("./file.asm8")
	var code[] string

	if err != nil {
		fmt.Println(err)
		return
	}
	code = removeCommentsAndScape(file)

	defer file.Close()

	fileCreate, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Erro ao criar o arquivo: ", err)
		return
	}

	defer fileCreate.Close()

	codeLen := len(code)

    for {

		if codeLen == 0 { 
			break
		}
		instruction := strings.Join(code[:3],"")
		r := switchInstructions(instruction, &code, fileCreate)

		if r == -1 {
			fmt.Println("Compiled Error")
			os.Remove("output.txt")
			break
		}

		codeLen = len(code)
	}

	if codeLen == 0 {	
		fmt.Println("finish compile")
	}

}

func switchInstructions(instruction string, array *[]string, write *os.File) int {
	err := 0
	
	switch instruction {
	case "ldi":
		r := verifyArgument((*array)[4:7])
		if r == -1 {
			break
		}
		bytes := "0x04" + strings.Join((*array)[3:7], "")
		writeFile(write, bytes)
		getArgument(array, 7)
	case "add":
		r := verifyArgument((*array)[4:7])
		if r == -1 {
			break
		}
		bytes := "0x05" + strings.Join((*array)[3:7], "")
		writeFile(write, bytes)
		getArgument(array, 7)
	case "oti":
		writeFile(write, "0x02")
		getArgument(array, 3)
	case "hlt":
		writeFile(write, "0x00")
		getArgument(array, 3)
	default:
		err = -1
	}

	return err
}

func writeFile(file *os.File, bytes string) int {
	_, err := file.WriteString(bytes)
	if err != nil {
		fmt.Println("Error write File", err);
		return -1
	}
	return 0
}

func verifyArgument(argument []string) int {
	if(argument[0] != "x") {
		return -1
	}
	value := strings.Join(argument[1:], "")
	number, _ := strconv.ParseInt(value, 16, 64)
	return int(number)
}

func getArgument(code *[]string, number int){
	slice := (*code)[number:]
	*code = slice
}

func removeCommentsAndScape(file *os.File) [] string {
	var valid int
	var s []string 

	reader := bufio.NewReader(file)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
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