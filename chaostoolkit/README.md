# 📦 Deploy Chaos Toolkit + Toxiproxy dengan Docker Compose



## 1) Struktur & Lokasi File
Pastikan file `docker-compose.yml` berada di:
- `<ROOT_PROJECT>/chaostoolkit/docker-compose.yml`

## 2) Isi `docker-compose.yml`
```yaml
version: "3.9"

services:
  toxiproxy:
    image: ghcr.io/shopify/toxiproxy:latest
    container_name: toxiproxy
    ports:
      - "8474:8474"   # REST API Toxiproxy (management)
      - "6060-6070:6060-6070" # Port range for proxies
      # Port proxy akan terbuka sesuai yang kita buat (contoh: 8666)

  chaostoolkit:
    build:
      context: ./chaostoolkit
    container_name: chaostoolkit
    working_dir: /experiments
    volumes:
      - ./experiments:/experiments
    # Biarkan container hidup agar bisa eksekusi perintah interaktif
    command: [ "bash", "-lc", "echo 'Chaos container ready'; tail -f /dev/null" ]
```

> **Catatan**  
> - Folder `./chaostoolkit` berisi `Dockerfile` untuk image Chaos Toolkit (install `chaostoolkit` dan plugin yang diperlukan).  
> - Folder `./experiments` untuk menyimpan file eksperimen chaos engineering.

## 3) Jalankan Docker Compose (build & start)
Dari direktori root proyek, jalankan:
```bash
docker-compose -f <ROOT_PROJECT>/chaostoolkit/docker-compose.yml up -d --build
```

## 4) Validasi Toxiproxy
Cek endpoint management Toxiproxy:
```bash
curl -s localhost:8474/proxies
```
**Ekspektasi:** mengembalikan `{}` (kosong) jika belum ada proxy yang dibuat.

Contoh output:
```json
{}
```

## 5) Validasi Chaos Toolkit
Pastikan container `chaostoolkit` aktif lalu cek CLI:
```bash
docker exec -it chaostoolkit chaos
```
**Ekspektasi:** menampilkan bantuan/usage dari perintah `chaos` (CLI Chaos Toolkit).

