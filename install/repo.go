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

	//  Виконуємо скрипт install-repos.sh
	cmd := exec.Command("bash", "assets/pkglists/repos/install-repos.sh") // Правильний шлях
	cmd.Stdout = os.Stdout                                                //  Перенаправляємо вивід
	cmd.Stderr = os.Stderr                                                //  Перенаправляємо помилки
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%sERROR: Failed to install repositories: %v%s\n", red, err, reset)
		return err //  Повертаємо помилку, а не os.Exit, щоб дати головній програмі вирішити
	}

	fmt.Printf("%s====================================%s\n", yellow, reset)
	fmt.Printf("%sRepositories installed successfully!%s\n", green, reset)
	fmt.Printf("%s====================================%s\n", yellow, reset)
	return nil //  Успішне завершення
}
