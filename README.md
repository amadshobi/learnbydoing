# LEARNBYDOING CLI 🚀
> Platform Latihan Mandiri Pemrograman Berbasis Staged TDD untuk Mahasiswa Semester 1.

`learn` adalah CLI lokal yang terinspirasi dari CodeCrafters, dirancang khusus untuk memandu mahasiswa baru belajar pemrograman (C/C++ dan ke depannya bahasa lain) tahap demi tahap dengan umpan balik otomatis langsung di terminal.

```text
█░░ █▀▀ ▄▀█ █▀█ █▄░█                  ▄█▄
█▄▄ ██▄ █▀█ █▀▄ █░▀█                ▄█████▄     │
                                  ▄█████████▄   │
──────── STAGED LEARNING ──────── ▀█████████▀  (o)
                                    ███████
```

---

## 📥 Cara Instalasi Otomatis (Cukup 1 Baris di Terminal)

Teman-teman tidak perlu repot download manual atau atur Environment PATH sendiri. Cukup jalankan perintah instalasi di bawah ini:

### Untuk Windows (PowerShell):
Buka PowerShell, lalu ketik/paste:
```powershell
irm https://raw.githubusercontent.com/amadshobi/learnbydoing/main/install.ps1 | iex
```

### Untuk Linux / WSL / macOS (Bash):
Buka terminal Bash, lalu ketik/paste:
```bash
curl -fsSL https://raw.githubusercontent.com/amadshobi/learnbydoing/main/install.sh | bash
```

> **Catatan:** Script installer otomatis menempatkan binary ke folder `~/.learn/bin` dan mendaftarkannya ke PATH. Selesai instalasi, perintah `learn` langsung bisa dipanggil di direktori mana saja!

---

## ✨ Fitur Utama
- **100% Offline & Portable:** Menggunakan single binary Go (`learn.exe`) dengan kurikulum, starter code, dan test case ter-embed langsung di dalamnya (`//go:embed`). Tidak butuh koneksi internet atau server eksternal saat mengerjakan soal.
- **Dukungan C & C++:** Otomatis mendeteksi file kode mahasiswa (`main.c` menggunakan `gcc` atau `main.cpp` menggunakan `g++`).
- **Pengujian Stdin/Stdout Dua Lapis:**
  - `learn test`: Menguji dengan sample test terbuka beserta visualisasi diff jika salah.
  - `learn submit`: Menguji validasi penuh (termasuk hidden test) dan mencatat stage selesai.
- **Proteksi Anti-Infinite-Loop:** Proses eksekusi dibatasi timeout 3 detik untuk mencegah terminal macet akibat perulangan `while`/`for` tanpa henti.
- **Pesan Ramah Mahasiswa:** Semua log instruksi, tips compiler, dan diff disajikan dalam Bahasa Indonesia yang komunikatif.

---

## 💻 Panduan Penggunaan

### 1. Persiapan Compiler
Pastikan laptop sudah terinstall compiler C/C++ (misal: MinGW-w64 di Windows atau `build-essential` di Linux) sehingga perintah `gcc` atau `g++` dapat dipanggil di terminal.

### 2. Mulai Tantangan
Buka terminal (CMD / PowerShell / Bash) di folder tempat kamu ingin belajar:
```bash
# Menampilkan panduan lengkap
learn --help

# Melihat daftar semua stage yang tersedia
learn stage list

# Membuka stage pertama (Kalkulator Sederhana)
learn init 1
```

> **💡 Catatan Flag `--force`:**  
> Jika kodinganmu error parah dan kamu ingin mereset `main.c` kembali ke template bersih awal dari nol, gunakan:
> ```bash
> learn init 1 --force
> ```
> Ini adalah tombol RESET aman, bukan perintah berbahaya.

### 3. Mengerjakan Tantangan
Perintah `learn init 1` akan membuat folder `learn-c/easy-stage-1/`:
1. Masuk ke folder stage:
   ```bash
   cd learn-c/easy-stage-1
   ```
2. Baca materi dan spesifikasi interaksi di `README.md`.
3. Buka file `main.c` di editor favoritmu (VS Code, Dev-C++, Notepad++, dll) dan tulis solusinya.

### 4. Uji Coba Solusi
Ketik kapan saja di terminal:
```bash
learn test
```
Jika ada kesalahan logika atau kompilasi, CLI akan menampilkan pesan ramah serta perbandingan output yang diharapkan vs output kodemu.

### 5. Kumpulkan dan Buka Stage Berikutnya
Jika semua sample test sudah lulus:
```bash
learn submit
```
CLI akan menguji kasus rahasia (hidden test). Jika berhasil, progresmu akan disimpan dan kamu akan diarahkan ke stage berikutnya:
```bash
learn init 2 --easy
```

### 6. Cek Progres Belajar
Kapan saja kamu ingin melihat stage yang sudah berhasil diselesaikan:
```bash
learn status
```

---

## 🛠️ Build Mandiri dari Source Code (Opsional)

Jika ingin mengompilasi binary sendiri dari source code:
```bash
# Build untuk Windows
go build -o learn.exe .

# Cross-compile dari Linux/WSL untuk Windows
GOOS=windows GOARCH=amd64 go build -o learn.exe .

# Build untuk Linux
GOOS=linux GOARCH=amd64 go build -o learn .
```
