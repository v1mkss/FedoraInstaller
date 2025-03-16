#!/bin/bash

# Кольори для виводу
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

# ASCII арт
echo -e "${CYAN}"
cat << "EOF"
 ______ _                    ______ _
|  ____(_)                  |  ____(_)
| |__   _ _ __   ___  ___  | |__   _ _  __
|  __| | | '_ \ / _ \/ __| |  __| | | |/ /
| |    | | | | |  __/\__ \ | |    | |   <
|_|    |_|_| |_|\___||___/ |_|    |_|_|\_\
EOF
echo -e "${NC}"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}Installing System Configurations...${NC}"
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

# Встановлення конфігурації Fish
log "Installing Fish configuration..."
./fish/install.sh
check_error "Failed to install Fish configuration"

# Встановлення конфігурації Starship
log "Installing Starship configuration..."
./starship/install.sh
check_error "Failed to install Starship configuration"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ All configurations have been installed successfully!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Додаткова інформація
echo -e "${CYAN}"
echo "Configurations installed:"
echo "• Fish Shell"
echo "• Starship Prompt"
echo -e "${NC}"
