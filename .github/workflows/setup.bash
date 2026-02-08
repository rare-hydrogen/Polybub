#!/bin/bash

# INSTRUCTIONS
# Use this script as a template to setup the deploy agent the first time 
# on a remote server. You shouldn't need to run this again. Designed for 
# root user to execute on an Ubuntu 24.04 (LTS) x64 server with Snap installed already.

# Add the deployment user
sudo useradd \
  --system \
  --create-home \
  --home-dir /home/github_agent \
  --shell /usr/bin/bash \
  github_agent
/

echo "added user"

# Allow the deployment user to use SSL
sudo -u github_agent mkdir -p /home/github_agent/.ssh
sudo -u github_agent chmod 700 /home/github_agent/.ssh
sudo touch /home/github_agent/.ssh/authorized_keys
sudo chown github_agent:github_agent /home/github_agent/.ssh/authorized_keys
sudo chmod 600 /home/github_agent/.ssh/authorized_keys
sudo passwd -l github_agent

echo "added SSL"

# Add directories for the app
sudo mkdir -p /var/www/app/Polybub

echo "added dir"

# Restrict agent to running in non-interactive mode
DEPLOY_DTL="no-port-forwarding,no-agent-forwarding,no-X11-forwarding"
# MANUALLY add a PUBLIC key here
DEPLOY_SSH_PUBLIC_KEY="" # EX: ssh-rsa AAAA...= github_agent
cat >> /home/github_agent/.ssh/authorized_keys <<EOF
$DEPLOY_DTL $DEPLOY_SSH_PUBLIC_KEY
EOF

echo "added authorized_keys"

# Enable agent to use password-less sudo only for specific commands
# or manually modify using sudo nano /etc/sudoers.d/polybub-deploy
echo 'github_agent ALL=(root) NOPASSWD: \
  /bin/systemctl stop Polybub, \
  /bin/systemctl reset-failed Polybub, \
  /bin/systemctl daemon-reload, \
  /usr/bin/systemd-run --unit=golang-web-app --property=User=github_agent --property=AmbientCapabilities=CAP_NET_BIND_SERVICE --property=ReadWritePaths=/var/www/app/golang-web-app/.db /var/www/app/golang-web-app/golang-web-app
  ' | sudo tee /etc/sudoers.d/polybub-deploy

echo "added password-less sudo"

# Or MANUALLY add a sqlitedb file UNDER .db
sudo mkdir -p /var/www/app/Polybub/.db
touch /var/www/app/Polybub/.db/sqlitedb

# Or MANUALLY add a private.pem UNDER .certs
sudo mkdir -p /var/www/app/Polybub/.certs
touch /var/www/app/Polybub/.certs/private.pem

echo "added placeholder paths and files"

# Then, change ownership on all files:
sudo chown github_agent:github_agent /var/www/app/Polybub -R

echo "updated ownership"

# Then, add fail2ban and lock-down users to reduce attacks
sudo apt install fail2ban -y

sudo tee /etc/ssh/sshd_config.d/99-hardening.conf >/dev/null <<'EOF'
PasswordAuthentication no
PermitRootLogin no
EOF

echo "installed and started fail2ban and locked down users"