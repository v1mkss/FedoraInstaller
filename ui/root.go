package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
		Log("Elevating Privilages...")          // Log action
		cmd := exec.Command("sudo", os.Args...) // Execute the program again with sudo
		cmd.Stdin = os.Stdin                    // Pass standard input
		cmd.Stdout = os.Stdout                  // Pass standard output
		cmd.Stderr = os.Stderr                  // Pass standard error
		err := cmd.Run()                        // Run the command
		if err != nil {
			fmt.Println(ansi.Color(fmt.Sprintf("Error elevating privileges: %v", err), "red"))
			LogError("Error elevating privileges", err) // Log error
			os.Exit(1)                                  // Exit if we can't get root
		}
		os.Exit(0) // Exit the current process after elevation; sudo will restart it
	}
}

// PrintWelcomeScreen displays the welcome screen
func PrintWelcomeScreen() {
	elevatePrivileges() // Try to get root
	ClearScreen()       // Clear any previous output

	if !CheckIfFedora() {
		Log("Not a Fedora system, exiting.") // Log OS check failure
		fmt.Println(ansi.Color("This script is intended for Fedora only.", "red"))
		os.Exit(1)
	}

	fmt.Println(ansi.Color("Welcome to the simple Fedora CLI installer!", "cyan+b"))
	fmt.Println(ansi.Color("This tool will guide you through a basic Fedora installation.", "cyan"))
	fmt.Println()
}

var rootCmd = &cobra.Command{
	Use:   "FedoraInstaller",
	Short: "Simple Fedora CLI installer",
	Long:  `This tool will guide you through a basic Fedora installation.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := InitializeLogger(); err != nil { // Initialize logger here
			fmt.Println("Failed to initialize logger:", err)
			os.Exit(1)
		}
		PrintWelcomeScreen()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		CloseLogger() // Close logger when command finishes
	},
	Run: func(cmd *cobra.Command, args []string) {
		showMainMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleInstallAction(installFunc func() error, installingMessage string) {
	fmt.Println(ansi.Color(installingMessage, "blue"))
	Log(installingMessage) // Log install action start
	err := installFunc()
	if err == nil {
		fmt.Println(ansi.Color("Successfully", "green"))
		Log(installingMessage + " Successfully") // Log install success
	} else {
		LogError(installingMessage, err) // Log install failure
	}
	fmt.Println(ansi.Color("Press Enter to return to menu", "yellow"))
	fmt.Scanln()
	ClearScreen()
	PrintWelcomeScreen()
}

// showMainMenu displays the main menu and handles user input.
func showMainMenu() {
	for { // Loop until the user exits
		fmt.Println(ansi.Color("\nMain Menu:", "yellow+b"))
		fmt.Println(ansi.Color("1. Install System", "green"))
		fmt.Println(ansi.Color("2. Options", "green"))

		fmt.Println(ansi.Color("\n3. Exit", "red"))
		fmt.Print(ansi.Color("Enter your choice (1-3) or 'q' to quit: ", "yellow"))

		var choiceStr string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choiceStr = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(ansi.Color("Error reading input:", "red"))
			LogError("Error reading main menu input", err) // Log input error
			continue                                       // Back to the main menu loop
		}

		if choiceStr == "q" {
			fmt.Println(ansi.Color("Exiting...", "red"))
			Log("Exiting installer via user input 'q'.") // Log exit action
			os.Exit(0)
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-3) or 'q'.", "red"))
			LogError("Invalid main menu input", err) // Log invalid input
			time.Sleep(1 * time.Second)
			continue // Back to the main menu loop
		}

		switch choice {
		case 1:
			handleInstallAction(install.RunSystemUpdate, "Installing System...")
			Log("Starting system installation.") // Log menu choice
			return                               // Exit the main menu after starting the installation
		case 2:
			ClearScreen()
			showOptionsMenu() // Go to the options menu
			ClearScreen()
			PrintWelcomeScreen()
		case 3:
			fmt.Println(ansi.Color("Exiting...", "red"))
			Log("Exiting installer via menu choice.") // Log exit action
			os.Exit(0)                                // Exit the program
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2 or 3.", "red"))
			LogError("Invalid main menu choice", fmt.Errorf("choice: %d", choice)) // Log invalid choice
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

		fmt.Println(ansi.Color("\n5. Back to Main Menu", "red"))
		fmt.Print(ansi.Color("Enter your choice (1-5) or 'q' to quit: ", "yellow"))

		var choiceStr string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			choiceStr = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(ansi.Color("Error reading input:", "red"))
			LogError("Error reading options menu input", err) // Log input error
			continue                                          // Back to the options menu loop
		}

		if choiceStr == "q" {
			fmt.Println(ansi.Color("Exiting...", "red"))
			Log("Exiting installer via user input 'q' from options menu.") // Log exit action
			os.Exit(0)
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println(ansi.Color("Invalid input. Please enter a number (1-5) or 'q'.", "red"))
			LogError("Invalid options menu input", err) // Log invalid input
			time.Sleep(1 * time.Second)
			continue // Back to the options menu loop
		}

		switch choice {
		case 1:
			handleInstallAction(configs.InstallDnfConfig, "Installing DNF config...")
			Log("Installing DNF config from options menu.") // Log option choice
		case 2:
			handleInstallAction(install.InstallRepositories, "Installing repositories...")
			Log("Installing repositories from options menu.") // Log option choice

		case 3:
			handleInstallAction(configs.InstallFishConfig, "Installing Fish config...")
			Log("Installing Fish config from options menu.") // Log option choice

		case 4:
			handleInstallAction(configs.InstallStarshipConfig, "Installing Starship config...")
			Log("Installing Starship config from options menu.") // Log option choice

		case 5:
			ClearScreen()
			PrintWelcomeScreen()
			Log("Returning to main menu from options menu.") // Log menu navigation
			return                                           // Back to the main menu
		default:
			fmt.Println(ansi.Color("Invalid choice. Please enter 1, 2, 3, 4 or 5.", "red"))
			LogError("Invalid options menu choice", fmt.Errorf("choice: %d", choice)) // Log invalid choice
			time.Sleep(1 * time.Second)
		}
	}
}
