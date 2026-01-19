## Task Tracer (CLI)

A small CLI app to practice Go basics: argument parsing, JSON persistence, and simple command handling.

### Spec / Requirements

See `domain/task-tracer/project.md`.

### Run

From repo root:

```bash
go run ./cmd/task-tracer <command> [args...]
```

Examples:

```bash
go run ./cmd/task-tracer add "Buy groceries" "Milk, eggs, bread"
go run ./cmd/task-tracer list
go run ./cmd/task-tracer list todo
go run ./cmd/task-tracer update 123 "New title" "New description"
go run ./cmd/task-tracer mark-in-progress 123
go run ./cmd/task-tracer mark-done 123
go run ./cmd/task-tracer delete 123
```

### Data

Tasks are stored in `tasks.json` in the current working directory.

