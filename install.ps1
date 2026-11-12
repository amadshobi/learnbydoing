# Learn CLI Windows Installer (PowerShell)
# Usage: irm https://raw.githubusercontent.com/<username>/learnbydoing/main/install.ps1 | iex

$ErrorActionPreference = "Stop"

Write-Host "🚀 Memulai instalasi Learn CLI..." -ForegroundColor Cyan

$InstallDir = Join-Path $HOME ".learn\bin"
$ExePath = Join-Path $InstallDir "learn.exe"

# 1. Buat direktori instalasi jika belum ada
if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

# 2. Download binary dari GitHub Release
# Catatan: ganti URL di bawah sesuai username repo GitHub kamu
$Repo = "ahmads/learnbydoing" # Default placeholder, akan mengikuti release repo
$DownloadUrl = "https://github.com/$Repo/releases/latest/download/learn-windows-amd64.exe"

Write-Host "⬇️  Mengunduh binary learn.exe..." -ForegroundColor Yellow
try {
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $ExePath -UseBasicParsing
} catch {
    Write-Host "⚠️  Gagal mengunduh dari release online ($DownloadUrl)." -ForegroundColor DarkYellow
    Write-Host "   Jika repo baru dibuat, pastikan file 'learn-windows-amd64.exe' sudah di-upload ke GitHub Releases." -ForegroundColor DarkYellow
}

# 3. Tambahkan ke User PATH jika belum terdaftar
$UserPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User)
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "⚙️  Menambahkan $InstallDir ke User PATH Windows..." -ForegroundColor Cyan
    $NewPath = "$UserPath;$InstallDir"
    [Environment]::SetEnvironmentVariable("Path", $NewPath, [EnvironmentVariableTarget]::User)
}

# 4. Refresh PATH di sesi PowerShell saat ini
$env:Path = "$env:Path;$InstallDir"

# 5. Cek ketersediaan gcc / g++
Write-Host "🔍 Memeriksa compiler C/C++ (gcc)..." -ForegroundColor Cyan
if (Get-Command gcc -ErrorAction SilentlyContinue) {
    Write-Host "✓ GCC terdeteksi! Sistem siap digunakan." -ForegroundColor Green
} else {
    Write-Host "⚠️  Peringatan: 'gcc' belum terdeteksi di PATH." -ForegroundColor Yellow
    Write-Host "   Pastikan kamu sudah menginstall MinGW-w64 agar bisa mengompilasi kode C/C++." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🎉 Instalasi selesai! Kamu sekarang bisa mengetik 'learn' di terminal mana saja." -ForegroundColor Green
Write-Host "👉 Coba ketik: learn --help" -ForegroundColor Cyan
