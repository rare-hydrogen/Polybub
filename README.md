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
4. Add a `private.pem` under `Polybub/.certs/`
5. **BUILD** with `go mod tidy` to get the modules.
6. **RUN** with `go run main.go` or the `launch.json` file.

### Remote

To run your web app on a remote server, you need to handle the same things you did to get it running locally (`database, code, secrets, certs`), before running it on a remote machine.

I've provided an example `setup-template.bash` to configure a VM before using `deploy.yml` to build the project and deploy it via Github Actions. That script needs the following secrets:

- `DEPLOY_HOST`
- `DEPLOY_SSH_PRIVATE_KEY`
- `DEPLOY_USER`

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
├── Static # Public Files (minified CSS / JS)
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

TODO: add this section

### Add an Endpoint

TODO: add this section

### Add a Unit Test

TODO: add this section
