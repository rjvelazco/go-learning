## Task Tracer (CLI)

A small CLI app to practice Go basics: argument parsing, JSON persistence, and simple command handling.

### Spec / Requirements

See `cmd/tasktracer/project.md`.

### Run

From repo root:

```bash
go run ./cmd/tasktracer <command> [args...]
```

Examples:

```bash
go run ./cmd/tasktracer add "Buy groceries" "Milk, eggs, bread"
go run ./cmd/tasktracer list
go run ./cmd/tasktracer list todo
go run ./cmd/tasktracer update 123 "New title" "New description"
go run ./cmd/tasktracer mark-in-progress 123
go run ./cmd/tasktracer mark-done 123
go run ./cmd/tasktracer delete 123
```

Tip: if you want flag-style help (`-h` / `--help`), build the binary and run it directly:

```bash
go build -o tasktracer ./cmd/tasktracer
./tasktracer --help
```

### Data

Tasks are stored in `tasks.json` in the current working directory.

