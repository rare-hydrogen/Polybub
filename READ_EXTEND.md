This article exists to give a basic explanation of how to add new features to the project. If you read all the sections in order, they walk-through how to add features to your web app.

## File Tree

Here's a basic overview of where things belong in Polybub:

```bash
├── Auth # Basic, OAuth2, JWT, TOTP
├── Data # Everything for the Database
|   ├── Schema # Database schema for tests and deployment
│   ├── Services # Backend Logic
├── Routes # HTTP server
│   ├── Api # Endpoints that return JSON
│   ├── Jsend # Structured output
│   ├── Ui # Endpoints and templates that return HTML + CSS + JS
│   ├── Static # Public Files (minified CSS / JS)
├── Tests # Unit, Integration, and End-to-End Tests
├── Utilities # Logger, list of perms, SMTP, Swagger, Config.
```

## Add a Permission

Permissions are compressed and signed inside the JWT stored on the clients browser. Every request validates th

1. Open `Utilities/Permissions/Permissions.go`
2. Add a `var` with the `OAuth2.NewPerm()` function.
3. In the database, add a matching `permissions` record for any user that needs it.
4. The next time the user logs in, they'll get your permission.

I've validated that you can add up to 400 permissions comfortably. You could probably add more by removing the some of the audit fields stored in the JWT.

## Check a Permission

- Use `OAuth2.JwtPermit()` if you're verifying a whole route.
- Use `OAuth2.CheckPerm()` if you need to check perms from WITHIN a route.

## Modifying Application Schema

- Use `Data/Models` to track objects
- Use `Data/Schema` to configure database tables
- Use `Data/Services` to configure CRUD operations on the database
- Use `Data/Validators` to define what an acceptable object looks like

### Migrations

I personally do not make use of GORM migrations, I just write the SQL myself and make the changes directly. Feel free to handle this as you desire. However, make sure you keep the `Data/Schema/schema.sql` file updated so your tests work properly.

### Adding Frontend

In general, there are two rules:

- The design hierarchy goes: `Wrappers > Pages > Components`
- A component OR page OR wrapper always requires two things: HTML + Go Handler.

#### Pages:

To add a new page, go to `PageHandlers.go` and add a simple get handler for your new page. Also, put a new `page.name.html` file into the PageTemplates folder. In that HTML file, define the layout and what components will be included using html/template {{ template "comp.name" }} to reference what you want. Pages should not need much data (app name / user details), but if they do, pass that data into the template `top-down`. Pages (mostly) wrap components in <body> tag.

#### Components:

To add a new component, you'll want to create a new `NameHandler.go` file under ComponentHandlers. You'll also want to add a new `comp.name.html` file into the ComponentTemplates folder. In that HTML file, use {{ define "comp.name" }} to create the template you want to reference in your page. Lastly, to hydrate your component with data, DO NOT pass data into the template. Just use HTMX to call the get endpoint for your component, and lazy load it. This allows for a `bottom-up` data flow for components. Components (mostly) begin with a <section> tag.
