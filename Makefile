# Variables
BINARY_NAME=syswatch
SRC_DIR=./cmd/main.go
TEST_DIR=./tests/metrics
STATIC_DIR=./internal/server/static/
CSS_INPUT=./app/styles/app.css
CSS_OUTPUT=${STATIC_DIR}/output.css

# Default command (build)
all: build

# Development
dev:
	@echo "Starting development mode..."
	@bash -c 'trap "kill 0" SIGINT; \
	templ generate -watch & \
	air & \
	wait'
# Run dev_win if you are using windows 
dev_win:
	@echo "Starting development mode..."
	@powershell -Command "Start-Process templ -ArgumentList 'generate','-watch'; Start-Process air; \
		try { Wait-Process -Name templ, air } catch { Write-Host 'Development mode terminated.'}"

# Build the binary
build:
	@echo "Building the project..."
	@go build -o ./bin/$(BINARY_NAME) ${SRC_DIR}

# Run tests
test:
	@echo "Running tests..."
	@cd ${TEST_DIR} && go test . -v

# Running the app locally
run:
	@./bin/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f ./bin/${BINARY_NAME}

# Watch and build CSS
css:
	@echo "Building Tailwind CSS..."
	@npx tailwindcss -i ${CSS_INPUT} -o ${CSS_OUTPUT} --watch

# Production CSS build (minified)
css-build:
	@echo "Building Tailwind CSS for production..."
	@npx tailwindcss -i ${CSS_INPUT} -o ${CSS_OUTPUT} --minify

temple:
	@temple generate -watch &
# Instructions for new users
help:
	@echo "Usage:"
	@echo "  make dev           Start the development server (with live reloading)"
	@echo "  make build         Build the project"
	@echo "  make test          Run the tests"
	@echo "  make clean         Clean build artifacts"
	@echo "  make css           Watch and build CSS with Tailwind"
	@echo "  make css-build     Build and minify CSS for production"

.PHONY: dev build test clean css css-build help
