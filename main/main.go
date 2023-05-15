package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	words := strings.Split(str, `\n`)
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for _, l := range word {
				for lineIndex, line := range lines {
					if lineIndex == (int(l)-32)*9+h {
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}
