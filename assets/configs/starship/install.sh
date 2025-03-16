#!/bin/bash

# Кольори
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

# ASCII арт
echo -e "${CYAN}"
cat << "EOF"
 _____ _                 _     _
/  ___| |               | |   (_)
\ `--.| |_ __ _ _ __ ___| |__  _ _ __
 `--. \ __/ _` | '__/ __| '_ \| | '_ \
/\__/ / || (_| | |  \__ \ | | | | |_) |
\____/ \__\__,_|_|  |___/_| |_|_| .__/
                                | |
                                |_|
EOF
echo -e "${NC}"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}Installing Starship Configuration...${NC}"
echo -e "${YELLOW}========================================${NC}"

# Функція для логування
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

# Створення конфігураційної директорії
log "Creating configuration directories..."
mkdir -p /etc/skel/.config/
check_error "Failed to create configuration directory in /etc/skel"

# Копіювання конфігурації для нових користувачів
log "Copying configuration for new users..."
cp ./assets/configs/starship/install.sh /etc/skel/.config/starship.toml
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
            mkdir -p "$user_home/.config"

            # Копіювання конфігурації
            cp ./starship.toml "$user_home/.config/starship.toml"

            # Встановлення правильних прав власності
            chown -R $user:$user "$user_home/.config/starship.toml"
            check_error "Failed to configure for user $user"
        fi
    fi
done
echo -e "${YELLOW}-----------------------------------${NC}"

# Перевірка наявності starship
if ! command -v starship &> /dev/null; then
    echo -e "${YELLOW}Warning: Starship is not installed!${NC}"
    echo -e "${CYAN}Please ensure Starship is installed via your package manager${NC}"
fi

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ Starship configuration completed!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Додаткова інформація
echo -e "${CYAN}"
echo "Configuration Summary:"
echo "• Global config: /etc/skel/.config/starship.toml"
echo "• User config: ~/.config/starship.toml"
echo "• Shell integration required in shell's rc file:"
echo "  - For bash: eval \"\$(starship init bash)\""
echo "  - For fish: starship init fish | source"
echo "  - For zsh: eval \"\$(starship init zsh)\""
echo -e "${NC}"
