# Belajar Go REST API
Belajar REST API dengan bahasa pemrograman Go.

# Daftar Isi
* [Fitur](#fitur)
* [Konsep Arsitektural](#konsep-arsitektural)
* [Stack Teknologi](#stack-teknologi)
* [Utilitas Pihak Ketiga](#utilitas-pihak-ketiga)
* [Menjalankan Projek](#menjalankan-projek)
  * [Persiapan](#persiapan)
  * [Mode Development](#mode-development)
  * [Mode Production](#mode-production)

# Fitur
Projek ini sudah siap dijalankan dan dapat dimodifikasi sesuai kebutuhan masing-masing. Adapun kemampuan projek ini:
1. Dapat menjadi backend dari REST API.
2. Dapat menjadi service dalam rangkaian microservices.
3. Dapat menyediakan dokumentasi API melalui Swagger.
4. Support berbagai koneksi database, seperti MySQL, PostgreSQL, SQLite, dan SQL Server.
5. Support migrasi database dalam bentuk file SQL.
6. Support [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection).

# Konsep Arsitektural
Struktur proyek ini mengikuti konsep-konsep arsitektural seperti:
- DRY (Don't Repeat Yourself)
- SOLID Principle
- Clean Architecture

# Stack Teknologi
Nama | Peran
-|-
[Go](https://golang.org) | Bahasa Pemrograman
[Fiber](https://docs.gofiber.io) | Web Framework
[PostgreSQL](https://www.postgresql.org) | Database Management System
[Gorm](https://gorm.io/index.html) | Object-Relational Mapper
[Jwt](https://jwt.io) | Token untuk Otentikasi

# Utilitas Pihak Ketiga
Nama | Peran
-|-
[Validator](https://github.com/go-playground/validator) | Validasi data
[Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) | Migrasi database
[Swag](https://github.com/swaggo/swag) | Otomatis generate dokumentasi REST API (swagger)
[GNU/Make](https://www.gnu.org/software/make/) | Build tool

# Menjalankan Projek
## Persiapan
Agar projek dapat dijalankan, diperlukan beberapa hal untuk dipersiapkan, antara lain:
1. Instalasi Fiber CLI.
```bash
go get -u github.com/gofiber/cli/fiber
```
2. Instalasi Swaggo.
```bash
go get -u github.com/swaggo/swag/cmd/swag
```
3. Instalasi Migrate.
Ikuti panduan [di sini](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

4. Menjalankan migration dengan menyesuaikan akses database.
```bash
migrate -path db/migrations -database "postgres://localhost/test_db?sslmode=disable" up
```

## Mode Development
1. Clone projek ke komputer lokal, dengan perintah:
  ```bash
  git clone https://github.com/hadihammurabi/belajar-go-rest-api
  ```
2. Menjalankan projek dengan mode development (pengembangan) dengan perintah:
  ```bash
  make dev
  ```

## Mode Production
1. Clone projek ke komputer lokal, dengan perintah:
  ```bash
  git clone https://github.com/hadihammurabi/belajar-go-rest-api
  ```
2. Menjalankan projek siap guna dengan perintah:
  ```bash
  make && ./main
  ```
  
