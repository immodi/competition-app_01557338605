# Technology Competition App

This repository contains both the frontend and backend components for the **Technology Competition App**.

---

## 📦 Frontend (Vite + React)

### Live Demo

👉 [https://areeb-submission.netlify.app/](https://areeb-submission.netlify.app/)

### Running Locally

#### Prerequisites

- Node.js (v16+ recommended)
- npm

#### Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/immodi/technology-competition-app.git
    cd technology-competition-app/frontend
    ```

2. **Install dependencies:**

    ```bash
    npm install
    ```

3. **Configure environment variables:**

    The project includes a `.env_example` file. To connect to your custom backend, create a `.env` file based on it:

    ```bash
    cp .env_example .env
    ```

    Then edit `.env` to set the `VITE_API_URL` to your backend API endpoint.

    **⚠️ Warning:** Do **NOT** add a trailing slash `/` at the end of the URL.

    Example `.env`:

    ```env
    VITE_API_URL=http://localhost:8020
    ```

4. **Start the development server:**

    ```bash
    npm run dev
    ```

    The app will be available at [http://localhost:5173](http://localhost:5173) by default.

### Default Admin User

You can log in with the following credentials:

- **Username:** `admin`
- **Password:** `admin`

---

## 🚀 Backend (Go + Chi)

A RESTful API built with Go and [Chi router](https://github.com/go-chi/chi) for managing events and users.

### Quick Start

#### Prebuilt Binary (Recommended)

Download a binary from the [Releases](https://github.com/immodi/technology-competition-app/releases) page.

- Windows: `myapp-windows-amd64.exe`
- Linux: `myapp-linux-amd64`
- macOS: `myapp-darwin-amd64`

Make it executable if necessary (`chmod +x`), then run it.

#### Build from Source

##### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- Git installed

##### Clone the Repository

```bash
git clone https://github.com/immodi/technology-competition-app.git
cd technology-competition-app/backend
```

##### Setup Environment Variables

Use the `.env_example` file as a template:

```bash
cp .env_example .env
```

Edit it with appropriate values.

Example `.env`:

```env
JWT_SECRET_KEY=your-very-strong-secret-key
```

##### Install Dependencies and Build

```bash
go mod download
go build -o myapp .
```

##### Run the Application

```bash
./myapp  # on Unix systems
myapp.exe  # on Windows
```

The API will listen on port 8020 by default.

### Additional Notes

- Keep `JWT_SECRET_KEY` secure and never commit it.
- Multiple environments supported—adjust environment variables accordingly.
- For prebuilt binaries, ensure environment variables are set before running.