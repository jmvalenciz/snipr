# Snippet Manager CLI

Snippet Manager is a simple and efficient command-line tool for managing code snippets. It allows you to create, edit, search, and organize snippets by name and tags. The tool uses SQLite for storage and leverages Bubble Tea to provide a customizable TUI (Text-based User Interface).

## Features

- **Create** and **store** code snippets with names and tags.
- **Fuzzy search** snippets by name and tags with a flexible search format.
- **Organize** snippets by multiple tags.
- **Sync** your snippet database across devices using Syncthing.
- **Simple TUI** interface built with Bubble Tea and Bubble components.

## Table of Contents

- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- Go 1.16 or later
- SQLite3 installed on your system
- Syncthing (optional, for syncing snippets across devices)

### Clone the repository

```bash
git clone https://github.com/jmvalenciz/snipr.git
```

### Compile
```bash
make build
```
# Contributing

We welcome contributions! Here’s how you can help:
- Fork the repository.
- Create a new branch (`git checkout -b <feature-branch>`).
- Commit your changes (`git commit -am 'Add some feature'`).
- Push to the branch (`git push origin feature-branch`).
- Create a new pull request.

# License

This project is licensed under the [BSD 3-Clause License](https://github.com/jmvalenciz/snipr/tree/main?tab=BSD-3-Clause-1-ov-file#readme).
