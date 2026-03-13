# GitHub Repository CLI Tool

A simple command-line utility written in Go to fetch and display metadata from GitHub repositories using the official REST API

## Features
- Accepts a repository URL or path as an argument
- Performs an HTTP request to the GitHub API
- Parses the JSON response into a structured format
- Displays key repository information in a clean, readable output

## Installation / Building

To build the executable from source, run:

```bash
go build -o analyzeRepo

```

## Usage

Run the program by passing the full URL of the repository you want to analyze as a command-line argument:

```bash
./analyzeRepo [https://github.com/Ganesha1967/golang-course](https://github.com/Ganesha1967/golang-course)

```

### Example Output:

```text
Name:         golang-course
Description:  Homework for GoLang course 2026
Stars:        0
Forks:        0
Created:      2026-03-09T13:03:23Z

```
