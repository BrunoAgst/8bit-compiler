package errorsHandling

import "fmt"

func PrintError(errors *[]string, errorsLen int) {
	for i := 0; i < errorsLen; i++ {
		exceptionHandling((*errors)[i])
	}
	fmt.Println("[ERROR] - Compilation Error")
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
