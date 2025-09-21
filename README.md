# 📘 Special Academy: Chaos Engineering  

## Pendahuluan  
**Special Academy** merupakan program internal di SPE Solution yang dirancang untuk menjadi wadah *sharing session* dan *knowledge transfer* seputar topik **Chaos Engineering**.  

Chaos Engineering sendiri adalah metode untuk menguji seberapa tangguh (*resilient*) sebuah sistem dengan cara **menciptakan gangguan secara terkontrol** (contohnya: service down, latency, resource exhaustion, atau dependency failure). Lewat eksperimen ini, kita bisa:  
- Menemukan titik lemah yang tersembunyi.  
- Meningkatkan *observability* dan monitoring.  
- Memastikan sistem tetap berjalan walau ada sebagian komponen yang gagal.  
- Mendorong budaya engineering yang fokus pada *reliability*.  

## Tujuan Special Academy  
- Memberikan pemahaman dasar tentang Chaos Engineering.  
- Berbagi pengalaman nyata implementasi di layanan internal maupun eksternal.  
- Menjadi forum diskusi untuk membahas *best practice* dan tantangan implementasi.  
- Membantu engineer mengasah kemampuan membangun sistem yang tahan banting (*resilient systems*).  

## Format Kegiatan  
- **Sesi Sharing**: Penjelasan konsep dan studi kasus.  
- **Hands-on**: Praktik langsung melakukan chaos experiment di lingkungan yang aman.  
- **Diskusi**: Mengulas hasil eksperimen dan langkah perbaikan sistem.  
- **QnA**: Sesi tanya jawab untuk memperdalam pemahaman peserta.  

## Outcome yang Diharapkan  
- Peserta memahami konsep dasar Chaos Engineering.  
- Peserta bisa merancang dan menjalankan eksperimen chaos dengan aman.  
- Terbentuk *shared knowledge* di tim terkait *resilience engineering*.  
- Meningkatkan *system reliability* di berbagai project yang sedang berjalan.

## 🛠️ Prasyarat (Pre-requisite) Chaos Engineering Hands-on

Sebelum mengikuti sesi hands-on Chaos Engineering, pastikan Anda sudah menyiapkan:

- **Docker**  
  Instalasi Docker diperlukan untuk menjalankan container environment secara terisolasi.  
  [Panduan instalasi Docker](https://docs.docker.com/get-docker/)

- **Docker Compose**  
  Digunakan untuk mengelola multi-container environment (misal: Chaos Toolkit + Toxiproxy) dengan satu perintah.  
  [Panduan instalasi Docker Compose](https://docs.docker.com/compose/install/)

- **Clone Repository Ini**  
  Sebelum memulai, pastikan Anda sudah meng-*clone* repository ini ke komputer Anda.  
  Jalankan perintah berikut di terminal:
  ```bash
  git clone https://gitlab.spesolution.net/data/rnd/special-academy-chaos-engineering.git
  cd special-academy-chaos-engineering
  ```



> **Catatan:**  
> - Pastikan perintah `docker` dan `docker-compose` dapat dijalankan di terminal Anda.
> - Tidak perlu install Python/Chaos Toolkit secara manual, semua sudah disediakan di dalam container.


## 📦 Referensi: Setup Environment Chaos Toolkit

Untuk panduan lengkap setup environment Chaos Toolkit + Toxiproxy menggunakan Docker Compose, silakan lihat dokumentasi berikut:

[📄 Setup Chaos Toolkit + Toxiproxy (docs/setup-chaostoolkit-toxiproxy.md)](./chaostoolkit/README.md)

Dokumentasi tersebut mencakup:
- Struktur dan lokasi file yang diperlukan
- Contoh `docker-compose.yml`
- Cara build & menjalankan container
- Validasi service Toxiproxy & Chaos Toolkit

Pastikan mengikuti langkah-langkah pada dokumen tersebut sebelum memulai eksperimen chaos engineering di sesi hands-on.

## 📦 Referensi: Setup Project Microchaos

Untuk panduan setup dan struktur project eksperimen microchaos (Toxiproxy di antara ChaosToolkit & microservices), silakan lihat:

[📄 Setup Project Microchaos (microchaos/README.md)](./microchaos/README.md)

Dokumentasi tersebut mencakup:
- Penjelasan arsitektur integrasi Toxiproxy dengan berbagai microservices (MySQL, RabbitMQ, MongoDB, dsb)
- Daftar file eksperimen chaos beserta deskripsi singkat fungsinya
- Langkah-langkah menjalankan eksperimen toxic (misal: latency, bandwidth, timeout, disconnect, dll)
- Panduan melakukan rollback dan validasi hasil eksperimen untuk memastikan sistem kembali normal

Pastikan membaca dokumen tersebut untuk memahami alur eksperimen dan menyesuaikan skenario sesuai kebutuhan project Anda.

