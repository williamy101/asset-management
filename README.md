# Asset Management API
### Sistem backend untuk mengelola aset, termasuk operasi CRUD, kontrol akses berbasis role, dan manajemen pengguna.

***Terdapat fitur transaksi baru yang belum dilakukan testing yaitu***:
 - Request pinjam aset (user create request pinjam aset, admin bisa approve atau reject)
 - Tabel penyimpanan aset yang dipinjam
 - Request maintenance (request dibuat oleh user, admin bisa approve atau reject, create maintenance manual tetap ada untuk maintenance rutin)
 
> _Langkah selanjutnya: testing dan dokumentasi fitur baru, integrasi project ke docker, lalu deployment ke cloud hosting_ 

## Preview ERD 
![image](https://github.com/user-attachments/assets/a98a8631-191a-477c-adb1-9b6673eee1a6)
> *Note: Untuk melihat rancangan database lebih jelasnya di branch main*

(Link: https://dbdiagram.io/d/ERD-Asset-Management-676dac255406798ef7b59c4e)

## Slide Presentasi 
## ![image](https://github.com/user-attachments/assets/876e6272-ddb8-4e9a-b5dd-ab4a1c88d037)
(Link: https://www.canva.com/design/DAGcmM3nHpE/RlvXGsFwEORdt432r6Bd3A/edit?utm_content=DAGcmM3nHpE&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)




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
