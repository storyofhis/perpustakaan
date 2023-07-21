package views

import "time"

// pinjam buku
type CreateTransaction struct {
	Id                  uint      `json:"id"`
	NamaMahasiswa       string    `json:"full_name"`
	JudulBuku           string    `json:"judul_buku"`
	TanggalPeminjaman   time.Time `json:"tanggal_peminjaman"`
	TanggalPengembalian time.Time `json:"tanggal_pengembalian"`
	BatasPengembalian   time.Time `json:"batas_pengembalian"`
}

type GetTransaction struct {
}
