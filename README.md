# SkillTimer

A simple command-line countdown timer written in Go.

## Usage

```bash
go run main.go <duration>
```

### Duration Format

The duration accepts Go's `time.ParseDuration` format:
- `1h` - 1 hour
- `30m` - 30 minutes
- `1m30s` - 1 minute 30 seconds
- `45s` - 45 seconds
- `1h30m15s` - 1 hour 30 minutes 15 seconds

If no duration is provided, defaults to `30s`.

### Examples

```bash
# 5 minute timer
go run main.go 5m

# 1 hour 30 minute timer
go run main.go 1h30m

# 90 second timer
go run main.go 90s

# Default 30 second timer
go run main.go
```

### Controls

During the countdown, type `p` + Enter to **pause/resume** the timer.

## Building

```bash
go build -o skilltimer main.go
./skilltimer 5m
```

### Build and Install

Use the provided script to build and install globally as `timer`:

```bash
./buildandinstall.sh
```

This builds the binary and installs it to `/usr/bin/timer` (requires sudo). After installation, you can run `timer <duration>` from anywhere.

## How It Works

The timer counts down from the specified duration, updating the display every second in `MM:SS` format. When the timer reaches zero, it displays "00:00 time elapsed" with a terminal bell.