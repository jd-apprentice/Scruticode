#!/bin/bash

# --- Configuration ---
REPO="jd-apprentice/Scruticode"
INSTALL_DIR="/usr/local/bin"
DOWNLOAD_URL=""

error_exit() {
  echo "Error: $1" >&2
  echo "Installation failed." >&2
  exit 1
}

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

if ! command_exists "curl"; then
  error_exit "'curl' command not found. Please install it."
fi

if ! command_exists "tar"; then
  error_exit "'tar' command not found. Please install it."
fi

ARCH=$(uname -m)
case "$ARCH" in
  x86_64)
    ARCH="x86_64"
    ;;
  aarch64)
    ARCH="arm64"
    ;;
  *)
    error_exit "Unsupported CPU architecture: $ARCH"
    ;;
esac

API_URL="https://api.github.com/repos/$REPO/releases/latest"

RELEASES_JSON=$(curl -s "$API_URL")

DOWNLOAD_URL=$(echo "$RELEASES_JSON" | grep "browser_download_url" | grep -E "scruticode-linux-$ARCH\.tar\.gz" | cut -d : -f 2,3 | tr -d '"' | tr -d ' ' | tr -d ',')

if [ -z "$DOWNLOAD_URL" ]; then
  error_exit "Could not find a compatible Scruticode binary for ($ARCH) in the latest release."
fi

echo "Downloading and installing the binary from GitHub..."
TEMP_DIR=$(mktemp -d)
TEMP_FILE="$TEMP_DIR/scruticode.tar.gz"

curl -L -o "$TEMP_FILE" "$DOWNLOAD_URL" || error_exit "Failed to download file."
tar -xzf "$TEMP_FILE" -C "$TEMP_DIR" || error_exit "Failed to extract file."

BINARY_NAME="scruticode"
BINARY_PATH="$TEMP_DIR/$BINARY_NAME"
INSTALL_PATH="$INSTALL_DIR/$BINARY_NAME"

if [ ! -f "$BINARY_PATH" ]; then
  error_exit "Binary file '$BINARY_NAME' not found after extraction."
fi

sudo mv "$BINARY_PATH" "$INSTALL_PATH" || error_exit "Failed to move binary. You may need sudo permissions."
sudo chmod +x "$INSTALL_PATH" || error_exit "Failed to set execution permissions."

rm -rf "$TEMP_DIR"

echo "Scruticode was installed successfully!"
echo "You can now run 'scruticode' in your terminal."

exit 0