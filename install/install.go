package install

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mgutz/ansi"
)

func RunSystemUpdate() error {
	green := ansi.ColorCode("green")
	yellow := ansi.ColorCode("yellow")
	red := ansi.ColorCode("red")
	reset := ansi.ColorCode("reset")

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sRunning System Update...%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)

	// Running script 00-system-update.sh
	cmd := exec.Command("bash", "assets/scripts/00-system-update.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%sERROR: Failed to run system update script: %v%s\n", red, err, reset)
		return err
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sSystem update completed (or at least the script ran)%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)
	return nil
}
