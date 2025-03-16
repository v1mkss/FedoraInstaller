package ui

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mgutz/ansi"
)

func CheckIfFedora() bool {
	if runtime.GOOS != "linux" {
		fmt.Println(ansi.Color("This script is intended for Linux/Fedora only.", "red"))
		return false
	}
	_, err := os.Stat("/etc/fedora-release")
	if err != nil {
		if os.IsNotExist(err) {
			// The /etc/fedora-release file does not exist, it's not Fedora.
			fmt.Println(ansi.Color("This script is intended for Fedora only.", "red"))
			return false
		}
		// Another error (not IsNotExist), print the message but continue.
		// Could also panic(err) if it's a critical error.
		fmt.Println(ansi.Color(fmt.Sprintf("Error checking for Fedora: %v", err), "red"))
		return false // Or true, depending on whether you want to allow installation if access errors occur.
	}
	return true
}
