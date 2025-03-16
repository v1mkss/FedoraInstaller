package configs

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mgutz/ansi"
)

func InstallDnfConfig() error {
	// Кольори
	green := ansi.ColorCode("green")
	blue := ansi.ColorCode("blue")
	yellow := ansi.ColorCode("yellow")
	red := ansi.ColorCode("red")
	reset := ansi.ColorCode("reset")

	// Читаємо ASCII-арт з файлу
	asciiArtBytes, err := os.ReadFile("ui/ansi/dnf.txt") //  Правильний шлях!
	if err != nil {
		fmt.Printf("%sError reading DNF ASCII art file: %v%s\n", red, err, reset)
		//  Можна або os.Exit(1), або продовжити без арту
	} else {
		fmt.Print(blue) // Синій колір
		fmt.Println(string(asciiArtBytes))
		fmt.Println(reset)
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sInstalling DNF Configuration...%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)

	log := func(message string) {
		fmt.Printf("%s[%s] %s%s\n", blue, time.Now().Format("2006-01-02 15:04:05"), message, reset)
	}

	checkError := func(err error, message string) {
		if err != nil {
			fmt.Printf("%sERROR: %s: %v%s\n", red, message, err, reset)
			os.Exit(1)
		}
	}

	configFilePath := "assets/configs/dnf/dnf.conf" //  Правильний шлях!
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Printf("%sERROR: Configuration file not found: %s%s\n", red, configFilePath, reset)
		os.Exit(1)
	}

	log("Creating backup of current DNF configuration...")
	if _, err := os.Stat("/etc/dnf/dnf.conf"); err == nil {
		err = os.Rename("/etc/dnf/dnf.conf", "/etc/dnf/dnf.conf.old")
		checkError(err, "Failed to backup old configuration")
		log("Backup created at /etc/dnf/dnf.conf.old")
	}

	log("Installing new DNF configuration...")
	sourceFile, err := os.Open(configFilePath)
	checkError(err, "Failed to open source configuration file")
	defer sourceFile.Close()

	destFile, err := os.Create("/etc/dnf/dnf.conf")
	checkError(err, "Failed to create destination configuration file")
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	checkError(err, "Failed to copy configuration")

	if _, err := os.Stat("/etc/dnf/dnf.conf"); err == nil {
		fmt.Printf("%s-----------------------------------%s\n", yellow, reset)
		fmt.Printf("%sDNF Configuration installed successfully!%s\n", green, reset)
		fmt.Printf("%s-----------------------------------%s\n", yellow, reset)
	} else {
		fmt.Printf("%sFailed to verify DNF configuration%s\n", red, reset)
		os.Exit(1)
	}
	return nil
}
