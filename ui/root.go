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

// ClearScreen clears the terminal screen. Platform agnostic, but relies on the "clear" command.
func ClearScreen() {
	var cmd *exec.Cmd

	cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// elevatePrivileges attempts to elevate the program's privileges using sudo.
// It only works on Linux and exits the program if elevation fails.
func elevatePrivileges() {
	if runtime.GOOS != "linux" {
		return // No need to elevate privileges on non-Linux systems
	}
	if os.Geteuid() != 0 {
		fmt.Println(ansi.Color("Elevating Privilages...", "blue"))
		cmd := exec.Command("sudo", os.Args...) // Execute the program again with sudo
		cmd.Stdin = os.Stdin                    // Pass standard input
		cmd.Stdout = os.Stdout                  // Pass standard output
		cmd.Stderr = os.Stderr                  // Pass standard error
		err := cmd.Run()                        // Run the command
		if err != nil {
			fmt.Println(ansi.Color(fmt.Sprintf("Error elevating privileges: %v", err), "red"))
			os.Exit(1) // Exit if we can't get root
		}
		os.Exit(0) // Exit the current process after elevation; sudo will restart it
	}
}

// PrintWelcomeScreen displays the welcome screen, including ASCII art and initial messages.
func PrintWelcomeScreen() {
	elevatePrivileges() // Try to get root
	ClearScreen()       // Clear any previous output

	asciiArtBytes, err := os.ReadFile("ui/ansi/install.txt")
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		return
	}
	asciiArt := string(asciiArtBytes)
	fmt.Println(ansi.Color(asciiArt, "green+b"))
	lines := strings.Count(asciiArt, "\n")

	time.Sleep(2 * time.Second) // Give the user a moment to admire the art
	fmt.Print("\033[1B")        // Move the cursor down one line
	fmt.Print("\033[J")         // Clear the screen from the cursor down

	if !CheckIfFedora() {
		os.Exit(1) // If not fedora, then we are done here
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// showMainMenu displays the main menu and handles user input.
func showMainMenu() {
	for { // Loop until the user exits
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
			continue // Back to the main menu loop
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-3).", "red"))
			time.Sleep(1 * time.Second)
			continue // Back to the main menu loop
		}

		switch choice {
		case 1:
			fmt.Println(ansi.Color("Installing System...", "blue"))
			install.RunSystemUpdate()
			return // Exit the main menu after starting the installation
		case 2:
			showOptionsMenu() // Go to the options menu
		case 3:
			fmt.Println(ansi.Color("Exiting...", "red"))
			os.Exit(0) // Exit the program
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2 or 3.", "red"))
			time.Sleep(1 * time.Second)
		}
	}
}

// showOptionsMenu displays the options menu and handles user input.
func showOptionsMenu() {
	for { // Loop until the user goes back to the main menu
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
			continue // Back to the options menu loop
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-5).", "red"))
			time.Sleep(1 * time.Second)
			continue // Back to the options menu loop
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
			return // Back to the main menu
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2, 3, 4 or 5.", "red"))
			time.Sleep(1 * time.Second)
		}
	}
}
