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
			// Файл /etc/fedora-release не існує, це не Fedora.
			fmt.Println(ansi.Color("This script is intended for Fedora only.", "red"))
			return false
		}
		// Інша помилка (не IsNotExist), виводимо повідомлення, але продовжуємо.
		// Можна було б і panic(err), якщо це критична помилка.
		fmt.Println(ansi.Color(fmt.Sprintf("Error checking for Fedora: %v", err), "red"))
		return false // Або true, залежно від того, чи хочете ви дозволити встановлення при помилках доступу.
	}
	return true
}
