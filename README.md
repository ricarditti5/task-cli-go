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

**Available statuses:** `todo` · `in-progress` · `done`

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
task-cli list todo
task-cli list "in-progress"
task-cli list done

# Show available commands
task-cli help
```

## Install

### Windows

1. Go to the [**latest release page**](https://github.com/ricarditti5/task-cli-go/releases/latest)
2. Download `task-cli-windows.zip`
3. Extract the `.zip` to a folder of your choice
4. Open the extracted folder — you will find `task-cli.bat`
5. Open a terminal inside that folder and run:

```bash
task-cli help
```

> 💡 To use `task-cli` from any folder, add the extracted folder to your system `PATH`.

### Linux / macOS

1. Go to the [**latest release page**](https://github.com/ricarditti5/task-cli-go/releases/latest)
2. Download `task-cli-linux.tar.gz` or `task-cli-macos.tar.gz`
3. Extract and make it executable:

```bash
tar -xzf task-cli-linux.tar.gz
chmod +x task-cli-linux
./task-cli-linux help
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