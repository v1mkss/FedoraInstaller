#!/bin/bash

dnf install --nogpgcheck --repofrompath 'terra,https://repos.fyralabs.com/terra$releasever' terra-release
sed -i 's/repo_gpgcheck=1/repo_gpgcheck=0/' /etc/yum.repos.d/terra.repo
