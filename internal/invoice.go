package internal

import (
	"fmt"
	"strings"
)

// FormatRupiah memformat angka menjadi format rupiah
func FormatRupiah(amount float64) string {
	// Format dengan pemisah ribuan
	intAmount := int64(amount)
	str := fmt.Sprintf("%d", intAmount)

	// Tambahkan pemisah ribuan
	var result strings.Builder
	length := len(str)
	for i, digit := range str {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(digit)
	}

	return "Rp " + result.String()
}
