package handlers

import (
	"InvoiceLaundryGolang/internal"
	"InvoiceLaundryGolang/models"
	"encoding/json"
	"net/http"
	"time"
)

// GetServices mengembalikan daftar layanan yang tersedia
func GetServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	services := models.GetAllServices()
	json.NewEncoder(w).Encode(services)
}

// GetStatusPembayaran mengembalikan daftar status pembayaran
func GetStatusPembayaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	statuses := models.GetAllStatusPembayaran()
	json.NewEncoder(w).Encode(statuses)
}

// CreateInvoice membuat invoice baru
func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.InvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if req.Nama == "" || req.NomorTelepon == "" || req.Berat <= 0 {
		http.Error(w, "Data tidak lengkap", http.StatusBadRequest)
		return
	}

	// Buat invoice
	layanan := models.GetService(req.Layanan)
	if layanan.Name == "Unknown" {
		http.Error(w, "Layanan tidak valid", http.StatusBadRequest)
		return
	}

	invoice := models.Invoice{
		Nama:             req.Nama,
		NomorTelepon:     req.NomorTelepon,
		Berat:            req.Berat,
		Layanan:          layanan,
		StatusPembayaran: req.StatusPembayaran,
	}

	// Set tanggal dan hitung total
	invoice.SetTanggal(time.Now())
	invoice.HitungTotal()

	// Buat response
	response := models.InvoiceResponse{
		Nama:             invoice.Nama,
		NomorTelepon:     invoice.NomorTelepon,
		Berat:            invoice.Berat,
		Layanan:          invoice.Layanan,
		HargaTotal:       invoice.HargaTotal,
		HargaTotalFormat: internal.FormatRupiah(invoice.HargaTotal),
		TanggalPesanan:   invoice.TanggalPesanan.Format("02-01-2006"),
		PerkiraanSelesai: invoice.PerkiraanSelesai.Format("02-01-2006"),
		StatusPembayaran: invoice.StatusPembayaran,
	}

	json.NewEncoder(w).Encode(response)
}
