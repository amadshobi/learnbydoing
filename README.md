# LEARN CLI 🚀
> Platform Latihan Mandiri C/C++ Berbasis Staged TDD untuk Mahasiswa Semester 1.

`learn` adalah CLI lokal yang terinspirasi dari CodeCrafters, dirancang khusus untuk memandu mahasiswa baru belajar pemrograman C/C++ tahap demi tahap dengan umpan balik otomatis langsung di terminal.

---

## ✨ Fitur Utama
- **100% Offline & Portable:** Menggunakan single binary Go (`learn.exe`) dengan kurikulum, starter code, dan test case ter-embed langsung di dalamnya (`//go:embed`). Tidak butuh koneksi internet atau server eksternal.
- **Dukungan C & C++:** Otomatis mendeteksi file kode mahasiswa (`main.c` menggunakan `gcc` atau `main.cpp` menggunakan `g++`).
- **Pengujian Stdin/Stdout Dua Lapis:**
  - `learn test`: Menguji dengan sample test terbuka beserta visualisasi diff jika salah.
  - `learn submit`: Menguji validasi penuh (termasuk hidden test) dan mencatat stage selesai.
- **Proteksi Anti-Infinite-Loop:** Proses eksekusi dibatasi timeout 3 detik untuk mencegah terminal macet akibat perulangan `while`/`for` tanpa henti.
- **Pesan Ramah Mahasiswa:** Semua log instruksi, tips compiler, dan diff disajikan dalam Bahasa Indonesia yang komunikatif.

---

## 🛠️ Cara Build Binary (Untuk Dibagikan ke Teman)

### 1. Build untuk Windows (`learn.exe`):
```bash
go build -o learn.exe .
```

### 2. Cross-Compile dari Linux/WSL untuk Windows:
```bash
GOOS=windows GOARCH=amd64 go build -o learn.exe .
```

Kamu tinggal membagikan file `learn.exe` tersebut ke teman-teman kuliahmu!

---

## 💻 Panduan Penggunaan Mahasiswa

### Langkah 1: Persiapan
Pastikan laptop sudah terinstall compiler C/C++ (misal: MinGW-w64 di Windows) sehingga perintah `gcc` atau `g++` dapat dipanggil di terminal.

### Langkah 2: Mulai Tantangan
Letakkan `learn.exe` di folder belajar kamu, lalu buka terminal (CMD / PowerShell / Bash):
```bash
# Menampilkan panduan bantuan
learn --help

# Melihat daftar stage tantangan yang tersedia
learn stage list

# Membuka stage pertama (Kalkulator Sederhana)
learn init 1
```

> **Catatan `--force`:** Jika kodinganmu error parah dan kamu ingin mereset `main.c` kembali ke template bersih awal dari nol, gunakan:
> ```bash
> learn init 1 --force
> ```
> Ini adalah tombol RESET aman, bukan perintah berbahaya.

### Langkah 3: Mengerjakan Tantangan
Perintah di atas akan membuat folder `learn-c/easy-stage-1/`.
1. Masuk ke folder stage:
   ```bash
   cd learn-c/easy-stage-1
   ```
2. Baca panduan materi dan spesifikasi di `README.md`.
3. Buka file `main.c` di IDE atau text editor favorit kamu (VS Code, Dev-C++, Notepad++, dll) dan tulis solusinya.

### Langkah 4: Uji Coba Solusi
Ketik kapan saja di terminal:
```bash
learn test
```
Jika ada kesalahan logika atau kompilasi, CLI akan menampilkan pesan ramah serta perbandingan output yang diharapkan.

### Langkah 5: Kumpulkan dan Buka Stage Berikutnya
Jika semua sample test sudah lulus:
```bash
learn submit
```
CLI akan memvalidasi skenario lanjutan. Jika berhasil, progresmu akan dicatat dan kamu akan dipandu untuk membuka stage berikutnya:
```bash
learn init 2 --easy
```

### Cek Progres
Kapan saja kamu ingin melihat progres stage yang sudah kamu selesaikan:
```bash
learn status
```
