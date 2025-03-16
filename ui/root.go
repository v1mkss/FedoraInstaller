package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"FedoraInstaller/install"
	"FedoraInstaller/install/configs"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

func ClearScreen() {
	var cmd *exec.Cmd

	cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func elevatePrivileges() {
	if runtime.GOOS != "linux" {
		return
	}
	if os.Geteuid() != 0 {
		fmt.Println(ansi.Color("Elevating Privilages...", "blue"))
		cmd := exec.Command("sudo", os.Args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(ansi.Color(fmt.Sprintf("Error elevating privileges: %v", err), "red"))
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func PrintWelcomeScreen() {
	elevatePrivileges()
	ClearScreen()

	asciiArtBytes, err := os.ReadFile("ui/ansi/install.txt")
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		return
	}
	asciiArt := string(asciiArtBytes)
	fmt.Println(ansi.Color(asciiArt, "green+b"))
	lines := strings.Count(asciiArt, "\n")

	time.Sleep(2 * time.Second)
	fmt.Print("\033[1B")
	fmt.Print("\033[J")

	if !CheckIfFedora() {
		os.Exit(1)
	}

	fmt.Println(ansi.Color("Welcome to the simple Fedora CLI installer!", "cyan+b"))
	fmt.Println(ansi.Color("This tool will guide you through a basic Fedora installation.", "cyan"))
	fmt.Print("\033[1A")
	fmt.Printf("\033[%dB", lines)
}

var rootCmd = &cobra.Command{
	Use:   "FedoraInstaller",
	Short: "Simple Fedora CLI installer",
	Long:  `This tool will guide you through a basic Fedora installation.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintWelcomeScreen()
		showMainMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showMainMenu() {
	for {
		fmt.Println(ansi.Color("\nMain Menu:", "yellow+b"))
		fmt.Println(ansi.Color("1. Install System", "green"))
		fmt.Println(ansi.Color("2. Options", "green"))

		fmt.Println(ansi.Color("\n3. Exit", "red"))
		fmt.Print(ansi.Color("Enter your choice (1-3): ", "yellow"))

		var choiceStr string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choiceStr = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(ansi.Color("Error reading input:", "red"))
			continue
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-3).", "red"))
			time.Sleep(1 * time.Second)
			continue
		}

		switch choice {
		case 1:
			fmt.Println(ansi.Color("Installing System...", "blue"))
			install.RunSystemUpdate()
			return
		case 2:
			showOptionsMenu()
		case 3:
			fmt.Println(ansi.Color("Exiting...", "red"))
			os.Exit(0)
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2 or 3.", "red"))
			time.Sleep(1 * time.Second)
		}
	}
}
func showOptionsMenu() {
	for {
		fmt.Println(ansi.Color("\nOptions Menu:", "yellow+b"))
		fmt.Println(ansi.Color("1. Install DNF Config", "green"))
		fmt.Println(ansi.Color("2. Install Repositories", "green"))
		fmt.Println(ansi.Color("3. Install Fish", "green"))
		fmt.Println(ansi.Color("4. Install Starship", "green"))

		fmt.Println(ansi.Color("\n5. Back to Main Menu", "yellow"))
		fmt.Print(ansi.Color("Enter your choice (1-5): ", "yellow"))

		var choiceStr string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choiceStr = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(ansi.Color("Error reading input:", "red"))
			continue
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-5).", "red"))
			time.Sleep(1 * time.Second)
			continue
		}

		switch choice {
		case 1:
			fmt.Println(ansi.Color("Installing DNF config...", "blue"))
			configs.InstallDnfConfig()
		case 3:
			fmt.Println(ansi.Color("Installing Fish config...", "blue"))
			configs.InstallFishConfig()
		case 2:
			fmt.Println(ansi.Color("Installing repositories...", "blue"))
			install.InstallRepositories()
		case 4:
			fmt.Println(ansi.Color("Installing Starship config...", "blue"))
			configs.InstallStarshipConfig()
		case 5:
			return
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2, 3, 4 or 5.", "red"))
			time.Sleep(1 * time.Second)
		}
	}
}
