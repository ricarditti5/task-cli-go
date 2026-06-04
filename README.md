# Task CLI Go

A minimal command-line task manager written in Go. Tasks are stored persistently in a `task.json` file.

## Commands

| Command | Description |
|---|---|
| `task-cli add "<description>"` | Add a new task |

## Usage

```bash
# Add a task
task-cli add "buy groceries"
```

## Install

Run `install.bat` (Windows) — it builds from source if Go is available, otherwise uses the pre-built binary, and optionally adds it to your `PATH`.

Or build manually:

```bash
go build -o task-cli.exe .
```

## Data

Tasks are saved as JSON in the same directory as the executable.

## License

MIT
