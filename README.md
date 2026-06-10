# Interview Question 009 — Post & Comment App (IT 08-1)

Backend: Go + Gin · Frontend: Angular 22

## Installation

Install the following:

- Go 1.26+ — <https://go.dev/dl/>
- Node.js 24.15+ (or 26+) — <https://nodejs.org/>

Then install frontend dependencies:

```bash
cd frontend
npm install
```

## Run

**Backend** (http://localhost:18080):

```bash
cd api
go run ./cmd/server
```

**Frontend** (http://localhost:14200):

```bash
cd frontend
npm start
```

Open <http://localhost:14200>, type a comment in the box and press **ENTER**.

## Run tests

```bash
cd api && go test ./...      # backend
cd frontend && npm test      # frontend
```
