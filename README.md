# obsidianOptimizeMCP

A Go-based MCP (Multi-Channel Platform) server for managing and optimizing files in an Obsidian vault, with RESTful APIs, structured logging (Logrus + Gin), and support for custom tools and prompt resources.

## Features
- **File Management**: Create, read, update, delete, and list files in your Obsidian vault via API endpoints and MCP tools.
- **Markdown Optimization**: Resource endpoint for optimizing markdown content (customizable logic).
- **Structured Logging**: Logrus integration for all HTTP and Gin logs, including file/line info.
- **Sampling/Confirmation**: Interactive sampling for dangerous operations (e.g., delete confirmation).
- **Extensible Tools/Prompts**: Easily add new tools and prompt resources for custom workflows.

## Getting Started

### Prerequisites
- Go 1.18+
- [Obsidian](https://obsidian.md/) running with API plugin enabled (or compatible backend)

### Build & Run
```sh
make build      # Build the binary
make run        # Build and run the server
make test       # Run all tests
```

### Configuration
Set environment variables or use command-line flags:
- `OBSIDIAN_URL` (default: `http://127.0.0.1:27123`)
- `OBSIDIAN_TOKEN` (your Obsidian API token)

Example:
```sh
export OBSIDIAN_URL="http://localhost:27123"
export OBSIDIAN_TOKEN="your-token-here"
make run
```

### API/Tool Examples
- **Create File**: `obsidian_create_file`
- **Read File**: `obsidian_read_file`
- **Delete File**: `obsidian_delete_file` (with confirmation)
- **List Files**: `obsidian_list_files`
- **Markdown Optimize Prompt**: `obsidian_md_optimize`

## Project Structure
- `main.go`            — Entry point, Gin server setup
- `pkg/ob-tools.go`    — Tools for file operations
- `pkg/prompts.go`     — Prompt resources (e.g., markdown optimizer)
- `types/`             — Shared types and schemas
- `utils.go`           — Logging and utility functions
- `Makefile`           — Build/test/run automation
- `.gitignore`         — Standard ignores for Go, build, and editor files

## Extending
- **Add a Tool**: Implement the `Tool` interface and register in `pkg/reg.go`
- **Add a Prompt Resource**: Implement the `Prompt` interface and register in `pkg/reg.go`

## License
MIT

---

For more details, see the code comments and each package's documentation.
