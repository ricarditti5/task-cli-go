# Task CLI Go

A minimal command-line task manager written in Go. Tasks are stored persistently in a `task.json` file.

## Commands

| Command | Description |
|---|---|
| `task-cli add "<description>"` | Add a new task |
| `task-cli delte id` | Delete an task |
| `task-cli change-status id "<New status>"` | Change the status of task |
| `task-cli list ` | List all tasks |
| `task-cli list "<not done or in-progress or done>"` | List task by progress |
| `task-cli update id "<New task>"` | Update the task |

## Usage

```bash
# Add a task
task-cli add "buy groceries"
# Change status of task
task-cli change-status 1 "done"
# Delete a task
task-cli delete 1
# List all tasks
task-cli list
# List filtered tasks
task-cli list done or "not done" or in-progress
# Update the description of the task
task-cli update 1 "Buy cow meat"
# Show some commands of the cli app
task-cli help
```

## Install

Run `install.bat` (Windows) — it builds from source if Go is available, otherwise uses the pre-built binary, and optionally adds it to your `PATH`.

Or build manually:

```bash
go build -o task-cli.exe .
```

## Data

Tasks are saved as JSON in the same directory as the executable.

### Little Warning
Don't let the source project in trasferences directory!!!

## License

MIT
