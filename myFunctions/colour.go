package myFunctions

import "strings"

func ApplyColour(text string, colour string) string {
	// Convert the colour to uppercase for case-insensitive matching
	colour = strings.ToUpper(colour)

	// Apply the colour based on the RGB values
	switch colour {
	case "RED":
		return "\033[31m" + text + "\033[0m" // Red color
	case "GREEN":
		return "\033[32m" + text + "\033[0m" // Green color
	case "BLUE":
		return "\033[34m" + text + "\033[0m" // Blue color
	case "CYAN":
		return "\033[36m" + text + "\033[0m" // Cyan color
	case "MAGENTA":
		return "\033[35m" + text + "\033[0m" // Magenta color
	case "YELLOW":
		return "\033[33m" + text + "\033[0m" // Yellow color
	case "WHITE":
		return "\033[37m" + text + "\033[0m" // White color
	case "ORANGE":
		return "\033[38;5;208m" + text + "\033[0m" // Orange color
	case "PINK":
		return "\033[38;5;205m" + text + "\033[0m" // Pink color
	case "PURPLE":
		return "\033[38;5;165m" + text + "\033[0m" // Purple color
	case "BLACK":
		return "\033[38;5;232m" + text + "\033[0m" // Black color
	default:
		// If the provided colour is not supported, return the original text
		return text
	}
}
