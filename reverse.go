package main

import (
	"ascii_art/myFunctions"
	"os"
)

func ReverseAsciiArt(string) {
	examplefile := myFunctions.ReadFileLineByLine()
	input, err := os.ReadFile(examplefile)
}
