<p align="center">
  <br>
  <br>
  <a href="https://github.com/rare-hydrogen/Polybub" target="_blank" rel="noopener noreferrer">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="./polybub.png">
      <source media="(prefers-color-scheme: light)" srcset="./polybub.png">
      <img alt="polybub logo" src="./polybub.png" height="60">
    </picture>
  </a>
  <br>
  <br>
</p>
<br/>

# Polybub

> Ultralight backend and fronted template for web apps, built with:

- :hamster: Golang
- :horse: HTMX
- :dragon_face: Bulma
- :thinking: GORM (Postgres or Sqlite)

Polybub is a web app template for creating efficient, scalable, low-dependency, CRUD-focused web apps with a simple, opinionated stack. All with only 15 modules, so say goodbye to dependency hell!

## Built-In Features:

1. :love_letter: User Management (Encryption + JWT + Resets + TOTP 2FA)
2. :princess: Fine-Grained User Permissions
3. :speak_no_evil: Audit Logs, Soft Deletes, No Nulls,
4. :sunglasses: Swagger
5. :page_with_curl: HTTP Request Logging with systemctl.
6. :japanese_goblin: Unit, Integration, and End-to-End tests in one repo!
7. :bread: Unified FE-BE validation and toast notifications.

## Getting Started

To run Polybub locally:

1. Clone the repo.
2. Add your secrets to a `config.json` file. (see example_config.json)
3. Add `private.pem` under `/.certs/` for encryption.
4. Run `Data/Schema/schema.sql` on your database.
5. Run `Data/Schema/create-admin.sql` on your database.
6. Run `go mod tidy` to get the modules.
7. Run `go run main.go` or `launch.json` file.

### Helpful Commands:

```bash
# Create private.pem
openssl genrsa -out private.pem 4096

# Create the pepper, don't forget to manually add to config.json
openssl rand -base64 32
```

## Additional Information

- [READ_DEPLOY.md](./READ_DEPLOY.md) explains the prod deployment process in more detail.
- [READ_EXTEND.md](./READ_EXTEND.md) explains how to customize / extend this template.

---

# Credits:

I am deeply grateful to the following people / projects for making Polybub possible:

- [ThePrimeagen](https://www.youtube.com/c/theprimeagen)
- [Golang](https://go.dev/)
- [HTMX](https://htmx.org/)
- [Bulma](https://bulma.io/)
- [Gorm](https://gorm.io/)
- [Sqlite](https://sqlite.org/index.html)
- [Postgres](https://www.postgresql.org/)
- [clevergo/jsend](https://github.com/clevergo/jsend)
- [Rhymond/go-money](github.com/Rhymond/go-money)
- [go-ozzo/ozzo-validation](github.com/go-ozzo/ozzo-validation)
- [go-swagno/swagno-http](github.com/go-swagno/swagno-http)
