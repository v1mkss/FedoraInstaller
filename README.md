# Fedora Installer

A simple command-line installer for Fedora. Automates basic setup, including DNF, repositories, and essential packages.

## Features

*   **Interactive CLI Menu:** Guides through installation.
*   **Automated DNF Configuration:** Optimizes DNF.
*   **Repository Installation:** Installs software repositories.
*   **Essential Package Installation:** Installs key Fedora packages.
*   **Theme and Font Configuration:** Sets up default themes and fonts.

## Getting Started

### Prerequisites

*   Fedora system.
*   Root privileges (`sudo`).
*   Internet connection.

### Installation

1.  **Download:** [Releases page](https://github.com/v1mkss/FedoraInstaller/releases/latest).
2.  **Run:** `./fedorainstaller`.

### Usage

Menu options:

1.  **Install System:** Starts full Fedora setup.
2.  **Options:** Individual configuration steps.
3.  **Exit:** Quit.

Use number keys + Enter.

## Configuration

Configuration files in `assets/`:

*   `assets/configs/dnf/dnf.conf`: DNF configuration.
*   `assets/configs/fish/config.fish`: Fish shell config.
*   `assets/configs/starship/starship.toml`: Starship prompt config.
*   `assets/pkglists/fedora_pkglist.txt`: Full package list.
*   `assets/pkglists/pkgs/base.txt`: Base system packages.
*   `assets/pkglists/pkgs/desktop.txt`: Desktop packages.
*   `assets/pkglists/pkgs/drivers.txt`: Drivers.
*   `assets/pkglists/pkgs/groups.txt`: Package groups.
*   `assets/pkglists/repos/`: Repository scripts.
*   `assets/scripts/`: Installation scripts.

Customize in `assets/`.

## Contributing

Contributions welcome! Open issue or pull request.

## License

MIT License.
