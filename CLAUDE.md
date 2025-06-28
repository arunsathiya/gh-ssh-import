# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gh-ssh-import is a terminal user interface (TUI) application built in Go that helps users upload their local SSH keys to GitHub using the GitHub CLI. The application presents a list of available SSH public keys from the user's ~/.ssh directory and allows them to select and upload keys to their GitHub account.

## Commands

### Build and Run
```bash
go build -o gh-ssh-import
./gh-ssh-import
```

### Development
```bash
go run main.go
```

### Dependencies
```bash
go mod tidy
go mod download
```

### Testing
Currently no tests are present in the codebase.

## Architecture

The application follows a modular architecture with the following key components:

### Core Packages
- **main.go**: Entry point that initializes the Bubble Tea program
- **ssh/**: Handles SSH key discovery and management
  - `ssh.go`: Contains `GetLocalSSHKeys()` function that scans ~/.ssh for .pub files
- **github/**: GitHub API integration via gh CLI
  - `github.go`: Contains `UploadSshPublicKey()` function using gh CLI commands
- **tui/**: Terminal user interface implementation using Bubble Tea framework
  - `tui.go`: Main TUI model with state management and event handling
  - `delegate.go`: List item rendering delegate
  - `keys/keys.go`: Keyboard input mapping and help system
  - `styles/styles.go`: Lipgloss styling definitions for the interface

### Key Dependencies
- **Bubble Tea**: TUI framework for building interactive terminal applications
- **Bubbles**: Reusable TUI components (used for list widget)
- **Lipgloss**: Styling library for terminal interfaces
- **go-gh**: Official Go library for GitHub CLI integration

### Data Flow
1. Application starts and scans ~/.ssh directory for public keys (.pub files)
2. Keys are loaded into a Bubble Tea list model
3. User navigates the list using keyboard controls (Ctrl+j/k for navigation, Enter to select)
4. When a key is selected, it's uploaded to GitHub via the gh CLI
5. Application displays the result and exits

### State Management
The TUI uses a simple state machine with currently only one state (`browsing`). The architecture is prepared for additional states if needed for future features like 1Password integration.

## Dependencies

Requires GitHub CLI (`gh`) to be installed and authenticated for the upload functionality to work.