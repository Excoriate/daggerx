#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Check if the necessary environment variables are set
if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN is not set. Please set it and try again."
    exit 1
fi

# Extract the module path from the Go module file
MODULE_PATH=$(go list -m)

# Try to get the latest tag
if git describe --tags --abbrev=0 &>/dev/null; then
    MODULE_VERSION=$(git describe --tags --abbrev=0)
else
    echo "No tags found. Please ensure your repository has at least one tag."
    exit 1
fi

echo "Publishing Go module $MODULE_PATH@$MODULE_VERSION"

# Authenticate with GitHub
echo "machine github.com login $GITHUB_ACTOR password $GITHUB_TOKEN" > ~/.netrc

# Push the module version to the Go module proxy
GOPROXY=proxy.golang.org go list -m "$MODULE_PATH@$MODULE_VERSION"

echo "Module published successfully!"

# Trigger indexing on pkg.go.dev
curl "https://pkg.go.dev/$MODULE_PATH@$MODULE_VERSION"

echo "Triggered pkg.go.dev indexing for $MODULE_PATH@$MODULE_VERSION"
