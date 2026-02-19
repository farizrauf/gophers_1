# Gophers_1

Project ini dibuat untuk belajar dan eksplorasi bahasa Go.  
Tujuan: membangun aplikasi sederhana dengan struktur modular dan mudah dikembangkan.

## Fitur
- **CheckChange**: fungsi untuk mengecek apakah uang pembeli cukup untuk membeli barang.
  - Jika uang kurang, sistem menolak transaksi dan menampilkan pesan merah.
  - Jika uang cukup, sistem menerima transaksi dan menampilkan pesan hijau dengan jumlah kembalian.
- Menggunakan library [fatih/color](https://github.com/fatih/color) untuk output berwarna di terminal.

## Instalasi
Clone repository ini:
```bash
git clone https://github.com/farizrauf/gophers_1.git
cd gophers_1
go run main.go
