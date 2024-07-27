
# BooksInventory

BooksInventory adalah aplikasi manajemen inventaris buku yang dibangun menggunakan Go (Golang) dengan framework Fiber. Aplikasi ini menggunakan Redis sebagai cache, MySQL sebagai database, dan menyediakan API untuk operasi CRUD (Create, Read, Update, Delete) buku.

## Fitur

- **Manajemen Buku**: Menambahkan, melihat, mengedit, dan menghapus buku.
- **Middleware**: Logger, CORS, dan Recovery untuk penanganan error.
- **Autentikasi API Key**: Melalui middleware khusus.
- **Caching**: Menggunakan Redis untuk caching data.

## Persyaratan

- Go 1.16 atau lebih baru
- Air (untuk hot-reload)
- Redis
- MySQL
- Git

## Instalasi

1. **Clone repository ini:**

   ```bash
   git clone https://github.com/Dziqha/BooksInventory.git

   cd BooksInventory
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Install Air (hot-reload tool):**

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

   Pastikan $GOPATH/bin ada di dalam $PATH untuk menjalankan air.


4. **Buat file `.env` di root project:**

   Isi file `.env` dengan konfigurasi berikut:

   ```env
   API_KEY=your-api-key
   MYSQL_USER=root
   MYSQL_PASSWORD=your-password
   MYSQL_DB=books_db
   MYSQL_HOST=localhost
   MYSQL_PORT=3306
   ```

5. **Jalankan aplikasi:**

   ```bash
   air
   ```

   Aplikasi akan berjalan pada `localhost:3000`.

## Penggunaan

API ini menyediakan beberapa endpoint untuk mengelola buku. Gunakan alat seperti Postman atau cURL untuk mengakses API.

Contoh endpoint:

- **GET /api/v1/books**: Mendapatkan daftar buku
- **POST /api/v1/books**: Menambahkan buku baru
- **PUT /api/v1/books/:id**: Memperbarui informasi buku
- **DELETE /api/v1/books/:id**: Menghapus buku

## Strukur Direktori

- `app/`: Konfigurasi aplikasi seperti database.
- `src/controllers/`: Logika bisnis dan kontrol alur data.
- `src/services/`: Logika utama untuk memanipulasi data.
- `src/routers/`: Definisi routing API.
- `src/middleware/`: Middleware khusus.
- `tmp/`: Informasi log atau debug dari fitur hot-reloading.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repository ini dan buat pull request dengan penjelasan tentang perubahan yang Anda lakukan.

## Lisensi

Proyek ini dilisensikan di bawah MIT License. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.
