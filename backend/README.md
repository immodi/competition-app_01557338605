# Submission Backend API

A RESTful API built with Go and [Chi router](https://github.com/go-chi/chi) for managing events and users.

---

## Quick Start

### Use a Prebuilt Binary (Recommended)

If you just want to run the API without building from source, you can download a ready-made binary for your OS from the [Releases](https://github.com/immodi/technology-competition-app/releases) page.

- Windows: `myapp-windows-amd64.exe`
- Linux: `myapp-linux-amd64`
- macOS: `myapp-darwin-amd64`

Simply download the appropriate binary, make it executable if needed (`chmod +x` on Linux/macOS), and run it.

---

## Build from Source

If you want to build the API yourself, follow these steps:

### Prerequisites

- [Go 1.24+](https://golang.org/dl/) installed on your system
- Git installed
- Make sure `GOPATH` and `GOROOT` are properly configured

### Clone the repository

```bash
git clone https://github.com/immodi/technology-competition-app.git
cd technology-competition-app/backend
```

### Setup environment variables

The API requires certain environment variables to run. You can provide them by:

- Creating a `.env` file in the root of the project, **or**
- Setting system environment variables directly

> The repo contains a `.env_example` file. You can copy or rename it to `.env` and edit the values accordingly:

```bash
cp .env_example .env
```

### Install dependencies and build

```bash
go mod download
go build -o myapp ./cmd/myapp
```

Replace `./cmd/myapp` with your main package path if different.

---

### Run the application

```bash
./myapp
```

or on Windows

```powershell
myapp.exe
```

The API will start and listen on the configured port (default 8020 if not overridden).

---

## Environment Variable Configuration

You can use any method to supply environment variables:

- Place them in a `.env` file in the root folder
- Export them in your shell session before running
- Use a Docker environment configuration if containerizing

Example `.env` file:

```env
JWT_SECRET_KEY=your-very-strong-secret-key
```

---

## Additional Notes

- Make sure `JWT_SECRET_KEY` is kept secret and never committed publicly.
- The project supports multiple environments; adjust variables accordingly.
- If you use the prebuilt binaries, ensure your `.env` file or system environment variables are set up before running the binary.

---
