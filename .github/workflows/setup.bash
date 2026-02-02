#!/bin/bash

# INSTRUCTIONS
# Use this script as a template to setup the deploy agent the first time 
# on a remote server. You shouldn't need to run this again. Designed for 
# root user to execute on an Ubuntu server with Snap installed already.

# Add the deployment user
sudo useradd \
  --system \
  --create-home \
  --home-dir /home/github_agent \
  --shell /usr/sbin/nologin \
  github_agent
/

# Allow the deployment user to use SSL
sudo -u github_agent mkdir -p /home/github_agent/.ssh
sudo -u github_agent chmod 700 /home/github_agent/.ssh
sudo touch /home/github_agent/.ssh/authorized_keys
sudo chown github_agent:github_agent /home/github_agent/.ssh/authorized_keys
sudo chmod 600 /home/github_agent/.ssh/authorized_keys
sudo passwd -l github_agent

# Add directories for the app
sudo mkdir -p /var/www/app

# Restrict agent to running in non-interactive mode
DEPLOY_DTL="no-port-forwarding,no-agent-forwarding,no-X11-forwarding"
DEPLOY_SSH_PUBLIC_KEY="ssh AAAA..." # MANUALLY add a PUBLIC key here
cat >> /home/github_agent/.ssh/authorized_keys <<EOF
$DEPLOY_DTL $DEPLOY_SSH_PUBLIC_KEY
EOF

# Enable agent to use password-less sudo only for specific commands
# or manually modify using sudo nano /etc/sudoers.d/polybub-deploy
echo 'github_agent ALL=(root) NOPASSWD: \
  /bin/systemctl stop Polybub, \
  /bin/systemctl reset-failed Polybub, \
  /bin/systemctl daemon-reload, \
  /usr/bin/systemd-run --unit=Polybub --property=User=github_agent --property=ReadWritePaths=/var/www/app/Polybub/.db 
  ' | sudo tee /etc/sudoers.d/polybub-deploy

# MANUALLY ADD A private.pem UNDER .certs
# MANUALLY ADD A sqlite file UNDER .db

# Then, change ownership on all files:
sudo chown github_agent:github_agent /var/www/app/Polybub -R