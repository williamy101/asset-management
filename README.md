# Asset Management API

## Preview ERD 
![image](https://github.com/user-attachments/assets/a5243dbf-b45c-49ad-a926-0ce5697b3ab2)
*Note: Untuk melihat rancangan database lebih jelasnya di branch main*
(Link: https://dbdiagram.io/d/ERD-Asset-Management-676dac255406798ef7b59c4e)

## Slide Presentasi 
## ![image](https://github.com/user-attachments/assets/876e6272-ddb8-4e9a-b5dd-ab4a1c88d037)
(Link: https://www.canva.com/design/DAGcmM3nHpE/RlvXGsFwEORdt432r6Bd3A/edit?utm_content=DAGcmM3nHpE&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)


 ### Aplikasi backend untuk mengelola aset, termasuk operasi CRUD, kontrol akses berbasis role, dan manajemen pengguna.

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
