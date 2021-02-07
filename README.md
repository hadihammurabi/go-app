# Belajar Go REST API
Belajar REST API dengan bahasa pemrograman Go.

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
  
