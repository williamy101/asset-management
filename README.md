# Asset Management API
### Aplikasi backend untuk mengelola aset, termasuk operasi CRUD, kontrol akses berbasis role, dan manajemen pengguna.

### **3 Fungsi Untested:**
- maintenance request: User dapat create request maintenance aset tertentu, admin dapat approve atau reject, jika approve admin dapat menginput teknisi yang mengerjakan dll untuk membuat jadwal maintenance
- borrow asset request: User dapat request pinjam aset, admin dapat approve atau reject, jika approve maka akan terdata di tabel borrowed asset, dan aset tersebut terbind oleh userID
- borrowed asset: Tabel baru yang menyimpan data aset yang sedang dipinjam user

_In Progress & Coming Soon:_ Testing dan dokumentasi fitur baru menggunakan Postman, Integrasi Docker, Deployment ke cloud hosting Google. Integrasi Redis untuk Notification, Rate Limiting, Session Storage. 

>Link Postman: https://universal-desert-823258-1.postman.co/workspace/Asset-Management~d0881856-648f-4073-99bc-54a043912a33/collection/26349837-64b9586a-fdb5-4ec5-8a5b-900f407813dc?action=share&creator=26349837

## Preview ERD 
![image](https://github.com/user-attachments/assets/936fd75a-4908-4bf6-a3dd-8dd22e25301c)

*Note: Untuk melihat rancangan database lebih jelasnya di branch main*

>Link: https://dbdiagram.io/d/ERD-Asset-Management-676dac255406798ef7b59c4e

## Slide Presentasi 
## ![image](https://github.com/user-attachments/assets/876e6272-ddb8-4e9a-b5dd-ab4a1c88d037)
>Link: https://www.canva.com/design/DAGcmM3nHpE/RlvXGsFwEORdt432r6Bd3A/edit?utm_content=DAGcmM3nHpE&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton




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
