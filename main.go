package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	charWidth  = 9
	charHeight = 8
)

func main() {
	// Load banner files into a map
	bannerFiles := map[rune]string{
		' ': "banners/shadow.txt",
		'!': "banners/standard.txt",
		'"': "banners/tinkertoy.txt",
		// Add more banners as needed
	}

	// Read input from user
	fmt.Println("Enter a string:")
	var input string
	fmt.Scanln(&input)

	// Initialize an empty slice to store the ASCII art representation
	var art []string

	// Loop through each character in the input string
	for _, char := range input {
		if char == '\n' {
			// Add blank lines to the art slice for newline characters
			for i := 0; i < charHeight; i++ {
				art = append(art, "")
			}
			continue
		}
		// Get the banner file for the current character
		bannerFile, ok := bannerFiles[char]
		if !ok {
			// If there's no banner file for the current character, use a space banner as default
			bannerFile = bannerFiles[' ']
		}
		// Load the banner file into a string using os.ReadFile
		bannerBytes, err := os.ReadFile(bannerFile)
		if err != nil {
			panic(err)
		}
		bannerStr := string(bannerBytes)

		// Split the banner string into lines and add them to the art slice
		bannerLines := strings.Split(bannerStr, "\n")
		for i := 0; i < charHeight; i++ {
			line := bannerLines[i]
			// Add padding spaces to the line to ensure it's the same width as other characters
			if len(line) < charWidth {
				line += strings.Repeat(" ", charWidth-len(line))
			}
			art = append(art, line)
		}
	}

	// Print the ASCII art representation
	for i := 0; i < charHeight; i++ {
		for j := 0; j < len(input); j++ {
			index := j*charWidth + i
			if index >= len(art[0]) {
				break
			}
			fmt.Print(string(art[index]))
		}
		fmt.Println()
	}
}
