package myFunctions

import (
	"strings"
)

// ApplyColor applies the specified color to the input string
// applyColour applies the specified colour to the input string
func ApplyColour(word *string, color string) string {
	// Convert the color to uppercase for case-insensitive matching
	colour := strings.ToUpper(color)

	// Apply the color based on the RGB values
	switch colour {
	case "RED":
		return "\033[31m" + *word + "\033[0m" // Red color
	case "GREEN":
		return "\033[32m" + *word + "\033[0m" // Green color
	case "BLUE":
		return "\033[34m" + *word + "\033[0m" // Blue color
	case "CYAN":
		return "\033[36m" + *word + "\033[0m" // Cyan color
	case "MAGENTA":
		return "\033[35m" + *word + "\033[0m" // Magenta color
	case "YELLOW":
		return "\033[33m" + *word + "\033[0m" // Yellow color
	case "WHITE":
		return "\033[37m" + *word + "\033[0m" // White color
	case "ORANGE":
		return "\033[38;5;208m" + *word + "\033[0m" // Orange color
	case "PINK":
		return "\033[38;5;205m" + *word + "\033[0m" // Pink color
	case "PURPLE":
		return "\033[38;5;165m" + *word + "\033[0m" // Purple color
	case "BLACK":
		return "\033[38;5;232m" + *word + "\033[0m" // Black color
	default:
		// If the provided color is not supported, return the original word
		return *word
	}
}
