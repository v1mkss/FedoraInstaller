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
cat << "EOF"
 _____           _
|  __ \         | |
| |  \/ ___ _ __| |_ ___  __ _ _ __ ___
| | __ / _ \ '__| __/ _ \/ _` | '_ ` _ \
| |_\ \  __/ |  | ||  __/ (_| | | | | | |
 \____/\___|_|   \__\___|\__,_|_| |_| |_|

EOF
echo -e "${NC}"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}Installing System Packages...${NC}"
echo -e "${YELLOW}========================================${NC}"

# Logging function
log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')] $1${NC}"
}

# Error checking function
check_error() {
    if [ $? -ne 0 ]; then
        echo -e "${RED}ERROR: $1${NC}"
        exit 1
    fi
}

# Installing package groups
log "Installing package groups..."
echo -e "${YELLOW}-----------------------------------${NC}"
while read -r group; do
    echo -e "${CYAN}Installing group: $group${NC}"
    dnf install -y "$group"
    check_error "Failed to install group $group"
done < pkglists/pkgs/groups.txt
echo -e "${YELLOW}-----------------------------------${NC}"

# Installing base packages
log "Installing base packages..."
if [ -f ./assets/pkglists/pkgs/base.txt ]; then
    echo -e "${YELLOW}-----------------------------------${NC}"
    echo -e "${CYAN}Base packages:${NC}"
    cat ./assets/pkglists/pkgs/base.txt
    echo -e "${YELLOW}-----------------------------------${NC}"
    dnf install -y $(cat ./assets/pkglists/pkgs/base.txt)
    check_error "Failed to install base packages"
else
    echo -e "${RED}WARNING: base.txt not found${NC}"
fi

# Installing drivers
log "Installing drivers..."
if [ -f ./assets/pkglists/pkgs/drivers.txt ]; then
    echo -e "${YELLOW}-----------------------------------${NC}"
    echo -e "${CYAN}Drivers:${NC}"
    cat ./assets/pkglists/pkgs/drivers.txt
    echo -e "${YELLOW}-----------------------------------${NC}"
    dnf install -y $(cat ./assets/pkglists/pkgs/drivers.txt)
    check_error "Failed to install drivers"
else
    echo -e "${RED}WARNING: drivers.txt not found${NC}"
fi

# Installing desktop environment
log "Installing desktop environment..."
if [ -f ./assets/pkglists/pkgs/desktop.txt ]; then
    echo -e "${YELLOW}-----------------------------------${NC}"
    echo -e "${CYAN}Desktop packages:${NC}"
    cat ./assets/pkglists/pkglists/pkgs/desktop.txt
    echo -e "${YELLOW}-----------------------------------${NC}"
    dnf install -y $(cat ./assets/pkglists/pkglists/pkgs/desktop.txt)
    check_error "Failed to install desktop packages"
else
    echo -e "${RED}WARNING: desktop.txt not found${NC}"
fi

# Installing Flathub repository
log "Installing Flathub repository..."
echo -e "${YELLOW}-----------------------------------${NC}"
flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
check_error "Failed to add Flathub repository"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ All installations completed successfully!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Summary information
echo -e "${CYAN}"
echo "Installation Summary:"
echo "• Package groups installed"
echo "• Base packages installed"
echo "• Drivers installed"
echo "• Desktop environment installed"
echo -e "${NC}"
