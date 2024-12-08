# Project Title: AOC 2024 [Go]

This is a Go implementation of the Advent of Code 2024 puzzles.

## Table of Contents

- [Project Title: AOC 2024 \[Go\]](#project-title-aoc-2024-go)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)

## Installation

1. Install dependencies

```bash
go get -u github.com/ankuryd/aoc2024
```

2. Fetch session cookie from browser

    - Open browser and navigate to https://adventofcode.com/2024
    - Open developer tools and navigate to Application tab
    - Find cookie named `session` and copy its value

3. Set up environment variables

```bash
echo "SESSION_COOKIE=<session_cookie>" >> .env
```

## Usage

To get help:

```bash
go run main.go -h
```

To run a specific day:

```bash
go run main.go -d <day>
```

To run all days:

```bash
go run main.go -a
```

To run with test input:

```bash
go run main.go -d <day> -t
go run main.go -a -t
```
