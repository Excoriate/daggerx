#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Check if the necessary environment variables are set
if [ -z "$GITHUB_TOKEN" ]; then
    echo "GITHUB_TOKEN is not set. Please set it and try again."
    exit 1
fi

# Extract the module path and version from the Go module file
MODULE_PATH=$(go list -m)
MODULE_VERSION=$(git describe --tags --abbrev=0)

echo "Publishing Go module $MODULE_PATH@$MODULE_VERSION"

# Authenticate with GitHub
echo "machine github.com login $GITHUB_ACTOR password $GITHUB_TOKEN" > ~/.netrc

# Push the module version to the Go module proxy
GOPROXY=proxy.golang.org go list -m $MODULE_PATH@$MODULE_VERSION

echo "Module published successfully!"
