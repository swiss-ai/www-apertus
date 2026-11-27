#!/usr/bin/env bash

#------------------------------------------------------------------------------
# @file
# Builds a Hugo project hosted on Vercel.
#------------------------------------------------------------------------------

# Exit on error, undefined variables, or pipe failures
set -euo pipefail

# Define tool versions
DART_SASS_VERSION=1.108.2
GO_VERSION=1.27.4
HUGO_VERSION=0.170.1
NODE_VERSION=24.13.0

# Set the build time zone
TZ=Europe/Zurich

# Set the build cache directory
HUGO_CACHEDIR="${PWD}/.vercel/cache/hugo"

# Perform cleanup
cleanup() {
  if [[ -n "${build_temp_dir:-}" && -d "${build_temp_dir:-}" ]]; then
    rm -rf "${build_temp_dir}"
  fi
}

# Register the cleanup trap
trap cleanup EXIT SIGINT SIGTERM

main() {
  # Export the build time zone
  export TZ

  # Export the build cache directory
  export HUGO_CACHEDIR

  # Create a temporary directory for downloads
  build_temp_dir=$(mktemp -d)

  # Create a local tools directory
  mkdir -p "${HOME}/.local"

  # Install Dart Sass
  echo "Installing Dart Sass ${DART_SASS_VERSION}..."
  curl -sfL --output-dir "${build_temp_dir}" -O "https://github.com/sass/dart-sass/releases/download/${DART_SASS_VERSION}/dart-sass-${DART_SASS_VERSION}-linux-x64.tar.gz"
  tar -C "${HOME}/.local" -xf "${build_temp_dir}/dart-sass-${DART_SASS_VERSION}-linux-x64.tar.gz"
  export PATH="${HOME}/.local/dart-sass:${PATH}"

  # Install Go
  if [[ -f "go.mod" ]]; then
    echo "Installing Go ${GO_VERSION}..."
    curl -sfL --output-dir "${build_temp_dir}" -O "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"
    tar -C "${HOME}/.local" -xf "${build_temp_dir}/go${GO_VERSION}.linux-amd64.tar.gz"
    export PATH="${HOME}/.local/go/bin:${PATH}"
  fi

  # Install Hugo
  echo "Installing Hugo ${HUGO_VERSION}..."
  curl -sfL --output-dir "${build_temp_dir}" -O "https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_linux-amd64.tar.gz"
  mkdir -p "${HOME}/.local/hugo"
  tar -C "${HOME}/.local/hugo" -xf "${build_temp_dir}/hugo_${HUGO_VERSION}_linux-amd64.tar.gz"
  export PATH="${HOME}/.local/hugo:${PATH}"

  # Install Node.js
  if [[ -f "package-lock.json" ]]; then
    echo "Installing Node.js ${NODE_VERSION}..."
    curl -sfL --output-dir "${build_temp_dir}" -O "https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.gz"
    tar -C "${HOME}/.local" -xf "${build_temp_dir}/node-v${NODE_VERSION}-linux-x64.tar.gz"
    export PATH="${HOME}/.local/node-v${NODE_VERSION}-linux-x64/bin:${PATH}"
  fi

  # Log tool versions
  echo "Logging tool versions..."
  command -v sass &> /dev/null && echo "Dart Sass: $(sass --version)" || echo "Dart Sass: not installed"
  command -v go &> /dev/null && echo "Go: $(go version)" || echo "Go: not installed"
  command -v hugo &> /dev/null && echo "Hugo: $(hugo version)" || echo "Hugo: not installed"
  command -v node &> /dev/null && echo "Node.js: $(node --version)" || echo "Node.js: not installed"

  # Configure Git
  echo "Configuring Git..."
  git config --global core.quotepath false

  # Fetch full Git history
  if [[ $(git rev-parse --is-shallow-repository) == true ]]; then
    echo "Fetching full Git history..."
    git fetch --unshallow
  fi

  # Initialize Git submodules
  if [[ -f .gitmodules ]]; then
    echo "Initializing Git submodules..."
    git submodule update --init --recursive
  fi

  # Install Node.js dependencies
  if [[ -f package-lock.json ]]; then
    echo "Installing Node.js dependencies..."
    npm ci
  fi

  # Build the project
  echo "Building the project..."
  hugo build --gc --minify
}

main "$@"
