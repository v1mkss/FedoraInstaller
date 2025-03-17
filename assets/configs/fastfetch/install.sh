#!/bin/bash

# Функція для логування
log() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

# Функція для перевірки помилок
check_error() {
  if [ $? -ne 0 ]; then
    echo "ERROR: $1"
    exit 1
  fi
}

# Створення конфігураційної директорії
log "Creating configuration directories..."
sudo mkdir -p /etc/skel/.config/fastfetch
check_error "Failed to create configuration directory in /etc/skel"

# Копіювання конфігурації для нових користувачів
log "Copying configuration for new users..."
if [ -f assets/configs/fastfetch/config.jsonc ]; then
  sudo cp assets/configs/fastfetch/config.jsonc /etc/skel/.config/fastfetch/config.jsonc
  check_error "Failed to copy configuration for new users"
else
  echo "ERROR: assets/configs/fastfetch/config.jsonc not found"
  exit 1
fi

# Налаштування для існуючих користувачів
log "Configuring for existing users..."
for user_home in /home/*; do
  if [ -d "$user_home" ]; then
    user=$(basename "$user_home")

    # Пропускаємо системні директорії
    if [ "$user" != "lost+found" ]; then
      log "Configuring for user: ${user}"

      # Створення директорії конфігурації
      sudo mkdir -p "$user_home/.config/fastfetch"

      # Копіювання конфігурації
      if [ -f assets/configs/fastfetch/config.jsonc ]; then
        sudo cp assets/configs/fastfetch/config.jsonc "$user_home/.config/fastfetch/config.jsonc"
        check_error "Failed to configure for user $user"
      else
        echo "ERROR: assets/configs/fastfetch/config.jsonc not found"
        exit 1
      fi
    fi
  fi
done

# Перевірка наявності fastfetch
if ! command -v fastfetch &> /dev/null; then
  echo "Warning: Fastfetch is not installed!"
  echo "Please ensure Fastfetch is installed via your package manager"
fi
