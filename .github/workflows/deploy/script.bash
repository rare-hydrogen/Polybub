#!/bin/bash

# INSTRUCTIONS
# This script is run by the agent to update the deployed app.

# Stop the application if it is running
pkill Polybub

# Navigate to this repo
cd /var/www/app/Polybub

# Pull latest
git init
git fetch origin main
git reset --hard origin/main
git pull

# Build and Start Application
/snap/bin/go build -buildvcs=false
sudo systemctl restart polybub.service

echo Deployment Completed
exit 0