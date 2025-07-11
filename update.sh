#!/bin/bash

# Kill the running app process
pkill -f today-cli

# Configuration
REPO_URL="https://github.com/drclcomputers/Today-CLI"
APP_DIR="$(dirname "$(realpath "$0")")"

cd "$APP_DIR" || { echo "Failed to enter app directory."; exit 1; }

if [ ! -d ".git" ]; then
    echo "App directory is not a git repository. Cloning fresh copy..."
    cd ..
    rm -rf "$(basename "$APP_DIR")"
    git clone "$REPO_URL"
    echo "Cloned latest version."
else
    echo "Updating app from GitHub..."
    sudo git stash
    sudo git stash clear
    sudo git pull origin main
    sudo go build
    echo "Update complete."
fi