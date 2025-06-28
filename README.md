# 🔑 gh-ssh-import

<img width="1138" alt="image" src="https://github.com/arunsathiya/gh-ssh-import/assets/18581859/bb8ca897-2259-4b28-9a4a-91e5992206f6">

> **Work in Progress** - This project is currently under active development.

A terminal user interface (TUI) application that helps you upload your local SSH keys to GitHub. Built with Go and Bubble Tea, it scans your `~/.ssh` directory for public keys and allows you to select and upload them to your GitHub account using the GitHub CLI.

## Features

- 🔍 Automatically discovers SSH public keys in your `~/.ssh` directory
- 📋 Interactive list interface for key selection
- ⬆️ Direct upload to GitHub via GitHub CLI integration
- 🎨 Clean, modern terminal interface

## Prerequisites

- [GitHub CLI](https://cli.github.com/) installed and authenticated
- Go 1.21+ (for building from source)

## Usage

```bash
# Build and run
go build -o gh-ssh-import
./gh-ssh-import

# Or run directly
go run main.go
```

## Roadmap

- [ ] 1Password SSH key integration
- [ ] Key management features
- [ ] Bulk upload support

## Contributing

This is a work-in-progress project. Feel free to open issues or submit pull requests!
