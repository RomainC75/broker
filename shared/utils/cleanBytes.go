package utils

import "fmt"

func CleanByte(b []byte) []byte {
	if len(b) == 0 {
		return []byte{}
	}
	position := len(b) - 1
	fmt.Println("position : ", position)
	for position != -1 && b[position] == '\x00' {
		position--
	}
	newB := make([]byte, position+1)
	copy(newB, b[:position+1])
	return newB
}
