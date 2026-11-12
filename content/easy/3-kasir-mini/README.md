# Stage 3: Kasir Mini

Di tahap ketiga ini, kamu akan mengasah perulangan (`while` atau `for`) dan teknik akumulasi total nilai.

## Konsep yang Dipelajari
1. Perulangan `while` atau `for`
2. Pola akumulator (`total += harga`)
3. Kondisi berhenti (sentinel value `0`)

## Spesifikasi Interaksi Program
1. Cetak header:
   ```text
   === KASIR MINI ===
   ```
2. Minta input harga barang secara berulang:
   ```text
   Masukkan harga barang (0 untuk selesai): <harga>
   ```
3. Jika angka yang dimasukkan adalah `0`, perulangan berhenti.
4. Tampilkan total akhir belanja:
   ```text
   Total Belanja: Rp <total>
   ```

## Contoh Interaksi Terminal
```text
=== KASIR MINI ===
Masukkan harga barang (0 untuk selesai): 15000
Masukkan harga barang (0 untuk selesai): 25000
Masukkan harga barang (0 untuk selesai): 0
Total Belanja: Rp 40000
```

## Cara Menjalankan Tes
```bash
learn test
```
```bash
learn submit
```
