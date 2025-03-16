#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

# ASCII art
echo -e "${CYAN}"
echo -e "${YELLOW}========================================${NC}"
log "Creating configuration directories..."
mkdir -p /etc/fish/conf.d/
mkdir -p /etc/skel/.config/fish/
check_error "Failed to create configuration directories"

# Копіювання конфігурації для нових користувачів
log "Copying configuration for new users..."
sudo cp assets/configs/fish/config.fish /etc/skel/.config/fish/config.fish
check_error "Failed to copy configuration for new users"

# Налаштування для існуючих користувачів
log "Configuring for existing users..."
echo -e "${YELLOW}-----------------------------------${NC}"
for user_home in /home/*; do
    if [ -d "$user_home" ]; then
        user=$(basename "$user_home")

        # Пропускаємо системні директорії
        if [ "$user" != "lost+found" ]; then
            echo -e "${CYAN}Configuring for user: ${user}${NC}"

            # Створення директорії конфігурації
            mkdir -p "$user_home/.config/fish/"

            # Копіювання конфігурації
            sudo cp config.fish "$user_home/.config/fish/config.fish"

            # Встановлення правильних прав власності
            sudo chown -R $user:$user "$user_home/.config/fish"
            check_error "Failed to configure for user $user"
        fi
    fi
done
echo -e "${YELLOW}-----------------------------------${NC}"

# Встановлення fish як типової оболонки
log "Setting up Fish as default shell..."
if command -v fish &> /dev/null; then
    # Додавання fish до /etc/shells
    if ! grep -q "$(which fish)" /etc/shells; then
        echo "$(which fish)" >> /etc/shells
        log "Added Fish to /etc/shells"
    fi

    # Зміна типової оболонки
    sudo sed -i "s|^SHELL=.*|SHELL=$(which fish)|" /etc/default/useradd
    check_error "Failed to set Fish as default shell"

    echo -e "${GREEN}✓ Fish shell has been set as default for new users${NC}"
else
    echo -e "${RED}✗ Fish shell is not installed!${NC}"
    exit 1
fi

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ Fish shell configuration completed!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Додаткова інформація
echo -e "${CYAN}"
echo "Configuration Summary:"
echo "• Default shell: $(which fish)"
echo "• Config location: /etc/fish/conf.d/"
echo "• User config: ~/.config/fish/config.fish"
echo -e "${NC}"
}

# Logging function
log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')] $1${NC}"
}

# Функція для перевірки помилок
check_error() {
    if [ $? -ne 0 ]; then
        echo -e "${RED}ERROR: $1${NC}"
        exit 1
    fi
}
