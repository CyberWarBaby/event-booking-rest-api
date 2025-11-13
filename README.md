# event-booking-rest-api (auth-Rest-api-go)

This repository is a small event booking REST API written in Go using Gin and SQLite. It supports user signup and login (bcrypt password hashing), JWT-based authentication, event CRUD and user registration/cancellation for events.

This README is tailored to the code in this repo and documents the actual routes, files, and how to run the project locally.

## What this project contains

- main.go — app entry point that initializes the sqlite database and starts the Gin server.
- `db/db.go` — SQLite connection and schema creation for `users`, `events`, and `registrations` tables.
- `models/user.go` — User model: save user and validate credentials (uses bcrypt via `utils/hash.go`).
- `models/event.go` — Event model: save, list, get by id, update, delete, register/cancel registration.
- `routes/*.go` — HTTP handlers and route registrations (signup, login, events endpoints).
- `middlewares/auth.go` — JWT-based authentication middleware that extracts the Authorization header and sets `userId` in the Gin context.
- `utils/hash.go` — bcrypt helpers.
- `utils/jwt.go` — token generation and verification (HMAC, hard-coded secret in this repo).

## Routes 

Public endpoints
- POST /signup — register a new user. Request JSON: {"email":"...","password":"..."}
- POST /login — authenticate. Request JSON: {"email":"...","password":"..."}. Response includes a JWT token in the `token` field.
- GET  /events — list all events
- GET  /events/:id — get a single event by id

Protected endpoints (require Authorization header with the raw JWT token returned from /login)
- POST   /events — create a new event (JSON body: name, description, location, dateTime). The created event is associated with the authenticated user.
- PUT    /events/:id — update an event (only allowed by the event owner)
- DELETE /events/:id — delete an event (only allowed by the event owner)
- POST   /events/:id/register — register the authenticated user for the event
- DELETE /events/:id/register — cancel the authenticated user's registration for the event

Notes about authorization header: The middleware expects the token directly in the `Authorization` header (no "Bearer " prefix parsing). Example header value: `Authorization: <token>`

## Configuration & defaults

- The app uses SQLite by default and stores the DB in `api.db` in the repository root.
- JWT secret is currently a constant in `utils/jwt.go`:

   const secretKey = "cypheristesting-go"

   For production use, move this secret to an environment variable and do not commit it.

## How to run locally

Requirements: Go 1.20+ (the module declares 1.24), and the CGO-enabled sqlite driver (the repo already depends on `github.com/mattn/go-sqlite3`).

Run with:

```bash
go run ./
```

This will:
- initialize (or open) `api.db`, create tables if they don't exist, and start a Gin server on :8080.

Quick manual test flow

1) Signup

curl example:

```bash
curl -s -X POST http://localhost:8080/signup \
   -H "Content-Type: application/json" \
   -d '{"email":"alice@example.com","password":"secret"}'
```

2) Login (save the token from the response)

```bash
curl -s -X POST http://localhost:8080/login \
   -H "Content-Type: application/json" \
   -d '{"email":"alice@example.com","password":"secret"}'
```

The response JSON includes `token` — use that value for protected endpoints.

3) Create an event (protected):

```bash
curl -s -X POST http://localhost:8080/events \
   -H "Content-Type: application/json" \
   -H "Authorization: <TOKEN_FROM_LOGIN>" \
   -d '{"name":"Concert","description":"Live show","location":"Arena","dateTime":"2025-12-01T19:00:00Z"}'
```

4) Register for an event (protected):

```bash
curl -s -X POST http://localhost:8080/events/1/register \
   -H "Authorization: <TOKEN_FROM_LOGIN>"
```

## Database schema (created automatically)

- users(id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT UNIQUE, password TEXT)
- events(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT, location TEXT, dateTime DATETIME, user_id INTEGER)
- registrations(id INTEGER PRIMARY KEY AUTOINCREMENT, event_id INTEGER, user_id INTEGER)

## Key implementation details & caveats

- Passwords are hashed using bcrypt at cost 14 via `utils/hash.go`.
- JWT tokens are HMAC-signed with a repo-embedded secret in `utils/jwt.go`. Replace with env-based secret before publishing.
- The middleware expects the Authorization header to contain the token directly (no Bearer prefix). You can change `middlewares/auth.go` to parse `Bearer <token>` if desired.
- The project uses SQLite for simplicity. If you want Postgres or another DB, update `db/db.go` and SQL queries accordingly.
- Error handling is basic — handlers return generic 500 messages in many cases. Add more descriptive errors and input validation for production.

