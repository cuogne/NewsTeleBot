#!/usr/bin/env bash

## how to run: bash setup.sh

set -euo pipefail

echo "==> Initializing project setup..."

# 1. Install Go dependencies
echo "==> Installing Go dependencies..."
go mod download
go mod tidy
echo "[OK] Go dependencies installed"

# 2. Create .env from .env.example
touch .env
cat .env.example > .env
echo "[OK] Created .env from .env.example"

# 3. Install air (for hot reload) if not available
if ! command -v air >/dev/null 2>&1; then
	echo "==> air not found. Installing air..."
	
	go install github.com/air-verse/air@latest

	GOPATH_BIN="$(go env GOPATH)/bin"
	SHELL_RC=""

	if [ -n "${ZSH_VERSION:-}" ]; then
		SHELL_RC="$HOME/.zshrc"
	elif [ -n "${BASH_VERSION:-}" ]; then
		SHELL_RC="$HOME/.bashrc"
	else
		SHELL_RC="$HOME/.zshrc"
	fi

	if ! printf '%s' ":$PATH:" | grep -q ":${GOPATH_BIN}:"; then
		if [ -f "$SHELL_RC" ]; then
			grep -q 'export PATH="$PATH:$(go env GOPATH)/bin"' "$SHELL_RC" || \
				echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> "$SHELL_RC"
		else
			echo 'export PATH="$PATH:$(go env GOPATH)/bin"' > "$SHELL_RC"
		fi
		export PATH="$PATH:$GOPATH_BIN"
		echo "[OK] Added GOPATH/bin to PATH in $SHELL_RC"
	fi

	echo "[OK] air installed, you can use 'make dev' to start the development server with hot reload"
else
	echo "[SKIP] air is already installed"
fi

# done
echo "==> Setup completed successfully."
echo "You can fill in the .env file with your Telegram Bot Token and other necessary environment variables to configure the bot."