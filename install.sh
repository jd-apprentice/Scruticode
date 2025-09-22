#!/bin/bash

# --- Configuration ---
REPO="jd-apprentice/Scruticode"
INSTALL_DIR="/usr/local/bin"
DOWNLOAD_URL=""
CMD_TO_INSTALL="scruticode"

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

OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
  Linux)
    OS_NAME="Linux"
    ;;
  Darwin)
    OS_NAME="Darwin"
    ;;
  CYGWIN*|MINGW*|MSYS*)
    OS_NAME="Windows"
    ;;
  *)
    error_exit "Unsupported operating system: $OS"
    ;;
esac

case "$ARCH" in
  x86_64)
    ARCH_NAME="x86_64"
    ;;
  aarch64|arm64)
    ARCH_NAME="arm64"
    ;;
  i386|i686)
    ARCH_NAME="i386"
    ;;
  *)
    error_exit "Unsupported CPU architecture: $ARCH"
    ;;
esac

API_URL="https://api.github.com/repos/$REPO/releases/latest"
RELEASES_JSON=$(curl -s "$API_URL")

if [ "$OS_NAME" = "Windows" ]; then
  FILE_EXT="zip"
  CMD_TO_INSTALL="scruticode.exe"
else
  FILE_EXT="tar.gz"
fi

DOWNLOAD_URL=$(echo "$RELEASES_JSON" | grep "browser_download_url" | grep -E "Scruticode_${OS_NAME}_${ARCH_NAME}\.${FILE_EXT}" | cut -d : -f 2,3 | tr -d '"' | tr -d ' ' | tr -d ',')

if [ -z "$DOWNLOAD_URL" ]; then
  error_exit "Could not find a compatible Scruticode binary for ($OS_NAME/$ARCH_NAME) in the latest release."
fi

echo "Downloading and installing the binary from GitHub..."
TEMP_DIR=$(mktemp -d)
TEMP_FILE="$TEMP_DIR/scruticode.$FILE_EXT"

curl -L -o "$TEMP_FILE" "$DOWNLOAD_URL" || error_exit "Failed to download file."

if [ "$FILE_EXT" = "zip" ]; then
  if ! command_exists "unzip"; then
    error_exit "'unzip' command not found. Please install it."
  fi
  unzip -o "$TEMP_FILE" -d "$TEMP_DIR" || error_exit "Failed to extract file."
else
  if ! command_exists "tar"; then
    error_exit "'tar' command not found. Please install it."
  fi
  tar -xzf "$TEMP_FILE" -C "$TEMP_DIR" || error_exit "Failed to extract file."
fi

BINARY_PATH="$TEMP_DIR/Scruticode"
if [ "$OS_NAME" = "Windows" ]; then
  BINARY_PATH="${BINARY_PATH}.exe"
fi

if [ ! -f "$BINARY_PATH" ]; then
  error_exit "Binary file not found after extraction."
fi

if [ "$OS_NAME" = "Windows" ]; then
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"
  INSTALL_PATH="$INSTALL_DIR/scruticode.exe"
  mv "$BINARY_PATH" "$INSTALL_PATH" || error_exit "Failed to move binary."
  echo "Scruticode was installed successfully to $INSTALL_PATH"
  echo "Please ensure '$INSTALL_DIR' is in your system's PATH."
fi

if [ "$OS_NAME" = "Linux" ] || [ "$OS_NAME" = "Darwin" ]; then
  INSTALL_PATH="$INSTALL_DIR/scruticode"
  sudo mv "$BINARY_PATH" "$INSTALL_PATH" || error_exit "Failed to move binary. You may need sudo permissions."
  sudo chmod +x "$INSTALL_PATH" || error_exit "Failed to set execution permissions."
  echo "Scruticode was installed successfully!"
  echo "You can now run 'scruticode' in your terminal."
fi

rm -rf "$TEMP_DIR"

exit 0
