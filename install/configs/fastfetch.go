package configs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mgutz/ansi"
)

func InstallFastfetchConfig() error {
	green := ansi.ColorCode("green")
	yellow := ansi.ColorCode("yellow")
	red := ansi.ColorCode("red")
	reset := ansi.ColorCode("reset")

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sInstalling Fastfetch Configuration...%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)

	// Запускаємо скрипт install.sh
	cmd := exec.Command("bash", "assets/configs/fastfetch/install.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%sERROR: Failed to install Fastfetch configuration: %v%s\n", red, err, reset)
		return err
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sFastfetch configuration installed successfully!%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)
	return nil
}
