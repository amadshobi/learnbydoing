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
$Repo = "amadshobi/learnbydoing"
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

# 4. Tambahkan ke PowerShell Profile agar tab/terminal baru (termasuk VS Code) langsung mengenali 'learn'
$ProfilePaths = @(
    (Join-Path ([Environment]::GetFolderPath("MyDocuments")) "WindowsPowerShell\Microsoft.PowerShell_profile.ps1"),
    (Join-Path ([Environment]::GetFolderPath("MyDocuments")) "PowerShell\Microsoft.PowerShell_profile.ps1")
)

foreach ($Prof in $ProfilePaths) {
    try {
        $ProfDir = Split-Path $Prof
        if (-not (Test-Path $ProfDir)) {
            New-Item -ItemType Directory -Path $ProfDir -Force | Out-Null
        }
        $ProfileSnippet = "`n# Learn CLI PATH`nif (`$env:Path -notlike '*$InstallDir*') { `$env:Path = `"$InstallDir;`" + `$env:Path }`n"
        if (-not (Test-Path $Prof) -or (Get-Content $Prof -Raw) -notlike "*$InstallDir*") {
            Add-Content -Path $Prof -Value $ProfileSnippet -Force
        }
    } catch {
        # Abaikan jika ada issue permission profile
    }
}

# 5. Refresh PATH di sesi PowerShell saat ini
$env:Path = "$InstallDir;$env:Path"

# 6. Cek ketersediaan gcc / g++
Write-Host "🔍 Memeriksa compiler C/C++ (gcc)..." -ForegroundColor Cyan
if (Get-Command gcc -ErrorAction SilentlyContinue) {
    Write-Host "✓ GCC terdeteksi! Sistem siap digunakan." -ForegroundColor Green
} else {
    Write-Host "⚠️  Peringatan: 'gcc' belum terdeteksi di PATH." -ForegroundColor Yellow
    Write-Host "   Pastikan kamu sudah menginstall MinGW-w64 agar bisa mengompilasi kode C/C++." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "🎉 Instalasi selesai! Kamu sekarang bisa mengetik 'learn' di terminal mana saja." -ForegroundColor Green
Write-Host "💡 Catatan: Jika jendela terminal atau VS Code sudah terbuka sebelum install, cukup tutup dan buka kembali (restart terminal) agar PATH ter-refresh." -ForegroundColor Yellow
Write-Host "👉 Coba ketik: learn --help" -ForegroundColor Cyan
