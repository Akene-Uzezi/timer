# SkillTimer

A simple command-line countdown timer written in Go with two modes: Timer and Pomodoro.

[![Architecture diagram of akene-uzezi/timer](https://gitdiagram.com/akene-uzezi/timer/diagram.png)](https://gitdiagram.com/akene-uzezi/timer?utm_source=readme&utm_medium=picture)

## Installation

```bash
# Build and install globally as 'timer' (requires sudo)
./buildandinstall.sh
```

After installation, run `timer <duration>` from anywhere.

Or build locally:

```bash
go build -o skilltimer .
./skilltimer 5m
```

## Usage

```bash
timer [flags] <duration>
```

### Timer Mode (default)

```bash
timer <duration>
```

### Pomodoro Mode

```bash
timer -p <work-duration> <break-duration> [cycles]
```

- `work-duration`: Work session length (default: 25m)
- `break-duration`: Short break length (default: 5m)
- `cycles`: Number of work cycles before long break (default: 4)

## Duration Format

Uses Go's `time.ParseDuration` format:

| Format | Meaning |
|--------|---------|
| `1h` | 1 hour |
| `30m` | 30 minutes |
| `1m30s` | 1 minute 30 seconds |
| `45s` | 45 seconds |
| `1h30m15s` | 1 hour 30 minutes 15 seconds |

Defaults: `30s` for timer mode, `25m`/`5m`/`4` for pomodoro mode.

### Examples

```bash
# Timer examples
timer 5m              # 5 minute timer
timer 1h30m           # 1 hour 30 minute timer
timer 90s             # 90 second timer
timer                 # Default 30 second timer

# Pomodoro examples
timer -p                          # Default: 25m work / 5m break / 4 cycles
timer -p 50m 10m                  # 50m work / 10m break / 4 cycles
timer -p 1h 15m                   # 1h work / 15m break / 4 cycles
timer -p 25m 5m 3                 # 25m work / 5m break / 3 cycles before long break
```

## Controls

During countdown, type and press Enter:

| Key | Action |
|-----|--------|
| `p` | Pause / Resume |
| `s` | Stop / Quit |

## How It Works

### Timer Mode

Counts down from specified duration, updating display every second in `MM:SS` format. When reaching zero, displays "00:00 time elapsed" with a terminal bell.

### Pomodoro Mode

Runs alternating work and break cycles:

- Default: 25 minutes work / 5 minutes break
- After N work cycles (default 4), takes a long break (15 minutes)
- Cycles continue indefinitely until stopped
- Each transition plays a terminal bell
- Displays current cycle number during work sessions

## Project Structure

```
.
├── main.go           # Entry point, parses flags
├── timer.go          # Timer mode implementation
├── pomodoro.go       # Pomodoro mode implementation
├── util.go           # Shared utilities (display, input handling)
├── go.mod            # Go module definition
├── buildandinstall.sh # Build and install script
└── README.md
```

## Requirements

- Go 1.22+