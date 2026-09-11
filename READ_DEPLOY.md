There are a few methods you can use to deploy Polybub remotely. Either way, you'll need to handle your secrets with a custom implementation:

- Dockerfile
- Virtual Machine

## Docker

This is by-far the easier method, as it doesn't require as many manual steps. However, you need to use some kind of container deployment / orchestration tool, which may be vendor-specific.

```bash
# Build the image
docker build --tag 'polybub' .

# Run the image
docker run -p 8090:8090 polybub:latest
```

### WARNING:

The image built with the above will contain all files in your repo, including secrets, so DON'T JUST USE the sample Dockerfile. You'll need to decide for yourself how to handle your secrets. I'd recommend using a secrets manager and passing an access token through the docker run command as an argument.

That is not in this repo yet.

## Virtual Machine

This is a portion of my preferred hosting method, and relies on Github Actions + DigitalOcean droplets.

To mimic my setup, there will be a lot of configuration, but hopefully it gives you an idea of what you need to do to get the app running in your own preferred CI/CD + Hosting stack. In addition, I do not recommend you follow these steps exactly, (you should use a third-party secrets manager in production) but this will show you how to get by with minimal integration (AKA hosting it manually).

### Note on TLS:

- Change `is-secure:true` and `port:443`
- Don't use a self-signed cert like this article does, use a DNS provider like Cloudflare to handle TLS for you.
- Self-signed cert files are:
  - `.certs/tls.pem`
  - `.certs/tls.key`

### Two files you should understand:

- `.github/workflows/deploy-vm/deploy-vm.yml` Runs the CI/CD pipeline, updates the deployed app, and recaches the secrets.
- `.github/workflows/deploy-vm/setup.bash` One-time script to setup the VM to accept deployments from CI/CD pipeline.

### VM Setup Steps:

1. Create SSH creds for a deploy user:
   - `ssh-keygen`
2. Create self-signed cert files:
   - `openssl req`
3. Ensure your VM is up:
   - Create the VM.
   - Connect with SSH and setup config.
   - Edit the `.github/workflows/setup.bash` file to use your deploy user SSH creds.
   - Run the setup script with SSH, without copying the file.
     - `ssh droplet_name 'bash -s' < ./setup.bash`
4. Ensure Github has the following repository secrets:
   - `secrets.DEPLOY_USER` (github_agent)
   - `secrets.DEPLOY_HOST` (VM IP: 111.22.333.444)
   - `secrets.DEPLOY_SSH_PRIVATE_KEY` (---BEGIN ... END---)
   - `secrets.CONFIG` (contents of your json file)
   - `secrets.PRIVATE_PEM` (contents of your private.pem)
   - `secrets.CERT_KEY` (contents of your tls.key)
   - `secrets.CERT_PEM` (contents of your tls.pem)
5. Finally, run the `deploy-vm.yml` Github Action.

### Other useful commands:

```bash
# Run setup on remote target, without copying the file to the target
ssh droplet_name 'bash -s' < ./setup.bash`

# Create SSH key for deploy agent, to be used in the setup script
ssh-keygen -t rsa -b 4096

# Mount / Unmount to VM
sshfs droplet_name:/var/www/app/Polybub ~/Remote/mount
umount ~/Remote/mount

# Start / Stop the service
sudo systemctl restart Polybub.service
sudo systemctl stop Polybub.service

# Read logs from the VM
journalctl -u Polybub.service --since="15 min ago"
```

## If you get locked out of digital ocean

I only include this here because it happened three times to me, so I got sick of forgetting how to break back in.

1. reset root password
2. login to recovery terminal
3. add to top of /etc/ssh/ssh_config:
   PasswordAuthentication no
   PubkeyAuthentication yes
4. delete files under /etc/ssh/ssh_config.d (99-hardening... etc)
5. restart ssh `sudo systemctl restart ssh`
6. try to login again
