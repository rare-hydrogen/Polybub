# Polybub

A template for building Ultralight Web Apps.

## What is Polybub

Polybub is a basic web app designed to be used as a template for creating efficient, scalable, low-dependency, CRUD-focused web apps within a `Golang + HTMX + Bulma + Sqlite` stack.

If you've ever wanted to build something serious using these technologies, but didn't want to learn how to integrate them from scratch, this repo is for you.

### Benefits

1. High-Performance.
   - Embedded database.
   - Frontend is almost entirely SSR.
   - Build times are rapid thanks to Golang.
2. Minimal dependencies.
   - No NPM.
   - Built-in, minified HTMX and Bulma.
   - Few (15-ish) third-party Golang modules.
3. Local-Friendly DX.
   - Frontend, Backend, and Database all in one repo.
   - Serve REST and UI from one HTTP server.
   - Simple HTML templates, pipe data with [html/template](https://pkg.go.dev/html/template).
4. Ready for Scale
   - Go servers run great on cheap hardware.
   - Scale vertically on a budget.
   - Or, swap databases and scale horizontally.

### Features

1. Elegant CSS (Bulma)
2. Password Encryption (AES-256)
3. Password Resets (Sendgrid)
4. TOTP 2FA
5. Audit logs
6. Soft Deletes
7. Swagger
8. HTTP Request Logging.
9. Unit tests.
10. Fine-Grained user permissions.
11. Internal / External audiences.
12. Toast notifications in frontend.

## How to Get Started

### Local

To run the project with all features locally, you need to do a few things:

1. Clone / Pull the repo.
2. Add your secrets to a `Polybub/config.json` file (see example).
3. Create a Sqlite db under `Polybub/.db/` and apply the `Polybub/Data/Schema/schema.sql` to give it the right tables etc.
4. Create a `private.pem` under `Polybub/.certs/`
5. **BUILD** with `go mod tidy` to get the modules.
6. **RUN** with `go run main.go` or the `launch.json` file.

### Remote

To run your web app on a remote server, you need to handle the same things you did to get it running locally (`database, code, secrets, certs`), before running it on a remote machine.

I've provided an example `setup.bash` you need to run as root before using `deploy.yml` to build the project and deploy it remotely via Github Actions. That yml script needs the following secrets:

- `CONFIG` (your json file)
- `DEPLOY_HOST` (111.22.333.444)
- `DEPLOY_SSH_PRIVATE_KEY` (---BEGIN ... END---)
- `DEPLOY_USER` (github_agent)

If you want TLS/SSL (HTTPS) you'll need to change `is-secure:true` and `port:443` in the config.json file AND generate the following files and add them to `.certs` folder on your remote server:

- `tls.pem`
- `tls.key`

In addition, you may need to run some of the "Other Useful Commands" below to finish setting up your env. For example, if you forgot to change the variables in `setup.bash` those commands will help you fix your server. Regardless, you'll also need to apply schema and add data too to use the app.

## File Tree

Here's a basic overview of where things belong in the project.

```bash
├── Auth # Basic, OAuth2, JWT, TOTP
├── Data # Everything for the Database
│   ├── Services # Backend Logic
├── Routes # HTTP server
│   ├── Api # Endpoints that return JSON
│   ├── Jsend # Structured output for JSON
│   ├── Ui # Endpoints and templates that return HTML
│   ├── Static # Public Files (minified CSS / JS)
├── Tests # Unit tests
├── Utilities # Logger, list of perms, SMTP, Swagger, Config parsing.
```

## How To Extend

This section exists to remind me how to add new features to the project. If you read all the sections in order, they walk-through how to add features to your web app.

### Add a Permission

TODO: add this section

### Modify Schema

TODO: add this section

### Add a HTML Page or Component

**Pages:**

To add a new page, go to `PageHandlers.go` and add a simple get handler for your new page. Also, put a new `page.name.html` file into the PageTemplates folder. In that HTML file, define the layout and what components will be included using html/template {{ template "comp.name" }} to reference what you want. Pages should not need much data (app name / user details), but if they do, pass that data into the template `top-down`. Pages (mostly) wrap components in <body> tag.

**Components:**

To add a new component, you'll want to create a new `NameHandler.go` file under ComponentHandlers. You'll also want to add a new `comp.name.html` file into the ComponentTemplates folder. In that HTML file, use {{ define "comp.name" }} to create the template you want to reference in your page. Lastly, to hydrate your component with data, DO NOT pass data into the template. Just use HTMX to call the get endpoint for your component, and lazy load it. This allows for a `bottom-up` data flow for components. Components (mostly) begin with a <section> tag.

### Add an Endpoint

TODO: add this section

### Add a Unit Test

TODO: add this section

## Other useful commands:

The following are commands that may need to be run manually from time-to-time.

#### Run setup on remote target, without copying the file to the target

```bash
ssh droplet1 'bash -s' < ./setup.bash`
```

#### Create SSH key for github_agent, to be used in the setup script

```bash
ssh-keygen -t rsa -b 4096
```

#### Create cert for server with a passphrase

```bash
openssl genrsa -out private.pem 4096
```

#### Create the pepper, don't forget to add to config.json

```bash
openssl rand -base64 32
```

#### Setup mount to view server files

```bash
sshfs droplet1:/var/www/app/Polybub ~/Remote/mount
```

#### Remove mount

```bash
umount ~/Remote/mount
```

#### Start / Stop the app service

```bash
sudo systemctl restart Polybub.service
sudo systemctl stop Polybub.service
```

#### Read logs from the server

```bash
journalctl -u Polybub.service --since="15 min ago"
```

## Turso Swap

It is actually very easy to swap from embedded sqlite to turso instead. You only need to swap the data package's connection with the turso-friendly equivalent. Use this as your guide:
https://github.com/ytsruh/gorm-libsql?tab=readme-ov-file#pure-go-usage

I have not tried embedded sqlite with Turso. The fact that this works at all blows my mind. I essentially just replaced the existing sqlite connection with the one documented here. Very easy swap, in my opinion, especially since I already had the URL setup:
`url := Utilities.GlobalConfig.TursoDatabaseUrl + "?authToken=" + Utilities.GlobalConfig.TursoAuthToken`
