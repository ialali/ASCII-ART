package myFunctions

import (
	"fmt"
)

// ApplyColor applies the specified color to the input string
func ApplyColor(text string, color string) string {
	switch color {
	case "red":
		return applyANSIColor(text, 31)
	case "green":
		return applyANSIColor(text, 32)
	case "blue":
		return applyANSIColor(text, 34)
	default:
		return text
	}
}

// applyANSIColor applies the ANSI color code to the input string
func applyANSIColor(text string, colorCode int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, text)
}
