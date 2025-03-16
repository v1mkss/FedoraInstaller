#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

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

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}Starting System Setup...${NC}"
echo -e "${YELLOW}========================================${NC}"

# 1. DNF Configuration
log "Installing DNF configuration..."
sudo sh ./configs/dnf/install.sh
check_error "Failed to install DNF configuration"

# 2. Installing Repositories
log "Installing repositories..."
sudo sh ./pkglists/repos/install-repos.sh
check_error "Failed to install repositories"

# 3. System Update after adding repositories
log "Updating system..."
dnf upgrade -y --refresh
check_error "System update failed"

# 4. Installing Packages
log "Installing packages..."
sudo sh ./pkglists/install-pkgs.sh
check_error "Failed to install packages"

# 5. Installing Configurations (conditional theme setup)
log "Installing configurations..."
if command -v plasma-desktop &> /dev/null; then
    log "Plasma desktop detected. Running theme setup..."
    sudo sh ./configs/install.sh
    check_error "Failed to install configurations"
else
    log "Plasma desktop not detected. Skipping theme setup."
fi

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ System setup completed successfully!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Summary information
echo -e "${CYAN}"
echo "Setup Summary:"
echo "• DNF configured"
echo "• Repositories installed"
echo "• System updated"
echo "• Packages installed"
echo "• Configurations applied"
echo -e "${NC}"
