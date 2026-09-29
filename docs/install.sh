#!/usr/bin/env bash
set -e

echo "🚀 Installing devports..."

OS="$(uname -s)"
ARCH="$(uname -m)"
VERSION="v1.0.0"
REPO="selomrani/devports"

if [ "$OS" = "Linux" ]; then
    if [ "$ARCH" = "x86_64" ] || [ "$ARCH" = "amd64" ]; then
        FILE="devports-linux-amd64.tar.gz"
        BIN="devports-linux-amd64"
    else
        echo "❌ Unsupported Linux architecture: $ARCH (Currently only x86_64/amd64 is supported)"
        exit 1
    fi
elif [ "$OS" = "Darwin" ]; then
    if [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
        FILE="devports-darwin-arm64.tar.gz"
        BIN="devports-darwin-arm64"
    else
        echo "❌ Currently only Apple Silicon (M1/M2/M3) is supported on macOS."
        exit 1
    fi
else
    echo "❌ Unsupported OS: $OS. Please install via 'go install github.com/selomrani/devports@latest'"
    exit 1
fi

URL="https://github.com/$REPO/releases/download/$VERSION/$FILE"

echo "⬇️  Downloading $VERSION..."
curl -sL "$URL" -o "$FILE"

echo "📦 Extracting..."
tar -xzf "$FILE"

echo "🔧 Installing to /usr/local/bin/devports (this may prompt for your password)..."
chmod +x "$BIN"
sudo mv "$BIN" /usr/local/bin/devports

echo "🧹 Cleaning up..."
rm "$FILE"

echo ""
echo "✅ devports installed successfully! 🎉"
echo "👉 Just type 'devports' in your terminal to launch it."
