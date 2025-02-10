## Asset Management

### *Untuk backend code di branch master*

### 1. **Tabel asset_categories**
Tabel ini digunakan untuk menyimpan kategori aset. Setiap kategori seperti "Elektronik", "Furnitur", dan "Kendaraan" memiliki `category_id` yang unik. Selain itu, tabel ini juga mencatat kapan kategori tersebut dibuat dan terakhir diubah menggunakan `created_at` dan `updated_at`.

### 2. **Tabel statuses**
Tabel `statuses` berisi status yang dapat diberikan pada aset, pemeliharaan, dan peminjaman. Contoh status yang umum adalah "Available", "In Use", "In Maintenance", "Scheduled", "Borrowed", dan "Returned". Kolom `created_at` dan `updated_at` digunakan untuk mencatat kapan status dibuat dan terakhir diperbarui.

### 3. **Tabel roles**
Tabel `roles` menyimpan data terkait peran pengguna dalam sistem. Setiap peran seperti "Admin" atau "Technician" memiliki `role_id` yang digunakan untuk mengaitkan pengguna dengan peran mereka di sistem. Hal ini memungkinkan pengaturan akses berdasarkan role.

### 4. **Tabel assets**
Tabel `assets` menyimpan informasi tentang aset, termasuk nama aset, kategori (`category_id`), status (`status_id`), dan tanggal pemeliharaan berikutnya.
- **Kolom `user_id`** pada tabel ini menunjukkan siapa peminjam aset saat ini (jika ada).
- Jika aset sedang dalam peminjaman, maka statusnya akan berubah menjadi "Unavailable".
- Jika aset dikembalikan, maka statusnya akan kembali menjadi "Available".

### 5. **Tabel users**
Tabel `users` menyimpan data pengguna yang terdaftar dalam sistem, seperti nama, email, password (yang akan dienkripsi), dan `role_id` yang menunjukkan peran pengguna dalam sistem. Setiap pengguna memiliki peran yang menentukan hak akses mereka (misalnya, admin atau teknisi).

### 6. **Tabel maintenances**
Tabel `maintenances` digunakan untuk mencatat semua aktivitas pemeliharaan yang terjadi pada aset. Setiap pemeliharaan terkait dengan aset tertentu (`asset_id`), pekerja (`user_id`), serta status pemeliharaan (`status_id`).
- Status pemeliharaan bisa berupa "Scheduled", "In Progress", atau "Completed".
- Biaya pemeliharaan juga dapat dicatat dalam kolom `cost`.

### 7. **Tabel maintenance_requests**
Tabel `maintenance_requests` digunakan untuk mencatat permintaan pemeliharaan dari pengguna terhadap aset tertentu.
- Permintaan ini dapat diajukan oleh pengguna dan disetujui oleh admin.
- Status permintaan tercatat dalam `status_id` seperti "Pending Approval", "Approved", atau "Rejected".
- Jika permintaan disetujui, maka akan dijadwalkan dalam tabel `maintenances` dengan status "Scheduled".

### 8. **Tabel borrow_asset_requests**
Tabel `borrow_asset_requests` menyimpan permintaan peminjaman aset oleh pengguna.
- Pengguna dapat meminta peminjaman aset dengan menyertakan `requested_start_date` dan `requested_end_date`.
- Admin dapat menyetujui atau menolak permintaan ini, yang dicatat dalam `status_id`.
- Jika permintaan disetujui, aset akan dipindahkan ke tabel `borrowed_assets` dan statusnya diperbarui menjadi "Borrowed".

### 9. **Tabel borrowed_assets**
Tabel `borrowed_assets` menyimpan catatan aset yang sedang dipinjam oleh pengguna.
- `borrow_date` mencatat kapan aset dipinjam.
- `return_date` mencatat kapan aset dikembalikan.
- `status_id` menunjukkan status aset, seperti "Borrowed", "Returned", atau "Overdue".
- Jika aset sudah dikembalikan, maka status di `assets` diperbarui menjadi "Available".

### **Relasi dan Foreign Key**
- `maintenances` memiliki foreign key yang menghubungkan dengan `assets`, `users`, dan `statuses`. Ini memastikan bahwa setiap pemeliharaan terkait dengan aset yang benar, teknisi yang tepat, dan status yang sesuai.
- `maintenance_requests` dan `borrow_asset_requests` terhubung ke `statuses` untuk melacak progres persetujuan dan eksekusi.
- `borrowed_assets` menghubungkan aset yang sedang dipinjam dengan pengguna yang meminjamnya.
- Penggunaan **ON DELETE CASCADE** pada foreign key memastikan bahwa jika data yang terkait dihapus (misalnya, jika sebuah aset dihapus), data terkait di tabel lain juga akan dihapus secara otomatis untuk menjaga integritas data.

