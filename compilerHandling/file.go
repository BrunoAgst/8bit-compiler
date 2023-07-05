package compilerHandling

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

func writeFile(file *os.File, data uint8) {
	binary.Write(file, binary.LittleEndian, data)
}

func RemoveCommentsAndScape(file *os.File, errors *[]string) []string {
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
