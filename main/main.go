package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a word.")
		return
	}

	// Retrieve the first command-line argument
	word := os.Args[1]
	if word == "" {
		fmt.Println()
		return
	}

	// Set the default filename to "standard.txt"
	filename := "standard.txt"

	// Check if there are three command-line arguments
	if len(os.Args) == 3 {
		// Retrieve the second command-line argument
		font := os.Args[2]

		// Load the contents of the appropriate file based on the command-line argument
		switch font {
		case "shadow":
			filename = "shadow.txt"
		case "standard":
			filename = "standard.txt"
		case "thinkertoy":
			filename = "thinkertoy.txt"
		default:
			fmt.Println("Please specify a valid font. Valid fonts are shadow, standard, and thinkertoy.")
			return
		}
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var lines []string
	if filename == "thinkertoy.txt" {
		lines = strings.Split(string(content), "\r\n")
	} else {
		lines = strings.Split(string(content), "\n")
	}

	// Print the specified word with the selected font
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
