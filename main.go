package main

import (
	"ascii_art/myFunctions"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ascii-art [OPTIONS] WORD")
		flag.PrintDefaults()
	}

	// Define the color flag
	flag.String("colour", "", "Specify the color for the letters")

	// Parse the command-line arguments
	flag.Parse()

	// Retrieve the positional argument (word)
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

	word := args[0]

	// Set the default filename to "standard.txt"
	filename := "fonts/standard.txt"

	switch len(args) {
	case 1:
		// Only word argument provided, perform ASCII art on the word using default font
		// No need to do anything here
	case 2:
		// Second argument is provided, check if it's a color flag
		if args[1] == "--color" {
			if len(args) < 4 {
				fmt.Println("Please specify both a color and a word.")
				return
			}
			color := args[2]
			word = myFunctions.ApplyColour(&word, color)
		
			if len(args) < 5 {
				fmt.Println("Please specify a font.")
				return
			}
			font := args[3]
			switch font {
			case "shadow":
				filename = "fonts/shadow.txt"
			case "standard":
				filename = "fonts/standard.txt"
			case "thinkertoy":
				filename = "fonts/thinkertoy.txt"
			default:
				fmt.Println("Please specify a valid font. Valid fonts are shadow, standard, and thinkertoy.")
				return
			}
		}
	default:
		fmt.Println("Invalid number of arguments.")
		return
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var lines []string
	if filename == "fonts/thinkertoy.txt" {
		lines = strings.Split(string(content), "\r\n")
	} else {
		lines = strings.Split(string(content), "\n")
	}

	// Generate the ASCII art for the specified word
	var asciiArt strings.Builder
	for h := 1; h < 9; h++ {
		for _, l := range word {
			for lineIndex, line := range lines {
				if lineIndex == (int(l)-32)*9+h {
					asciiArt.WriteString(line)
				}
			}
		}
		asciiArt.WriteString("\n")
	}

	// Output the ASCII art
	output := asciiArt.String()

	// Check if the output needs to be piped into another command
	if myFunctions.IsPiped() {
		cmd := exec.Command("cat", "-e")
		cmd.Stdin = strings.NewReader(output)
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print(output)
	}
}
