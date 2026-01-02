# rpt (rocker-project-template)

`rpt` is a CLI tool that scaffolds RStudio development environments based on
`rocker/rstudio` Docker images.

## Features

- Initialize R projects with Docker and Docker Compose
- Select R version via CLI
- Minimal and reproducible setup

## Installation (Linux)

Download the binary from GitHub Releases and place it in your PATH.

## Usage

```bash
rpt init my-project --r-version 4.4
```

## Generated Files

- Dockerfile
- compose.yaml
- README.md
