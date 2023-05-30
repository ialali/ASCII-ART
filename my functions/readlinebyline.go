package myFunctions

import (
	"fmt"
	"os"
	"strings"
)

func ReadLineByLine(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines
}
