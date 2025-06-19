# Typ0 - CLI Typing Practice Tool

A interactive CLI tool for typing practice and speed tests built with Go and Bubble Tea.

![Typ0 Demo](https://img.shields.io/badge/Go-1.21+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Features

- **Interactive TUI** - Minimal terminal interface with real-time feedback
- **Statistics** - WPM and accuracy tracking
- **Mistype Analysis** - Shows which keys you struggle with most
- **Random Sentences** - Practice with different content every time
- **Configurable Length** - Choose your preferred word count
- **Easy Restart** - Press Enter to start a new race
- **Multiple Commands** - Use `race`, `r`, `type`, or `practice`

## Installation

### Using Homebrew (Recommended)

```bash
brew install TusharIbtekar/go-typ0/typ0
```

### Download Pre-built Binaries

Visit [Releases](https://github.com/TusharIbtekar/go-typ0/releases) and download for your platform:

- **macOS**: `typ0-darwin-amd64` or `typ0-darwin-arm64`
- **Linux**: `typ0-linux-amd64` or `typ0-linux-arm64`
- **Windows**: `typ0-windows-amd64.exe` or `typ0-windows-arm64.exe`

### From Source

```bash
git clone https://github.com/TusharIbtekar/go-typ0.git
cd go-typ0
go build
./typ0 race
```

## Usage

### Quick Start

```bash
# Start a typing race (20 words)
typ0 race

# Start a race with custom word count
typ0 race --words 30
typ0 race -w 30

# Use aliases
typ0 r          # Same as race
typ0 type       # Same as race
typ0 practice   # Same as race
```

### Command Options

```bash
# Show help
typ0 --help
typ0 race --help
```

## How It Works

1. **Start a Race** - Run `typ0 race` to begin
2. **Type the Sentence** - Follow the highlighted text with your cursor
3. **See Real-time Feedback** - Green text = correct, red text = mistakes
4. **View Results** - Get your WPM, accuracy, and mistype analysis
5. **Race Again** - Press Enter to start a new race

## Interface

- **Top Box**: Shows the target sentence with color-coded feedback
- **Bottom Box**: Shows your current input with blinking cursor
- **Stats Section**: Displays results and mistype analysis
- **Centered Layout**: Automatically adapts to your terminal size

## Understanding Results

- **WPM (Words Per Minute)**: Your typing speed
- **Accuracy**: Percentage of correctly typed characters
- **Mistypes**: Analysis of which keys you struggle with
- **Time**: Total time taken to complete the sentence

## Development

### Prerequisites

- Go 1.21 or higher
- Git

### Build from Source

```bash
git clone https://github.com/TusharIbtekar/go-typ0.git
cd go-typ0
go mod download
go build
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/typing-wizard`)
3. Commit your changes (`git commit -m 'Add typing wizard feature'`)
4. Push to the branch (`git push origin feature/typing-wizard`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Beautiful UI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- Styling with [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Released with [GoReleaser](https://goreleaser.com/)
