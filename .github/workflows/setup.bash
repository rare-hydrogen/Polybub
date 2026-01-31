#!/bin/bash

# INSTRUCTIONS
# Use this script as a template to setup the deploy agent the first time 
# on a remote server. You shouldn't need to run this again. Designed for 
# use on an Ubuntu server with Snap installed already.

# Define vars
# Do not commit the public key string!!!
DEPLOY_CMD="bash /var/www/app/Polybub/.github/workflows/deploy/script.bash"
DEPLOY_DTL=",no-port-forwarding,no-agent-forwarding,no-X11-forwarding"
PUB_KEY="ssh AAAA..."

# Add the user
sudo useradd \
  --system \
  --create-home \
  --home-dir /home/github_agent \
  --shell /usr/sbin/bash \
  github_agent
/

# Allow the user to use SSL
sudo -u github_agent mkdir -p /home/github_agent/.ssh
sudo -u github_agent chmod 700 /home/github_agent/.ssh
sudo touch /home/github_agent/.ssh/authorized_keys
sudo chown github_agent:github_agent /home/github_agent/.ssh/authorized_keys
sudo chmod 600 /home/github_agent/.ssh/authorized_keys
sudo passwd -l github_agent

# Add their command + ssh public key 
# TODO: Should I add commands here?
cat >> /home/github_agent/.ssh/authorized_keys <<EOF
command="$DEPLOY_CMD"$DEPLOY_DTL $PUB_KEY
EOF

# Add directories for repo
sudo mkdir -p /var/www/App
sudo chown github_agent:github_agent /var/www/app/Polybub -R

# Add service to systemctl
# TODO: Maybe I can skip this?
cp /var/www/app/Polybub/.github/workflows/deploy/polybub.service /etc/systemd/system/polybub.service
sudo systemctl daemon-reload

# Enable agent to use password-less sudo only for specific systemctl commands
echo 'github_agent ALL=(root) NOPASSWD: deployuser ALL=(root) NOPASSWD: \
  /usr/bin/systemd-run, \
  /bin/systemctl' | sudo tee /etc/sudoers.d/polybub-deploy
# or use sudo nano /etc/sudoers.d/polybub-deploy

# Enable the agent (and by extension, systemctl) to access the private.pem
sudo chown github_agent:github_agent /var/www/app/Polybub/.certs/private.pem
sudo chown github_agent:github_agent /var/www/app/Polybub/.certs