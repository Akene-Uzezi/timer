# SkillTimer

A simple command-line countdown timer written in Go with two modes: Timer and Pomodoro.

## Usage

```bash
go run main.go [flags] <duration>
```

### Timer Mode (default)

```bash
go run main.go <duration>
```

### Pomodoro Mode

```bash
go run main.go -p <pomodoro-duration> <break-duration>
```

## Duration Format

The duration accepts Go's `time.ParseDuration` format:
- `1h` - 1 hour
- `30m` - 30 minutes
- `1m30s` - 1 minute 30 seconds
- `45s` - 45 seconds
- `1h30m15s` - 1 hour 30 minutes 15 seconds

If no duration is provided, defaults to `30s` for timer mode and `25m`/`5m` for pomodoro mode.

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

# Pomodoro: 25min work / 5min break (default)
go run main.go -p

# Pomodoro: 50min work / 10min break
go run main.go -p 50m 10m

# Pomodoro: 1h work / 15min break
go run main.go -p 1h 15m
```

## Controls

During the countdown:
- Type `p` + Enter to **pause/resume** the timer
- Type `s` + Enter to **stop/quit** the timer

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

### Timer Mode
The timer counts down from the specified duration, updating the display every second in `MM:SS` format. When the timer reaches zero, it displays "00:00 time elapsed" with a terminal bell.

### Pomodoro Mode
Runs alternating work and break cycles:
- Default: 25 minutes work / 5 minutes break
- After 4 work cycles, takes a long break (15 minutes)
- Cycles continue indefinitely until stopped
- Each transition plays a terminal bell