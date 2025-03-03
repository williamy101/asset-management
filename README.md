# Asset Management API

Aplikasi backend untuk mengelola aset, termasuk operasi CRUD, kontrol akses berbasis peran, dan manajemen pengguna.
## *Untuk melihat rancangan database di branch main*

## Fitur

- Pendaftaran pengguna dengan peran (admin, user)
- Autentikasi pengguna dengan token JWT
- Kontrol akses berbasis peran untuk berbagai route
- Operasi CRUD untuk aset, kategori, dan status
- Manajemen pemeliharaan aset (Ongoing)
- Keamanan: Peran Admin dapat memberikan atau memodifikasi peran pengguna (Ongoing)

## Stack Teknologi

- **Go (Golang)**: Bahasa pemrograman backend
- **Gin-Gonic**: Framework web untuk Go
- **GORM**: ORM untuk manajemen database
- **MySQL**: Database
- **JWT**: JSON Web Tokens untuk autentikasi

```bash
git clone https://github.com/williamy101/asset-management.git
cd asset-management
