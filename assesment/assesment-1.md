# Assessment: Pembuatan dan Analisis Chaos Engineering Scenario untuk microchaos (order-api & notifier)

## 1. Ringkasan

Project **microchaos** terdiri dari dua service utama:
- **order-api**: bergantung pada **MySQL** dan **RabbitMQ**
- **notifier**: bergantung pada **MongoDB** dan **RabbitMQ**

Saat ini, skenario chaos engineering yang tersedia hanya untuk dependency **MySQL** (lihat: `chaostoolkit/experiments/02-microchaos/02.02-add-mysql-toxic-latency.json`). Belum ada skenario untuk **RabbitMQ** maupun **MongoDB**.

---

## 2. Assessment Kebutuhan Scenario Chaos Engineering

### Chaos Engineering Scenario: RabbitMQ & MongoDB (order-api & notifier)

Pada sistem **microchaos**, dependency utama yang perlu diuji ketahanannya adalah **RabbitMQ** (digunakan oleh `order-api` & `notifier`) dan **MongoDB** (digunakan oleh `notifier`). Gangguan pada kedua dependency ini dapat menyebabkan:
- Notifikasi tidak diproses atau hilang (RabbitMQ & MongoDB)
- Data menjadi tidak konsisten atau service error (MongoDB)
- Potensi kehilangan pesan, keterlambatan, atau error pada service

> **Instruksi untuk peserta:**  
> Buatlah skenario chaos engineering menggunakan Toxiproxy pada dependency **RabbitMQ** (`rabbitmq_proxy`) dan **MongoDB** (`mongo_proxy`). Untuk setiap dependency, lakukan hal berikut:
> - Suntikkan gangguan jaringan (_toxic_) sesuai jenis yang dipilih (misal: latency, bandwidth limit, timeout, disconnect, packet loss, slicer).
> - Lakukan aksi/probe pada service terkait (`order-api` dan/atau `notifier`) untuk mengamati dampak gangguan.
> - Pastikan ada aksi rollback untuk menghapus toxic setelah eksperimen selesai.
> - Dokumentasikan setiap skenario dalam file eksperimen journal ChaosToolkit (format JSON).
> - Setelah menjalankan eksperimen dan mendapatkan file journal (misal: `hasil.json`), Anda dapat menghasilkan report otomatis dalam format PDF menggunakan perintah:
>   ```
>   chaos report --export-format=pdf hasil.json hasil.pdf
>   ```
>   Gantilah `hasil.json` dan `hasil.pdf` sesuai nama file yang diinginkan. Report PDF ini berguna untuk dokumentasi, audit, atau presentasi hasil eksperimen chaos engineering.

**Contoh jenis gangguan (_toxic_) yang dapat diuji pada RabbitMQ & MongoDB:**
- **Latency**: Menambah delay komunikasi ke dependency.
- **Bandwidth limit**: Membatasi kecepatan transfer data.
- **Timeout**: Memaksa operasi menjadi timeout.
- **Disconnect**: Memutus koneksi secara tiba-tiba **(disable proxy)**.
- **Packet loss/slicer**: Memecah atau membuang sebagian paket data.

> **Analisis yang perlu dilakukan:**
> - Amati bagaimana service (`order-api` dan `notifier`) merespons ketika terjadi gangguan pada RabbitMQ atau MongoDB.
> - Evaluasi apakah sistem sudah resilien atau masih perlu perbaikan pada handling error, retry, atau observabilitas.
> - Lakukan probe pada endpoint health atau endpoint notifikasi untuk mengamati dampak gangguan.

> **Catatan:**  
> Untuk setiap skenario, pastikan rollback toxic dilakukan agar sistem kembali ke kondisi normal setelah eksperimen.

---


## 3. Petunjuk Pengumpulan

Setelah Anda membuat file skenario eksperimen (format JSON) dan dokumen analisis (format Markdown), lakukan langkah berikut untuk pengumpulan:

1. **Push ke Repository Pribadi**
   - Buat repository Git pribadi (misal di GitHub, GitLab, dsb).
   - Struktur direktori yang disarankan:
     ```
     /chaos-scenarios/
       ├── rabbitmq-latency.json
       ├── mongodb-bandwidth.json
       ├── ...
       └── analisis.md
     ```
   - Push seluruh file skenario (`*.json`) dan dokumen analisis (`analisis.md`) ke repository tersebut.

2. **Share Link Repository**
   - Setelah selesai, bagikan link repository Anda kepada instruktor..

3. **Format Dokumen Analisis**
   - Dokumen analisis (`analisis.md`) minimal memuat:
     - Penjelasan singkat setiap skenario (dependency, jenis toxic, tujuan).
     - Hasil pengamatan/probe (misal: screenshot, log, atau ringkasan hasil).
     - Analisis dampak dan rekomendasi perbaikan resiliency (jika ada).
     - (Opsional) Lampirkan report PDF/HTML jika diizinkan.

4. **Catatan**
   - Pastikan file JSON skenario dapat dijalankan di lingkungan ChaosToolkit + Toxiproxy.
   - Dokumentasi dan analisis dalam format Markdown agar mudah dibaca dan direview.

Dengan mengikuti langkah di atas, Anda dapat mengumpulkan hasil eksperimen chaos engineering beserta analisisnya.

