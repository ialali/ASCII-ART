package myFunctions

import "os"

func IsPiped() bool {
	fi, _ := os.Stdout.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}
