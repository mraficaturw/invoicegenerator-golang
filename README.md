# 🧺 Enja Laundry - Invoice Generator PWA

Aplikasi pembuatan invoice laundry otomatis dengan fitur kirim ke WhatsApp.

## ✨ Fitur

- 📝 Buat invoice laundry otomatis
- 💬 Kirim invoice langsung ke WhatsApp pelanggan
- 📱 Progressive Web App (PWA) - dapat diinstall di HP
- 🔄 Bekerja offline (setelah pertama kali dibuka)
- 💨 Cepat dan ringan

## 🚀 Cara Menjalankan Lokal

```bash
# Build aplikasi
go build -o laundry-api.exe .

# Jalankan server
./laundry-api.exe
```

Buka `http://localhost:8080` di browser.

---

## 📦 Cara Deploy ke Publik

### Opsi 1: Deploy ke Railway (Gratis & Mudah)

1. **Buat akun di [Railway.app](https://railway.app)**

2. **Push kode ke GitHub**

   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/username/enja-laundry.git
   git push -u origin main
   ```

3. **Deploy di Railway:**
   - Klik "New Project" → "Deploy from GitHub repo"
   - Pilih repository Anda
   - Railway akan otomatis mendeteksi Go dan deploy

4. **Setting Domain:**
   - Klik pada service → Settings → Domains
   - Generate domain atau tambahkan custom domain

### Opsi 2: Deploy ke Render.com (Gratis)

1. **Buat akun di [Render.com](https://render.com)**

2. **Buat file `render.yaml` di root project:**

   ```yaml
   services:
     - type: web
       name: enja-laundry
       env: go
       buildCommand: go build -o app .
       startCommand: ./app
       envVars:
         - key: PORT
           value: 8080
   ```

3. **Connect GitHub dan deploy**

### Opsi 3: Deploy ke Fly.io (Gratis tier tersedia)

1. **Install Fly CLI:**

   ```bash
   # Windows (PowerShell)
   iwr https://fly.io/install.ps1 -useb | iex
   ```

2. **Login dan deploy:**

   ```bash
   fly auth login
   fly launch
   fly deploy
   ```

### Opsi 4: Deploy ke VPS (Contoh: DigitalOcean, Vultr)

1. **SSH ke server:**

   ```bash
   ssh root@your-server-ip
   ```

2. **Install Go:**

   ```bash
   wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   ```

3. **Clone dan build:**

   ```bash
   git clone https://github.com/username/enja-laundry.git
   cd enja-laundry
   go build -o laundry-api .
   ```

4. **Jalankan dengan systemd:**

   ```bash
   sudo nano /etc/systemd/system/laundry.service
   ```

   ```ini
   [Unit]
   Description=Enja Laundry API
   After=network.target

   [Service]
   Type=simple
   User=root
   WorkingDirectory=/root/enja-laundry
   ExecStart=/root/enja-laundry/laundry-api
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

   ```bash
   sudo systemctl enable laundry
   sudo systemctl start laundry
   ```

5. **Setup Nginx + SSL:**

   ```bash
   sudo apt install nginx certbot python3-certbot-nginx
   
   sudo nano /etc/nginx/sites-available/laundry
   ```

   ```nginx
   server {
       server_name yourdomain.com;
       
       location / {
           proxy_pass http://localhost:8080;
           proxy_http_version 1.1;
           proxy_set_header Upgrade $http_upgrade;
           proxy_set_header Connection 'upgrade';
           proxy_set_header Host $host;
           proxy_cache_bypass $http_upgrade;
       }
   }
   ```

   ```bash
   sudo ln -s /etc/nginx/sites-available/laundry /etc/nginx/sites-enabled/
   sudo nginx -t
   sudo systemctl restart nginx
   sudo certbot --nginx -d yourdomain.com
   ```

---

## 📱 Cara Install PWA

### Di Android

1. Buka website di Chrome
2. Tap tombol "Install Aplikasi" di header, ATAU
3. Tap menu (3 titik) → "Add to Home Screen"

### Di iOS

1. Buka website di Safari
2. Tap icon Share (kotak dengan panah)
3. Scroll dan tap "Add to Home Screen"

### Di Desktop

1. Buka website di Chrome/Edge
2. Klik icon install di address bar, ATAU
3. Klik tombol "Install Aplikasi"

---

## 📁 Struktur Project

```text
InvoiceLaundryGolang/
├── main.go                    # API Server
├── go.mod                     # Go modules
├── handlers/
│   └── invoice_handler.go     # API Handlers
├── internal/
│   └── invoice.go             # Helper functions
├── models/
│   └── structure.go           # Data models
└── static/
    ├── index.html             # Frontend
    ├── style.css              # Styling
    ├── script.js              # JavaScript
    ├── manifest.json          # PWA Manifest
    ├── service-worker.js      # Service Worker
    └── icons/                 # PWA Icons
```

## 🔧 API Endpoints

| Method | Endpoint                 | Description                 |
| ------ | ------------------------ | --------------------------- |
| GET    | `/api/services`          | Daftar layanan tersedia     |
| GET    | `/api/status-pembayaran` | Daftar status pembayaran    |
| POST   | `/api/invoice`           | Buat invoice baru           |

---

## 📝 Catatan Penting untuk PWA

1. **HTTPS Wajib**: PWA membutuhkan HTTPS untuk Service Worker bekerja
2. **Icons**: Pastikan semua ukuran icon tersedia di folder `/static/icons/`
3. **Manifest**: File `manifest.json` harus valid dan dapat diakses

## 🎨 Membuat Icons PWA

Gunakan tool online untuk generate icons dari gambar 512x512:

- [PWA Asset Generator](https://progressier.com/pwa-icons-and-ios-splash-screen-generator)
- [RealFaviconGenerator](https://realfavicongenerator.net/)

---

Made with ❤️ by Enja Laundry
