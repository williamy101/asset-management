# Asset Management API

## Preview ERD
![image](https://github.com/user-attachments/assets/a5243dbf-b45c-49ad-a926-0ce5697b3ab2)
*Note: Untuk melihat rancangan database lebih jelasnya di branch main*


 ## Aplikasi backend untuk mengelola aset, termasuk operasi CRUD, kontrol akses berbasis peran, dan manajemen pengguna.

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

## Cara menjalankan project
```bash
git clone https://github.com/williamy101/asset-management.git
cd asset-management
