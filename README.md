# Task CLI Go

A minimal command-line task manager written in Go. Tasks are stored persistently in a `task.json` file located in the same directory as the executable.

## Commands

| Command | Description |
|---|---|
| `task-cli add "<description>"` | Add a new task |
| `task-cli delete <id>` | Delete a task |
| `task-cli update <id> "<new description>"` | Update a task description |
| `task-cli change-status <id> "<status>"` | Change the status of a task |
| `task-cli list` | List all tasks |
| `task-cli list "<status>"` | List tasks filtered by status |
| `task-cli help` | Show available commands |

**Available statuses:** `"not done"` · `in-progress` · `done`

## Usage

```bash
# Add a task
task-cli add "Buy groceries"

# Update a task description
task-cli update 1 "Buy cow meat"

# Change the status of a task
task-cli change-status 1 "in-progress"

# Delete a task
task-cli delete 1

# List all tasks
task-cli list

# List tasks filtered by status
task-cli list "not done"
task-cli list "in-progress"
task-cli list done

# Show available commands
task-cli help
```

## Install

Download the latest release here: **[github.com/ricarditti5/task-cli-go/releases/latest](https://github.com/ricarditti5/task-cli-go/releases/latest)**

### Windows

Download the latest release and extract the `.zip` file. It includes:

```
task-cli-windows.zip
├── task-cli.exe
└── task-cli.bat
```

Run `task-cli.bat` from the terminal or add the folder to your `PATH` to use it from anywhere.

### Linux / macOS

Download the binary from the latest release and make it executable:

```bash
chmod +x task-cli-linux
./task-cli-linux
```

### Build from source

Requires [Go](https://golang.org/dl/) installed.

```bash
git clone https://github.com/ricarditti5/task-cli-go
cd task-cli-go
go build -o task-cli.exe .
```

## Data

Tasks are saved as `task.json` in the same directory as the executable. The file is created automatically on first use.

> ⚠️ **Warning:** Do not run the executable from a temporary or transfers directory — the `task.json` file will be created there and may be lost.

## License

MIT