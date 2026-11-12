#!/usr/bin/env bash
# Learn CLI Unix/Linux/WSL Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/<username>/learnbydoing/main/install.sh | bash

set -e

echo "🚀 Memulai instalasi Learn CLI..."

INSTALL_DIR="$HOME/.learn/bin"
EXE_PATH="$INSTALL_DIR/learn"

mkdir -p "$INSTALL_DIR"

REPO="ahmads/learnbydoing"
DOWNLOAD_URL="https://github.com/$REPO/releases/latest/download/learn-linux-amd64"

echo "⬇️  Mengunduh binary learn..."
if curl -fsSL "$DOWNLOAD_URL" -o "$EXE_PATH"; then
    chmod +x "$EXE_PATH"
else
    echo "⚠️  Gagal mengunduh dari release online ($DOWNLOAD_URL)."
    echo "   Pastikan file 'learn-linux-amd64' sudah di-upload ke GitHub Releases."
fi

# Add to PATH in shell profile if not present
SHELL_RC=""
if [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
elif [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
fi

if [ -n "$SHELL_RC" ] && [ -f "$SHELL_RC" ]; then
    if ! grep -q "$INSTALL_DIR" "$SHELL_RC"; then
        echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_RC"
        echo "⚙️  Menambahkan $INSTALL_DIR ke $SHELL_RC"
    fi
fi

export PATH="$PATH:$INSTALL_DIR"

echo "🔍 Memeriksa compiler gcc..."
if command -v gcc >/dev/null 2>&1; then
    echo "✓ GCC terdeteksi! Sistem siap digunakan."
else
    echo "⚠️  Peringatan: 'gcc' belum terdeteksi. Jalankan 'sudo apt install build-essential'."
fi

echo ""
echo "🎉 Instalasi selesai! Jalankan 'source $SHELL_RC' atau buka terminal baru, lalu ketik 'learn'."
