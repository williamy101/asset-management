
# **Asset Management API**

## **Preview ERD**  
![image](https://github.com/user-attachments/assets/2bda7a98-e782-4cbd-84b5-fa80f15dce49)
🔗 **[Lihat ERD lebih jelas](https://dbdiagram.io/d/ERD-Asset-Management-676dac255406798ef7b59c4e)**  

## **Slide Presentasi**  
![Slide Presentasi](https://github.com/user-attachments/assets/876e6272-ddb8-4e9a-b5dd-ab4a1c88d037)  
🔗 **[Lihat Presentasi di Canva](https://www.canva.com/design/DAGcmM3nHpE/RlvXGsFwEORdt432r6Bd3A/edit?utm_content=DAGcmM3nHpE&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)**  

## **Deskripsi**  
Aplikasi backend untuk mengelola aset, mencakup operasi CRUD, kontrol akses berbasis peran, dan manajemen pengguna.  

---

## **Fitur Utama**
✅ **Manajemen Aset**  
- CRUD untuk aset, kategori, dan status.  
- Penyimpanan informasi pemeliharaan aset.  


## **API Endpoint Documentation - Go Asset Management**
| **Endpoint** | **Method** | **Access Level** | **Description** |
|-------------|-----------|------------------|-----------------|
| **🔹 User Management** |
| `/users/register` | `POST` | Public | Mendaftarkan user baru. |
| `/users/login` | `POST` | Public | Login user dan mendapatkan token. |
| `/users/admin/` | `GET` | Admin (1) | Mendapatkan daftar semua user. |
| `/users/admin/:id` | `GET` | Admin (1) | Mendapatkan detail user berdasarkan ID. |
| `/users/admin/role` | `PUT` | Admin (1) | Mengubah role user. |
| **🔹 Role Management** |
| `/roles/` | `POST` | Admin (1) | Membuat role baru. |
| `/roles/` | `GET` | Admin (1) | Mendapatkan daftar semua role. |
| `/roles/:id` | `GET` | Admin (1) | Mendapatkan detail role berdasarkan ID. |
| `/roles/:id` | `DELETE` | Admin (1) | Menghapus role berdasarkan ID. |
| **🔹 Status Management** |
| `/statuses/` | `POST` | Admin (1) | Membuat status baru. |
| `/statuses/` | `GET` | Admin (1) | Mendapatkan daftar semua status. |
| `/statuses/:id` | `GET` | Admin (1) | Mendapatkan detail status berdasarkan ID. |
| `/statuses/:id` | `PUT` | Admin (1) | Mengupdate status berdasarkan ID. |
| `/statuses/:id` | `DELETE` | Admin (1) | Menghapus status berdasarkan ID. |
| `/statuses/user/` | `GET` | Technician (2) & User (3) | Mendapatkan daftar semua status. |
| `/statuses/user/:id` | `GET` | Technician (2) & User (3) | Mendapatkan detail status berdasarkan ID. |
| **🔹 Asset Category Management** |
| `/categories/` | `POST` | Admin (1) | Membuat kategori aset baru. |
| `/categories/` | `GET` | Admin (1) | Mendapatkan daftar semua kategori aset. |
| `/categories/:id` | `GET` | Admin (1) | Mendapatkan detail kategori aset berdasarkan ID. |
| `/categories/:id` | `PUT` | Admin (1) | Mengupdate kategori aset berdasarkan ID. |
| `/categories/:id` | `DELETE` | Admin (1) | Menghapus kategori aset berdasarkan ID. |
| **🔹 Asset Management** |
| `/assets/` | `POST` | Admin (1) | Menambahkan aset baru. |
| `/assets/` | `GET` | Admin (1) | Mendapatkan daftar semua aset. |
| `/assets/:id` | `GET` | Admin (1) | Mendapatkan detail aset berdasarkan ID. |
| `/assets/:id` | `PUT` | Admin (1) | Mengupdate aset berdasarkan ID. |
| `/assets/:id` | `DELETE` | Admin (1) | Menghapus aset berdasarkan ID (soft delete). |
| `/assets/get/` | `GET` | Technician (2) & User (3) | Mendapatkan daftar semua aset. |
| `/assets/get/:id` | `GET` | Technician (2) & User (3) | Mendapatkan detail aset berdasarkan ID. |
| **🔹 Maintenance Management** |
| `/maintenances/` | `POST` | Admin (1) | Membuat entri perawatan aset. |
| `/maintenances/` | `GET` | Admin (1) | Mendapatkan daftar semua perawatan. |
| `/maintenances/:id` | `GET` | Admin (1) | Mendapatkan detail perawatan berdasarkan ID. |
| `/maintenances/:id` | `DELETE` | Admin (1) | Menghapus entri perawatan. |
| `/maintenances/total-cost` | `GET` | Admin (1) | Mendapatkan total biaya perawatan. |
| `/maintenances/total-cost/:asset_id` | `GET` | Admin (1) | Mendapatkan total biaya perawatan untuk aset tertentu. |
| `/maintenances/technician/:id/start` | `PUT` | Technician (2) | Memulai perawatan aset. |
| `/maintenances/technician/:id/end` | `PUT` | Technician (2) | Menyelesaikan perawatan aset. |
| `/maintenances/user/` | `GET` | Technician (2) & User (3) | Mendapatkan daftar perawatan berdasarkan pekerja. |
| **🔹 Maintenance Request Management** |
| `/maintenance-requests/` | `POST` | User (3) | Mengajukan permintaan perawatan aset. |
| `/maintenance-requests/admin/:id/approve` | `PUT` | Admin (1) | Menyetujui permintaan perawatan. |
| `/maintenance-requests/admin/:id/reject` | `PUT` | Admin (1) | Menolak permintaan perawatan. |
| **🔹 Borrow Asset Request Management** |
| `/borrow-requests/` | `POST` | User (3) | Mengajukan permintaan peminjaman aset. |
| `/borrow-requests/admin/:id/approve` | `PUT` | Admin (1) | Menyetujui permintaan peminjaman. |
| `/borrow-requests/admin/:id/reject` | `PUT` | Admin (1) | Menolak permintaan peminjaman. |
| **🔹 Borrowed Asset Management** |
| `/borrowed-assets/` | `GET` | User (3) | Mendapatkan daftar aset yang sedang dipinjam. |
| `/borrowed-assets/:id` | `GET` | User (3) | Mendapatkan detail aset yang sedang dipinjam berdasarkan ID. |
| `/borrowed-assets/:id/return` | `PUT` | User (3) | Memperbarui tanggal pengembalian aset yang dipinjam. |

## **Hak Akses (Authorization)**
| **Role** | **Level** | **Keterangan** |
|----------|---------|----------------|
| **Admin** | `1` | Memiliki akses penuh untuk CRUD aset, status, kategori, pengguna, dan perawatan. |
| **Technician** | `2` | Hanya bisa melihat dan memperbarui perawatan aset. |
| **User** | `3` | Hanya bisa meminjam aset dan mengajukan permintaan perawatan. |

✅ **Fitur Baru** *(Dalam Pengujian 🛠️)*  
- **Peminjaman Aset**  
  - Pengguna dapat membuat permintaan peminjaman aset.  
  - Admin dapat menyetujui atau menolak permintaan.  
- **Permintaan Pemeliharaan**  
  - Pengguna dapat mengajukan permintaan pemeliharaan aset.  
  - Admin dapat menyetujui dan menentukan jadwal pemeliharaan.  
- **Dokumentasi API di Postman**  
  🔗 **[Lihat Dokumentasi API](https://universal-desert-823258-1.postman.co/workspace/Asset-Management~d0881856-648f-4073-99bc-54a043912a33/collection/26349837-64b9586a-fdb5-4ec5-8a5b-900f407813dc?action=share&creator=26349837)**  

---

## **Stack Teknologi**
- **Go (Golang)** → Bahasa pemrograman backend.  
- **Gin-Gonic** → Web framework untuk Golang.  
- **GORM** → ORM untuk manajemen database.  
- **MySQL** → Database utama.  
- **JWT** → JSON Web Token untuk autentikasi.  

---

## **Cara Menjalankan Project**
```sh
# Clone repository
git clone https://github.com/williamy101/asset-management.git
cd asset-management

# Jalankan aplikasi
go run main.go
```

