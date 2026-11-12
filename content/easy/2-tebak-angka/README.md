# Stage 2: Tebak Angka

Di tahap kedua ini, kamu akan melatih logika percabangan kondisi menggunakan `if`, `else if`, dan `else`.

## Konsep yang Dipelajari
1. Pernyataan `if`, `else if`, `else`
2. Operator pembanding (`>`, `<`, `==`)
3. Alur percabangan keputusan

## Spesifikasi Interaksi Program
Angka rahasia yang harus ditebak pada tantangan ini adalah **50**.

1. Cetak header:
   ```text
   === GAME TEBAK ANGKA ===
   ```
2. Minta input tebakan angka:
   ```text
   Masukkan tebakan angka (1-100): <angka>
   ```
3. Evaluasi tebakan:
   - Jika tebakan lebih besar dari 50, cetak:
     ```text
     Terlalu besar!
     ```
   - Jika tebakan lebih kecil dari 50, cetak:
     ```text
     Terlalu kecil!
     ```
   - Jika tebakan sama dengan 50, cetak:
     ```text
     Selamat, tebakanmu benar!
     ```

## Contoh Interaksi Terminal
```text
=== GAME TEBAK ANGKA ===
Masukkan tebakan angka (1-100): 25
Terlalu kecil!
```

## Cara Menjalankan Tes
```bash
learn test
```
```bash
learn submit
```
