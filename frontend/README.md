# Technology Competition App (Frontend)

This is the frontend application for the Technology Competition App, built with **Vite** and **React**.

---

## Live Demo

If you don't want to run the app locally, you can access the live version here:

👉 [https://areeb-submission.netlify.app/](https://areeb-submission.netlify.app/)

---

## Running Locally

### Prerequisites

-   Node.js (v16+ recommended)
-   npm

### Setup

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

    ```
    VITE_API_URL=http://localhost:8020
    ```

4. **Start the development server:**

    ```bash
    npm run dev
    ```

    The app will be available at [http://localhost:5173](http://localhost:5173) by default.

---

## Default Admin User

You can log in with the following default admin credentials on the custom backend and the live demo as well:

-   **Username:** `admin`
-   **Password:** `admin`

---

## Notes

-   The app expects the backend API URL in the `VITE_API_URL` environment variable.

---
