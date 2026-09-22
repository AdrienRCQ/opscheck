# OPSCheck

![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)![cli](https://img.shields.io/badge/CLI-8A2BE2?style=for-the-badge)![sysops](https://img.shields.io/badge/SysOps-blue?style=for-the-badge)

**OPSCheck** is a command-line utility written in **Go**, designed to provide simple diagnostic and troubleshooting tools for SysOps and system administrators.

It provides several commands for retrieving system information, testing network connectivity and searching log files.

---

## Features

OPSCheck currently provides the following features:

- 🖥️ **System information**
  - Hostname
  - Operating system
  - Architecture
  - CPU

- 🌐 **Network checks**
  - TCP connectivity
  - HTTP/HTTPS connectivity

- 📄 **Log analysis**
  - Search for a string inside a log file

---

## Usage

```bash
opscheck <command> [subcommand] [options]
```

Available commands:

```text
info        Display system information
check       Perform connectivity checks
logs        Perform log-related operations
```

Use the `--help` option to display the available commands and options:

```bash
opscheck --help
```

---

## System Information

The `info` command displays information about the local system.

```bash
opscheck info
```

Information currently returned:

- Hostname
- Operating system
- Architecture
- CPU

Example:

```text
Hostname      : SERVER01
OS            : windows
Architecture  : amd64
CPU           : Intel(R) Xeon(R)
```

---

## Connectivity Checks

The `check` command provides different network connectivity tests.

```bash
opscheck check <subcommand>
```

Currently available:

```text
http
tcp
```

### TCP Check

Tests whether a TCP connection can be established to a remote host and port.

```bash
opscheck check tcp --host <destination> --port <port>
```

Example:

```bash
opscheck check tcp --hsot server01.example.com --port 443
```

This can be used to quickly determine whether a TCP service is reachable from the current machine.

### HTTP Check

Tests HTTP or HTTPS connectivity to a remote endpoint.

```bash
opscheck check http --url <url> --expect-status <http status code>
```

Example:

```bash
opscheck check http --url https://example.com --expect-status 200
```

This command can be used to verify that an HTTP endpoint is reachable.

---

## Log Analysis

The `logs` command provides utilities for inspecting log files.

```bash
opscheck logs <subcommand>
```

Currently available:

```text
filter
```

### Filter

Searches for a specific string inside a log file.

```bash
opscheck logs filter --file <file> --conatins <string>
```

Example:

```bash
opscheck logs filter --file /var/log/application.log --contains "ERROR"
```

This command can be useful for quickly checking whether a specific message, error or event exists in a log file.

---

## Examples

Display system information:

```bash
opscheck info
```

Check whether an HTTPS service is reachable over TCP:

```bash
opscheck check tcp --host example.com --port 443
```

Check an HTTP endpoint:

```bash
opscheck check http --url https://example.com --expect-status 200
```

Search for an error inside a log file:

```bash
opscheck logs filter --file application.log --contains "ERROR"
```

---

## Documentation

Command-specific documentation is available directly from the CLI:

```bash
opscheck --help
```

or:

```bash
opscheck <command> --help
```

For example:

```bash
opscheck check --help
opscheck check tcp --help
opscheck logs filter --help
```

---

## Project Status

OPSCheck is currently under active development.

Additional diagnostic and troubleshooting features will be added over time.

### Current commands

```text
opscheck
├── info
├── check
│   ├── http
│   └── tcp
└── logs
    └── filter
```

---

## Support

If you encounter an issue or would like to suggest a new feature, open an issue in the project repository.

---

## License

See the `LICENSE` file for more information.