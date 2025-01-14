## Asset Management
### *Untuk backend code di branch master*
### 1. **Tabel asset_categories**
Tabel ini digunakan untuk menyimpan kategori aset. Setiap kategori seperti "Perlengkapan IT", "Furnitur", dan "Kendaraan" memiliki category_id yang unik. Selain itu, tabel ini juga mencatat kapan kategori tersebut dibuat dan terakhir diubah menggunakan created_at dan updated_at.

### 2. **Tabel statuses**
Tabel statuses berisi status yang dapat diberikan pada aset dan pemeliharaan. Contoh status yang umum adalah "Available", "In Use", "In Maintenance", dan "Scheduled". Kolom created_at dan updated_at digunakan untuk mencatat kapan status dibuat dan terakhir kali diperbarui.

### 3. **Tabel roles**
Tabel roles menyimpan data terkait peran pengguna dalam sistem. Setiap role, seperti "Admin" atau "User", memiliki role_id yang digunakan untuk mengaitkan pengguna dengan peran mereka di sistem. Hal ini memungkinkan pengaturan akses berdasarkan role.

### 4. **Tabel assets**
Tabel assets menyimpan informasi tentang aset, termasuk nama aset, kategori yang dihubungkan ke tabel asset_categories, status yang dihubungkan ke tabel statuses, dan tanggal pemeliharaan berikutnya. Aset yang dimasukkan ke dalam pemeliharaan akan diperbarui statusnya sesuai dengan status yang berlaku, seperti "In Maintenance".

### 5. **Tabel users**
Tabel users menyimpan data pengguna yang terdaftar dalam sistem, seperti nama, email, password (yang akan dienkripsi), dan role_id yang menunjukkan peran pengguna dalam sistem. Setiap pengguna memiliki role yang dapat menentukan hak akses mereka (seperti admin atau teknisi).

### 6. **Tabel maintenances**
Tabel ini digunakan untuk mencatat semua aktivitas pemeliharaan yang terjadi pada aset. Setiap pemeliharaan terkait dengan aset tertentu dan teknisi yang bertugas, serta status pemeliharaan (misalnya, "In Progress" atau "Completed"). Tabel ini memiliki kolom asset_id, user_id, dan status_id yang merujuk ke tabel lain untuk memastikan integritas data.

### **Relasi dan Foreign Key**
- Tabel maintenances memiliki foreign key yang menghubungkan dengan tabel assets, users, dan statuses. Ini memastikan bahwa setiap pemeliharaan terkait dengan aset yang benar, teknisi yang tepat, dan status yang sesuai.
- Penggunaan ON DELETE CASCADE pada foreign key memastikan bahwa jika data yang terkait dihapus (misalnya, jika sebuah aset dihapus), data terkait di tabel maintenances juga akan dihapus secara otomatis.
