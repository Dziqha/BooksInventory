
# BooksInventory

BooksInventory adalah aplikasi manajemen inventaris buku yang dibangun menggunakan Go (Golang) dengan framework Fiber. Aplikasi ini menggunakan Redis sebagai cache, MySQL sebagai database, dan menyediakan API untuk operasi CRUD (Create, Read, Update, Delete) buku.

## Fitur

- **Manajemen Buku**: Menambahkan, melihat, mengedit, dan menghapus buku.
- **Middleware**: Logger, CORS, dan Recovery untuk penanganan error.
- **Autentikasi API Key**: Melalui middleware khusus.
- **Caching**: Menggunakan Redis untuk caching data.

## Persyaratan

- Go 1.16 atau lebih baru
- Redis
- MySQL
- Git

## Instalasi

1. **Clone repository ini:**

   ```bash
   git clone https://github.com/your-username/BooksInventory.git
   cd BooksInventory
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Buat file `.env` di root project:**

   Isi file `.env` dengan konfigurasi berikut:

   ```env
   API_KEY=your-api-key
   MYSQL_USER=root
   MYSQL_PASSWORD=your-password
   MYSQL_DB=books_db
   MYSQL_HOST=localhost
   MYSQL_PORT=3306
   ```

4. **Jalankan aplikasi:**

   ```bash
   go run main.go
   ```

   Aplikasi akan berjalan pada `localhost:3000`.

## Penggunaan

API ini menyediakan beberapa endpoint untuk mengelola buku. Gunakan alat seperti Postman atau cURL untuk mengakses API.

Contoh endpoint:

- **GET /books**: Mendapatkan daftar buku
- **POST /books**: Menambahkan buku baru
- **PUT /books/:id**: Memperbarui informasi buku
- **DELETE /books/:id**: Menghapus buku

## Strukur Direktori

- `app/`: Konfigurasi aplikasi seperti database.
- `src/controllers/`: Logika bisnis dan kontrol alur data.
- `src/services/`: Logika utama untuk memanipulasi data.
- `src/routers/`: Definisi routing API.
- `src/middleware/`: Middleware khusus.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repository ini dan buat pull request dengan penjelasan tentang perubahan yang Anda lakukan.

## Lisensi

Proyek ini dilisensikan di bawah MIT License. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.
