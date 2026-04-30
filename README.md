# Training Tracker

Monorepo for a gym workout tracker with:

- `frontend` - Nuxt 3 + TanStack Query
- `backend` - Go + PostgreSQL + Redis

## Structure

```text
.
├── frontend
└── backend
```

## Quick start

1. Start infrastructure:

   ```bash
   make infra-up
   ```

2. Configure env files:

   ```bash
   cp backend/.env.example backend/.env
   cp frontend/.env.example frontend/.env
   ```

3. Install frontend dependencies:

   ```bash
   make frontend-install
   ```

4. Install backend dependencies:

   ```bash
   make backend-tidy
   ```

5. Run services:

   ```bash
   make dev
   ```

## Common commands

```bash
make dev              # run backend + frontend together
make backend-run      # run only backend
make frontend-run     # run only frontend
make backend-test     # run Go tests
make frontend-build   # run Nuxt production build
make goose-install    # install goose into ./bin
make goose-status     # show migrations status
make goose-up         # apply migrations
make goose-down       # rollback one migration
make goose-reset      # rollback all migrations
make goose-create name=add_users_table
```

## Backend architecture

The backend follows a feature-oriented clean architecture inspired by `golang-todoapp`:

```text
backend
├── cmd/training-tracker
├── internal/core
│   ├── config
│   ├── errors
│   ├── httpserver
│   ├── logger
│   ├── postgres
│   └── redis
└── internal/features
    ├── exercises
    ├── statistics
    └── workouts
```

Each feature is split into:

- `service` - business logic and repository contracts
- `repository/postgres` - PostgreSQL implementation
- `transport/http` - HTTP handlers and DTOs

## Next steps

- Add authentication and multi-user ownership
- Add PR records and progression analytics
- Add Redis-backed caching for heavy statistics endpoints
