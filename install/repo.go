package install

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mgutz/ansi"
)

func InstallRepositories() error {
	ansi.ColorCode("blue")
	green := ansi.ColorCode("green")
	yellow := ansi.ColorCode("yellow")
	red := ansi.ColorCode("red")
	reset := ansi.ColorCode("reset")

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sInstalling Repositories...%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)

	//  Executing the install-repos.sh script
	cmd := exec.Command("bash", "assets/pkglists/repos/install-repos.sh") // Correct path
	cmd.Stdout = os.Stdout                                                //  Redirecting output
	cmd.Stderr = os.Stderr                                                //  Redirecting errors
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%sERROR: Failed to install repositories: %v%s\n", red, err, reset)
		return err //  Returning the error, not os.Exit, to allow the main program to decide
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sRepositories installed successfully!%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)
	return nil //  Successful completion
}
