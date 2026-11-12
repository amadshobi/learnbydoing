# Stage 1: Kalkulator Sederhana

Selamat datang di tantangan pertama! Di tahap ini, kamu akan membuat program kalkulator interaktif sederhana dengan bahasa C/C++.

## Konsep yang Dipelajari
1. Input data menggunakan `scanf()`
2. Output data menggunakan `printf()`
3. Variabel bertipe `int` atau `float`
4. Operasi aritmatika (`+`, `-`, `*`, `/`)

## Spesifikasi Interaksi Program
Program harus menampilkan pesan interaktif persis seperti di bawah ini:

1. Cetak header:
   ```text
   === KALKULATOR MINI ===
   ```
2. Minta input angka pertama:
   ```text
   Masukkan angka pertama: <angka1>
   ```
3. Minta input angka kedua:
   ```text
   Masukkan angka kedua: <angka2>
   ```
4. Minta input operator (`+`, `-`, `*`, atau `/`):
   ```text
   Pilih operasi (+, -, *, /): <operator>
   ```
5. Tampilkan hasil perhitungan:
   ```text
   Hasil: <nilai>
   ```
   *Catatan:* Jika pembagian dengan angka 0 dilakukan (`operator == '/'` dan `angka2 == 0`), tampilkan:
   ```text
   Error: Tidak bisa membagi dengan nol!
   ```

## Contoh Interaksi Terminal
```text
=== KALKULATOR MINI ===
Masukkan angka pertama: 10
Masukkan angka kedua: 5
Pilih operasi (+, -, *, /): +
Hasil: 15
```

## Cara Menjalankan Tes
Ketik di terminal:
```bash
learn test
```
Jika sudah yakin dan ingin menyimpan progress, ketik:
```bash
learn submit
```
