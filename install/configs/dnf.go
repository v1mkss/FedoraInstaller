package configs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mgutz/ansi"
)

func InstallDnfConfig() error {
	green := ansi.ColorCode("green")
	yellow := ansi.ColorCode("yellow")
	red := ansi.ColorCode("red")
	reset := ansi.ColorCode("reset")

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sInstalling DNF Configuration...%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)

	// Running install.sh script
	cmd := exec.Command("bash", "assets/configs/dnf/install.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%sERROR: Failed to install DNF configuration: %v%s\n", red, err, reset)
		return err
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sDNF configuration installed successfully!%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)
	return nil
}
