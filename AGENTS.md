# Repository Guidelines

## Project Overview

`wol` is a Go command-line and web application for sending Wake-on-LAN magic packets, listing configured machines, and scheduling wake-ups.

## Architecture & Data Flow

`main.go` delegates to Cobra commands in `cmd/`. The root command loads configuration before command execution; `config/` owns YAML/environment loading, validation, and defaults. `send` resolves a MAC from a flag or configured machine, applies broadcast precedence, then uses `magicpacket/` to send the UDP packet. `serve` provides the web/API handlers, serves the embedded page, and starts configured schedules. Keep changes in the owning package; avoid moving configuration or packet behavior into command handlers.

## Key Directories

- `cmd/`: Cobra commands, web/API handlers, scheduling, and their tests.
- `config/`: configuration types, loading, precedence, and validation.
- `magicpacket/`: packet construction and UDP broadcast.
- `mac/`: MAC-address utilities.
- `templates/`: web UI template; `assets/images/`: documentation image.
- `docker/`, `.github/workflows/`: container and CI/release configuration.

## Development Commands

```sh
go run . --help
go run . send --help
go build ./...
go test ./...
go vet ./...
go test -race ./...
```

CI runs the build, vet, race-enabled tests, golangci-lint, and GoReleaser configuration checks; see `.github/workflows/ci.yaml` and `release.yaml`.

## Code Conventions & Common Patterns

- Format Go with `gofmt`; use standard Go package and exported-name conventions.
- Keep configuration definitions/loading in `config/`, packet/network operations in their owning packages, and CLI wiring in `cmd/`.
- Return errors from package operations and wrap underlying errors with `%w`; command boundaries report fatal execution errors through Cobra.
- Follow existing `context.Context` and `net/http` handler patterns in server code. This project has no UI framework or application-wide state container.
- Preserve CLI flags, config precedence, and defaults unless the requested behavior changes them. Config sources and broadcast precedence are implemented in `config/` and command code; consult those sources rather than duplicating rules.

## Important Files

- `main.go`: executable entry point.
- `cmd/root.go`: root Cobra command and shared config loading.
- `config/config.go`: configuration schema, load, and validation behavior.
- `config.example.yaml`: canonical configuration example.
- `cmd/serve.go`, `cmd/serveApi.go`, `cmd/schedule.go`: web serving, API, and schedules.
- `magicpacket/magicpacket.go`: magic-packet sending.
- `README.md`: user-facing setup, configuration, and usage.

## Runtime/Tooling Preferences

Use Go 1.27 as declared in `go.mod`; dependencies use Go modules (`go.mod`/`go.sum`). No JavaScript runtime or package manager is required. Use the Go toolchain and the CI versions/configuration as the source of truth for tooling.

## Testing & QA

Tests use Go's `testing` package and `testify` assertions. Add focused tests beside behavior changes (`*_test.go` in the owning package), especially for config precedence/validation and command-visible boundaries. Run `go test ./...` locally; before submitting, match CI with `go build ./...`, `go vet ./...`, and `go test -race ./...`. Update README examples when user-facing contracts change, and keep them consistent with `config.example.yaml` and command help.
