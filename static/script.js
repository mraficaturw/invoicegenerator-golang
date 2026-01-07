// API Base URL
const API_BASE = '';

// Store current invoice data for WhatsApp
let currentInvoice = null;

// Store install prompt for later use
let deferredPrompt = null;

// Register Service Worker
if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('/service-worker.js')
            .then((registration) => {
                console.log('✅ Service Worker registered:', registration.scope);

                // Check for updates
                registration.addEventListener('updatefound', () => {
                    const newWorker = registration.installing;
                    newWorker.addEventListener('statechange', () => {
                        if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
                            // New content available, show update notification
                            if (confirm('Versi baru tersedia! Muat ulang untuk update?')) {
                                window.location.reload();
                            }
                        }
                    });
                });
            })
            .catch((error) => {
                console.error('❌ Service Worker registration failed:', error);
            });
    });
}

// Handle PWA install prompt
window.addEventListener('beforeinstallprompt', (e) => {
    console.log('📱 PWA install prompt available');
    e.preventDefault();
    deferredPrompt = e;

    // Show install button if it exists
    const installBtn = document.getElementById('installBtn');
    if (installBtn) {
        installBtn.style.display = 'flex';
    }
});

// Install PWA function
function installPWA() {
    if (deferredPrompt) {
        deferredPrompt.prompt();
        deferredPrompt.userChoice.then((choiceResult) => {
            if (choiceResult.outcome === 'accepted') {
                console.log('✅ User accepted PWA install');
            } else {
                console.log('❌ User dismissed PWA install');
            }
            deferredPrompt = null;

            // Hide install button
            const installBtn = document.getElementById('installBtn');
            if (installBtn) {
                installBtn.style.display = 'none';
            }
        });
    }
}

// Handle successful PWA install
window.addEventListener('appinstalled', () => {
    console.log('🎉 PWA installed successfully!');
    deferredPrompt = null;
});

// Load dropdown options on page load
document.addEventListener('DOMContentLoaded', async () => {
    await loadServices();
    await loadStatusPembayaran();
});

// Load services for dropdown
async function loadServices() {
    try {
        const response = await fetch(`${API_BASE}/api/services`);
        const services = await response.json();

        const select = document.getElementById('layanan');
        services.forEach(service => {
            const option = document.createElement('option');
            option.value = service.name;
            option.textContent = `${service.name} (Rp ${formatNumber(service.price)}/kg)`;
            select.appendChild(option);
        });
    } catch (error) {
        console.error('Error loading services:', error);
    }
}

// Load status pembayaran for dropdown
async function loadStatusPembayaran() {
    try {
        const response = await fetch(`${API_BASE}/api/status-pembayaran`);
        const statuses = await response.json();

        const select = document.getElementById('statusPembayaran');
        statuses.forEach(status => {
            const option = document.createElement('option');
            option.value = status;
            option.textContent = status;
            select.appendChild(option);
        });
    } catch (error) {
        console.error('Error loading statuses:', error);
    }
}

// Format number with thousand separators
function formatNumber(num) {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

// Handle form submission
document.getElementById('invoiceForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = {
        nama: document.getElementById('nama').value,
        nomor_telepon: document.getElementById('nomorTelepon').value,
        berat: parseFloat(document.getElementById('berat').value),
        layanan: document.getElementById('layanan').value,
        status_pembayaran: document.getElementById('statusPembayaran').value
    };

    try {
        const response = await fetch(`${API_BASE}/api/invoice`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (!response.ok) {
            const error = await response.text();
            alert('Error: ' + error);
            return;
        }

        const invoice = await response.json();
        currentInvoice = invoice; // Store invoice data
        displayInvoice(invoice);
    } catch (error) {
        console.error('Error creating invoice:', error);
        alert('Gagal membuat invoice. Silakan coba lagi.');
    }
});

// Display invoice in the preview section
function displayInvoice(invoice) {
    document.getElementById('invTanggal').textContent = invoice.tanggal_pesanan;
    document.getElementById('invNama').textContent = invoice.nama;
    document.getElementById('invHP').textContent = invoice.nomor_telepon;
    document.getElementById('invBerat').textContent = invoice.berat.toFixed(1);
    document.getElementById('invLayanan').textContent = invoice.layanan.name;
    document.getElementById('invTotal').textContent = invoice.harga_total_format;
    document.getElementById('invSelesai').textContent = invoice.perkiraan_selesai;

    const statusSpan = document.getElementById('invStatus');
    statusSpan.textContent = invoice.status_pembayaran;
    statusSpan.className = invoice.status_pembayaran === 'Lunas' ? 'lunas' : 'belum-lunas';

    // Show invoice section, hide form section
    document.getElementById('formSection').style.display = 'none';
    document.getElementById('invoiceSection').style.display = 'block';
}

// Reset form and show form section
function resetForm() {
    document.getElementById('invoiceForm').reset();
    document.getElementById('formSection').style.display = 'block';
    document.getElementById('invoiceSection').style.display = 'none';
    currentInvoice = null;
}

// Format phone number for WhatsApp (remove leading 0, add 62)
function formatPhoneForWhatsApp(phone) {
    // Remove any non-digit characters
    let cleaned = phone.replace(/\D/g, '');

    // If starts with 0, replace with 62
    if (cleaned.startsWith('0')) {
        cleaned = '62' + cleaned.substring(1);
    }
    // If doesn't start with 62, add it
    else if (!cleaned.startsWith('62')) {
        cleaned = '62' + cleaned;
    }

    return cleaned;
}

// Send invoice to WhatsApp
function sendToWhatsApp() {
    if (!currentInvoice) {
        alert('Tidak ada data invoice!');
        return;
    }

    const phone = formatPhoneForWhatsApp(currentInvoice.nomor_telepon);

    // Create invoice message
    const message = `==============================
        *ENJA LAUNDRY*
Jl. Pondok Bambu Asri Blok A4/B4 No 57
WA: +6281384608887
==============================

📅 *Tanggal Pesanan:* ${currentInvoice.tanggal_pesanan}

👤 *Nama:* ${currentInvoice.nama}
📱 *No. HP:* ${currentInvoice.nomor_telepon}
⚖️ *Berat:* ${currentInvoice.berat.toFixed(1)} kg
🧺 *Layanan:* ${currentInvoice.layanan.name}
💰 *Harga Total:* ${currentInvoice.harga_total_format}

📆 *Perkiraan Selesai:* ${currentInvoice.perkiraan_selesai}

------------------------------
💳 *Status:* ${currentInvoice.status_pembayaran}
------------------------------

Terima kasih sudah menggunakan *EnjaLaundry!* 🙏
==============================`;

    // Encode the message for URL
    const encodedMessage = encodeURIComponent(message);

    // Create WhatsApp URL
    const whatsappUrl = `https://wa.me/${phone}?text=${encodedMessage}`;

    // Open WhatsApp in new tab
    window.open(whatsappUrl, '_blank');
}
