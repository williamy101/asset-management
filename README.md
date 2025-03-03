
# **Asset Management API**

## **Preview ERD**  
**![image](https://github.com/user-attachments/assets/2bda7a98-e782-4cbd-84b5-fa80f15dce49)
🔗 **[Lihat ERD lebih jelas](https://dbdiagram.io/d/ERD-Asset-Management-676dac255406798ef7b59c4e)**  

## **Slide Presentasi**  
![Slide Presentasi](https://github.com/user-attachments/assets/876e6272-ddb8-4e9a-b5dd-ab4a1c88d037)  
🔗 **[Lihat Presentasi di Canva](https://www.canva.com/design/DAGcmM3nHpE/RlvXGsFwEORdt432r6Bd3A/edit?utm_content=DAGcmM3nHpE&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)**  

## **Deskripsi**  
Aplikasi backend untuk mengelola aset, mencakup operasi CRUD, kontrol akses berbasis peran, dan manajemen pengguna.  

---

## **📌 Fitur Utama**
✅ **Manajemen Aset**  
- CRUD untuk aset, kategori, dan status.  
- Penyimpanan informasi pemeliharaan aset.  

✅ **Autentikasi & Kontrol Akses**  
- Pendaftaran pengguna dengan peran (`admin`, `user`).  
- Autentikasi menggunakan **JWT**.  
- **RBAC (Role-Based Access Control)** untuk membatasi akses berdasarkan peran.  

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

## **🛠️ Stack Teknologi**
- **Go (Golang)** → Bahasa pemrograman backend.  
- **Gin-Gonic** → Web framework untuk Golang.  
- **GORM** → ORM untuk manajemen database.  
- **MySQL** → Database utama.  
- **JWT** → JSON Web Token untuk autentikasi.  

---

## **🚀 Cara Menjalankan Project**
```sh
# Clone repository
git clone https://github.com/williamy101/asset-management.git
cd asset-management

# Jalankan aplikasi
go run main.go
```
