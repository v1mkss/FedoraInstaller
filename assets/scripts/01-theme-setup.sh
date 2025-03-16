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
 _____ _                           _____      _
|_   _| |                         /  ___|    | |
  | | | |__   ___ _ __ ___   ___ \ `--.  ___| |_ _   _ _ __
  | | | '_ \ / _ \ '_ ` _ \ / _ \ `--. \/ _ \ __| | | | '_ \
  | | | | | |  __/ | | | | |  __//\__/ /  __/ |_| |_| | |_) |
  \_/ |_| |_|\___|_| |_| |_|\___|\____/ \___|\__|\__,_| .__/
                                                       | |
                                                       |_|
EOF
echo -e "${NC}"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}Configuring Theme and Fonts...${NC}"
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

# Check for KDE Plasma
if [ ! -f /usr/bin/plasma-desktop ]; then
    echo -e "${RED}KDE Plasma is not installed!${NC}"
    exit 1
fi

# Configure defaults for new users
log "Configuring defaults for new users..."
mkdir -p /etc/skel/.config
check_error "Failed to create skel directory"

# KDE Configuration
cat << EOF > /etc/skel/.config/kdeglobals
[Icons]
Theme=Papirus

[General]
fixed=Cascadia Code,10,-1,5,50,0,0,0,0,0
font=Cascadia Code,10,-1,5,50,0,0,0,0,0
menuFont=Cascadia Code,10,-1,5,50,0,0,0,0,0
smallestReadableFont=Cascadia Code,8,-1,5,50,0,0,0,0,0
toolBarFont=Cascadia Code,9,-1,5,50,0,0,0,0,0
EOF
check_error "Failed to create KDE configuration"

# Configure for existing users
log "Configuring for existing users..."
echo -e "${YELLOW}-----------------------------------${NC}"
for user_home in /home/*; do
    if [ -d "$user_home" ]; then
        user=$(basename "$user_home")

        if [ "$user" != "lost+found" ]; then
            echo -e "${CYAN}Configuring for user: ${user}${NC}"

            # Create directories
            mkdir -p "$user_home/.config"
            cp /etc/skel/.config/kdeglobals "$user_home/.config/"
            chown -R $user:$user "$user_home/.config"

            # Configure using kwriteconfig5
            log "Applying theme settings for $user..."
            sudo -u $user kwriteconfig5 --file kdeglobals --group Icons --key Theme Papirus

            log "Applying font settings for $user..."
            sudo -u $user kwriteconfig5 --file kdeglobals --group General --key fixed "Cascadia Code,10,-1,5,50,0,0,0,0,0"
            sudo -u $user kwriteconfig5 --file kdeglobals --group General --key font "Cascadia Code,10,-1,5,50,0,0,0,0,0"
            sudo -u $user kwriteconfig5 --file kdeglobals --group General --key menuFont "Cascadia Code,10,-1,5,50,0,0,0,0,0"
            sudo -u $user kwriteconfig5 --file kdeglobals --group General --key smallestReadableFont "Cascadia Code,8,-1,5,50,0,0,0,0,0"
            sudo -u $user kwriteconfig5 --file kdeglobals --group General --key toolBarFont "Cascadia Code,9,-1,5,50,0,0,0,0,0"
        fi
    fi
done
echo -e "${YELLOW}-----------------------------------${NC}"

# Update font cache
log "Updating font cache..."
fc-cache -f -v
check_error "Failed to update font cache"

# Restart Plasma
log "Restarting Plasma..."
killall plasmashell
kstart5 plasmashell
check_error "Failed to restart Plasma"

echo -e "${YELLOW}========================================${NC}"
echo -e "${GREEN}✓ Theme and fonts configured successfully!${NC}"
echo -e "${YELLOW}========================================${NC}"

# Additional information
echo -e "${CYAN}"
echo "Configuration Summary:"
echo "• Papirus Icon Theme enabled"
echo "• Cascadia Code font set as default"
echo "• Font cache updated"
echo "• Plasma shell restarted"
echo -e "${NC}"
