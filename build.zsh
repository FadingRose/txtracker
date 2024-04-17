#!/bin/zsh


set -e

BUILD_DIR="./build/bin"
EXECUTABLE_NAME="txtracker"

echo "Cleaning up existing build directory..."
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

# echo "Running automated tests..."
# go test -count=1 ./tests/...

echo "Building the project..."
go build -o "${BUILD_DIR}/${EXECUTABLE_NAME}" ./cmd/txtracker

echo "Build successful! Executable is located at ${BUILD_DIR}/${EXECUTABLE_NAME}"