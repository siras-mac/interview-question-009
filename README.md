# Interview Question 009 — Post & Comment App (IT 08-1)

A small full-stack application built for example.com: a post feed page where a user
("Blend 285") can comment under a post. Typing a comment and pressing **ENTER**
appends it below the existing comments.

| Layer    | Stack                                          |
|----------|------------------------------------------------|
| Backend  | Go 1.26 + [Gin](https://gin-gonic.com) — clean architecture, in-memory store |
| Frontend | Angular 22 — standalone components, signals     |
| Database | In-memory (PostgreSQL schema provided in `db/schema.sql`) |

## Project structure

```
├── api/                        # Go REST API
│   ├── cmd/server/             # entry point + dependency wiring
│   └── internal/
│       ├── domain/             # entities + repository interfaces (framework-free)
│       ├── usecase/            # business logic
│       ├── repository/memory/  # in-memory repository implementation
│       └── handler/            # Gin HTTP handlers, router, CORS
├── frontend/                   # Angular 22 single-page app
└── db/schema.sql               # PostgreSQL schema + seed (reference)
```

The API follows clean architecture with a single dependency direction
`handler → usecase → domain ← repository`, so the in-memory store can be swapped
for PostgreSQL by implementing the same `domain.PostRepository` /
`domain.CommentRepository` interfaces — no changes needed in usecases or handlers.

## Prerequisites

- Go 1.26+
- Node.js 24.15+ (or 26+) and npm

## Run the backend

```bash
cd api
go run ./cmd/server     # serves on http://localhost:8080
```

### API endpoints

| Method | Path                      | Description                          |
|--------|---------------------------|--------------------------------------|
| GET    | `/api/posts/1`            | Post detail (author, date, image)    |
| GET    | `/api/posts/1/comments`   | Comments of the post, oldest first   |
| POST   | `/api/posts/1/comments`   | Add comment `{authorName, message}`  |
| GET    | `/healthz`                | Health check                         |

Validation: blank `message`/`authorName` → `400`, unknown post id → `404`.

## Run the frontend

```bash
cd frontend
npm install
npm start               # serves on http://localhost:4200, proxies /api to :8080
```

Open <http://localhost:4200>, type a comment in the box and press **ENTER**.

## Run the tests

```bash
# Backend — usecase + handler tests
cd api
go test ./...

# Frontend — service + component tests
cd frontend
npm test
```

## Database design

The app currently persists in memory (resets on restart). `db/schema.sql`
contains the equivalent PostgreSQL design with the same seed data:

- `posts` — id, author_name, image_url, posted_at
- `comments` — id, post_id (FK → posts, cascade), author_name, message, created_at
- index on `comments (post_id, created_at)` for the comment-list query
