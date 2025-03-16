# Fedora Installer

A simple, command-line Fedora installer written in Go. This tool automates the basic setup of a Fedora system, including configuring DNF, installing repositories, and installing essential packages.

## Features

*   **Interactive CLI Menu:** Provides an easy-to-use command-line interface to guide you through the installation process.
*   **Automated DNF Configuration:** Configures the DNF package manager for optimal performance.
*   **Repository Installation:** Installs essential repositories for accessing a wide range of software.
*   **Essential Package Installation:** Installs a curated set of packages necessary for a functional Fedora system.
*   **Theme and Font Configuration:** Sets up default themes and fonts for a visually appealing desktop experience.

## Getting Started

### Prerequisites

*   A running Fedora system (this tool is designed to automate setup on a fresh installation).
*   Root privileges (the installer will attempt to elevate privileges using `sudo`).
*   Internet connection.

### Installation

1.  **Download the latest release:** You can download the latest pre-built binary from the [Releases page](https://github.com/v1mkss/FedoraInstaller/releases/latest).
2.  **Run the installer:**
    ```bash
    ./fedorainstaller
    ```

### Usage

After running the installer, you'll be presented with a menu:

1.  **Install System:** Starts the automated Fedora setup process.
2.  **Options:** Provides access to individual configuration steps:
    *   Install DNF Config
    *   Install Repositories
    *   Install Fish Shell
    *   Install Starship Prompt
3.  **Exit:** Quits the installer.

Use the number keys to select an option and press Enter.

## Configuration

The installer's behavior is driven by configuration files located in the `assets/` directory:

*   `assets/configs/dnf/dnf.conf`: Configuration file for the DNF package manager.
*   `assets/configs/fish/config.fish`: Configuration file for the Fish shell.
*   `assets/configs/starship/starship.toml`: Configuration file for the Starship prompt.
*   `assets/pkglists/fedora_pkglist.txt`: A comprehensive list of packages available for installation.
*   `assets/pkglists/pkgs/base.txt`: Packages included in the base system installation.
*   `assets/pkglists/pkgs/desktop.txt`: Packages related to the desktop environment.
*   `assets/pkglists/pkgs/drivers.txt`: Packages for essential drivers.
*   `assets/pkglists/pkgs/groups.txt`: Package groups to install.
*   `assets/pkglists/repos/`: Scripts to install additional repositories.
*   `assets/scripts/`: Shell scripts that execute specific installation steps.

Feel free to modify these files to customize the installation process to your liking.

## Contributing

Contributions are welcome! If you have any suggestions or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
