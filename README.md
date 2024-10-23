# SysWatch

**SysWatch** is a system monitoring application written in Go. It provides real-time monitoring of system metrics, such as CPU, memory, disk usage, network activity, and processes. The application features a web-based interface using Go, WebSockets, HTMX, and TailwindCSS for dynamic metric updates and a modern UI.

## Features

- Real-time system monitoring
- WebSocket integration for live data updates every 3 seconds
- TailwindCSS styling with dark and light modes
- Metrics include CPU, memory, disk, network, processes, and system information
- Web-based UI for easy access and monitoring
- Mobile responsive design

## Project Structure

```
.
├── app
│   ├── components         # UI components for the frontend
│   └── styles             # CSS or Tailwind styles
├── bin                    # Compiled binaries
├── cmd                    # Main entry point of the application
├── go.mod                 # Go module definition
├── go.sum                 # Go module dependencies
├── internal               # Internal packages
│   ├── metrics            # Metrics gathering logic
│   ├── server             # Server-related code
│   │   ├── server.go      # HTTP server logic
│   │   └── static         # Embedded static files
│   └── ws                 # WebSocket handling logic
├── LICENSE                # License for the project
├── Makefile               # Makefile for build, test, and development commands
├── README.md              # Project documentation
├── tailwind.config.js      # TailwindCSS configuration
└── tests
    ├── metrics            # Unit tests for metrics components
```

## Requirements

- **Go 1.19+**
- **Node.js** (for building TailwindCSS)
- **Templ** (for template generation)
- **Air** (for live reloading in development)

## Setup

To get started with **SysWatch**, follow the steps below to set up the environment, build, and run the application.

### 1. Clone the repository

```bash
git clone https://github.com/thisisamr/SysWatch.git
cd SysWatch
```

### 2. Install dependencies

Ensure you have all Go module dependencies installed. Run:

```bash
go mod tidy
```

This will tidy up the `go.mod` and `go.sum` files and download the necessary modules.

### 3. Ensure Node.js is Installed

SysWatch uses **TailwindCSS** for styling, which requires Node.js to run. You will need **Node.js** and **npm** installed. Below are instructions for installing Node.js via `nvm` (Node Version Manager) for Linux and macOS, and direct binary installation for Windows.

#### Linux/macOS (Using `nvm`)

1. Install `nvm`:

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh | bash
```

2. Activate `nvm`:

```bash
export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
```

3. Install Node.js using `nvm`:

```bash
nvm install --lts
```

4. Verify installation:

```bash
node -v
npm -v
```

#### Windows (Using Node.js Binary Installer)

1. Download the latest **LTS** version of Node.js from the official [Node.js website](https://nodejs.org/en/download/).

2. Run the installer and follow the prompts.

3. After installation, open a terminal (PowerShell or Command Prompt) and verify that Node.js and npm were installed correctly:

```bash
node -v
npm -v
```

### 4. Specify Environment Variables

SysWatch requires environment variables for `HOST` and `PORT` to run properly. You can either:

1. Create an `.env` file in the root of the project and define the following:

```
HOST=localhost
PORT=3000
```

2. Or, specify them directly in your system environment:

```bash
export HOST=localhost
export PORT=3000
```

If no environment variables are set, the application will fall back to defaults (`localhost:3000`).

### 5. Run in Development Mode

For local development with live reloading, run:

- **Linux/macOS**: Use the command:

  ```bash
  make dev
  ```

- **Windows**: Use the command:

  ```bash
  make dev_win
  ```

On Windows, make sure to modify the `air.toml` configuration file as follows:

- Replace `bin = "./tmp/main"` with `bin = "./tmp/main.exe"`
- Replace `cmd = "go build -tags dev -o ./tmp/main cmd/main.go"` with `cmd = "go build -tags dev -o ./tmp/main.exe cmd/main.go"`

This will ensure that the correct binary (`main.exe`) is generated for Windows development.

### 6. Build for Production

To build the application, you can use the `make build` command, which will compile the binary into the `/bin` directory.

```bash
make build
```

You can then run the binary from the `/bin` folder:

```bash
./bin/syswatch
```

### 7. Run the Application

After building or running in development mode, access the application by navigating to:

```
http://host-name:3000
```

### 8. TailwindCSS

To watch and compile your TailwindCSS files, use the following command:

```bash
make css
```

This will watch the CSS files and compile them into `./static/output.css`.

### 9. Run Tests

SysWatch includes unit tests for various components. You can run all tests with:

```bash
make test
```

### 10. Clean Up

To clean up build artifacts, run:

```bash
make clean
```

This will remove the compiled binaries and other temporary files.

## Experimental Release

You can download the compiled binaries for the **experimental release v0.1.0** from the [releases section](https://github.com/thisisamr/SysWatch/releases). This version provides a basic functional version of SysWatch for testing purposes.

---

## Development Tools

1. **Templ**: For generating HTML templates. You need to install Templ before running the app.
2. **Air**: For live reloading during development. Install via:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

## Contributing

Contributions are welcome! Please open an issue or pull request for any improvements or bugs.

## License

This project is licensed under the MIT License.
