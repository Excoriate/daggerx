# Set default environment variables
export ENV := env_var_or_default("ENV", "dev")

# Load environment files
set dotenv-load

# 📋 Default recipe (list available recipes)
default:
    @echo "📋 Listing all available recipes in the justfile..."
    @just --list

# 🎣 Initialize and install required hooks
init-hooks:
    @echo "🎣 Initializing and installing required pre-commit hooks..."
    @echo "📦 Installing pre-commit package using pip..."
    pip install pre-commit
    @echo "🔧 Setting up pre-commit hooks in the local repository..."
    pre-commit install --install-hooks

# 🏃 Run all the hooks described in the .pre-commit-config.yaml file
run-hooks:
    @echo "🏃 Running all pre-commit hooks defined in .pre-commit-config.yaml..."
    pre-commit run --all-files

# 🚀 Execute all the Go CI tasks in the pkg/root module
go-ci: go-tidy go-fmt go-vet go-lint go-test
    @echo "🚀 Executing all Go CI tasks for the pkg/root module..."

# 📦 Publish the Go module to the registry
go-publish:
    @echo "📦 Publishing Go module to the registry..."
    @echo "🏃 Running publish-go-module.sh script..."
    @./scripts/publish-go-module.sh

# Go tasks (imported from taskfiles/Taskfile.go.yml)

# 🧹 Tidy Go module dependencies
go-tidy:
    @echo "🧹 Tidying Go module dependencies in the current directory..."
    go mod tidy

# 💅 Format Go code
go-fmt:
    @echo "💅 Formatting all Go code in the current directory and subdirectories..."
    go fmt ./...

# 🔍 Vet Go code
go-vet:
    @echo "🔍 Vetting all Go code in the current directory and subdirectories..."
    go vet ./...

# 🚨 Lint Go code
go-lint *ARGS:
    @echo "🚨 Linting Go code using golangci-lint..."
    golangci-lint run {{ARGS}}

# 🧪 Run Go tests
go-test:
    @echo "🧪 Running all Go tests in verbose mode..."
    go test -v ./...

# 🎭 Run pre-commit hooks on staged files
pc-staged:
    @echo "🎭 Running pre-commit hooks on staged files in the current repository..."
    pre-commit run

# 🔄 Update pre-commit hooks
pc-update:
    @echo "🔄 Updating all pre-commit hooks to their latest versions..."
    pre-commit autoupdate

# 🧼 Clean pre-commit cache
pc-clean:
    @echo "🧼 Cleaning pre-commit cache to ensure a fresh state..."
    pre-commit clean
