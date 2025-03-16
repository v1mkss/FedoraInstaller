#!/bin/bash

# Installing RPM Fusion
echo "Installing RPM Fusion repositories..."
sh ./assets/pkglists/repos/rpmfusion.sh

# Installing Terra
echo "Installing Terra repository..."
sh ./assets/pkglists/repos/terra.sh

# Installing Adoptium
echo "Installing Adoptium repository..."
sh ./assets/pkglists/repos/adoptium.sh

# Updating repositories cache
dnf clean all
dnf makecache
