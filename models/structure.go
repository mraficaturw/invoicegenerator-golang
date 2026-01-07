package models

import "time"

// Service menyimpan informasi layanan dan harga per kg
type Service struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// GetService mengembalikan informasi layanan berdasarkan nama layanan
func GetService(serviceName string) Service {
	switch serviceName {
	case "Setrika Only":
		return Service{Name: "Setrika Only", Price: 5000}
	case "Cuci Only":
		return Service{Name: "Cuci Only", Price: 5000}
	case "Cuci + Setrika":
		return Service{Name: "Cuci + Setrika", Price: 7000}
	default:
		return Service{Name: "Unknown", Price: 0}
	}
}

// GetAllServices mengembalikan semua layanan yang tersedia
func GetAllServices() []Service {
	return []Service{
		{Name: "Setrika Only", Price: 5000},
		{Name: "Cuci Only", Price: 5000},
		{Name: "Cuci + Setrika", Price: 7000},
	}
}

// GetAllStatusPembayaran mengembalikan semua status pembayaran
func GetAllStatusPembayaran() []string {
	return []string{"Belum Lunas", "Lunas"}
}

// InvoiceRequest untuk menerima data dari frontend
type InvoiceRequest struct {
	Nama             string  `json:"nama"`
	NomorTelepon     string  `json:"nomor_telepon"`
	Berat            float64 `json:"berat"`
	Layanan          string  `json:"layanan"`
	StatusPembayaran string  `json:"status_pembayaran"`
}

// InvoiceResponse untuk mengirim data invoice ke frontend
type InvoiceResponse struct {
	Nama             string  `json:"nama"`
	NomorTelepon     string  `json:"nomor_telepon"`
	Berat            float64 `json:"berat"`
	Layanan          Service `json:"layanan"`
	HargaTotal       float64 `json:"harga_total"`
	HargaTotalFormat string  `json:"harga_total_format"`
	TanggalPesanan   string  `json:"tanggal_pesanan"`
	PerkiraanSelesai string  `json:"perkiraan_selesai"`
	StatusPembayaran string  `json:"status_pembayaran"`
}

// Invoice menyimpan semua data invoice
type Invoice struct {
	Nama             string
	NomorTelepon     string
	Berat            float64
	Layanan          Service
	HargaTotal       float64
	TanggalPesanan   time.Time
	PerkiraanSelesai time.Time
	StatusPembayaran string
}

// HitungTotal menghitung harga total berdasarkan berat dan harga layanan
func (i *Invoice) HitungTotal() {
	i.HargaTotal = i.Berat * i.Layanan.Price
}

// SetTanggal mengatur tanggal pesanan dan perkiraan selesai
func (i *Invoice) SetTanggal(tanggalPesanan time.Time) {
	i.TanggalPesanan = tanggalPesanan
	i.PerkiraanSelesai = tanggalPesanan.AddDate(0, 0, 2) // +2 hari
}
