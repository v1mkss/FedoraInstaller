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

# Створення резервної копії
log "Creating backup of current DNF configuration..."
if [ -f /etc/dnf/dnf.conf ]; then
  sudo mv /etc/dnf/dnf.conf /etc/dnf/dnf.conf.old
  check_error "Failed to backup old configuration"
  log "Backup created at /etc/dnf/dnf.conf.old"
fi

# Встановлення нової конфігурації
log "Installing new DNF configuration..."
sudo cp assets/configs/dnf/dnf.conf /etc/dnf/dnf.conf
check_error "Failed to install new configuration"

# Перевірка встановлення
if [ -f /etc/dnf/dnf.conf ]; then
  log "DNF Configuration installed successfully!"
else
  echo "Failed to verify DNF configuration"
  exit 1
fi
