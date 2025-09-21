# Scruticode
A simple yet powerful tool to evaluate quality and processes in your SDLC. 🚀

<img width="1439" height="899" alt="image" src="https://github.com/user-attachments/assets/d64bdf74-5f31-4a89-bc1c-0d97e1815ec6" />

## Install

To install it you can run

```shell
curl -s https://raw.githubusercontent.com/jd-apprentice/Scruticode/master/install.sh | sudo bash
```

## Requirements (dev)

- Golang
- go-imports (https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- golangci-lint (https://golangci-lint.run/welcome/install/)
- air (https://github.com/air-verse/air)
- pre-commit (https://pre-commit.com/#install)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=jd-apprentice_Scruticode&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=jd-apprentice_Scruticode)

## Usage

Scruticode can be run with the following flags:

-   `--languages`: Specify the languages to analyze (e.g., --languages=golang,typescript).
-   `--platforms`: Specify the platforms (e.g., --platforms=github).
-   `--directory`: Specify the local directory to scan. Defaults to the current directory.
-   `--repository`: Specify a Git repository to clone and scan.

## Folder structure

Based on this [article](https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj)
