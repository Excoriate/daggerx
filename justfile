# Set default environment variables
export ENV := env_var_or_default("ENV", "dev")

# Load environment files
set dotenv-load

# Default recipe (list available recipes)
default:
    @just --list

# Initialize and install required hooks
pc-init:
    @echo "Initializing and installing required hooks..."
    @just precommit-hooks-init

# Run all the hooks described in the .pre-commit-config.yaml file
pc-run:
    @echo "Running all pre-commit hooks..."
    @just precommit-hooks-run

# Execute all the go CI tasks in the pkg/root module
go-ci:
    @echo "Executing Go CI tasks..."
    @cd cli && just go-tidy
    @cd cli && just go-fmt
    @cd cli && just go-vet
    @cd cli && just go-lint
    @cd cli && just go-test

# Publish the go module to the registry
go-publish:
    @echo "Publishing Go module to registry..."
    @./scripts/publish-go-module.sh

# Pre-commit tasks (imported from taskfiles/taskfile.precommit.yml)
precommit-hooks-init:
    @echo "Initializing pre-commit hooks..."
    # Add commands to initialize pre-commit hooks

precommit-hooks-run:
    @echo "Running pre-commit hooks..."
    # Add commands to run pre-commit hooks

# Go tasks (imported from taskfiles/Taskfile.go.yml)
go-tidy:
    @echo "Running go mod tidy..."
    go mod tidy

go-fmt:
    @echo "Formatting Go code..."
    go fmt ./...

go-vet:
    @echo "Vetting Go code..."
    go vet ./...

go-lint:
    @echo "Linting Go code..."
    golangci-lint run

go-test:
    @echo "Running Go tests..."
    go test -v ./...