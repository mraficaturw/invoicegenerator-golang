package main

import (
	"InvoiceLaundryGolang/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files (frontend)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API Routes
	http.HandleFunc("/api/services", handlers.GetServices)
	http.HandleFunc("/api/status-pembayaran", handlers.GetStatusPembayaran)
	http.HandleFunc("/api/invoice", handlers.CreateInvoice)

	port := ":8080"
	fmt.Printf("🧺 Enja Laundry Invoice API running on http://localhost%s\n", port)
	fmt.Println("📄 Open http://localhost:8080 in your browser")

	log.Fatal(http.ListenAndServe(port, nil))
}
