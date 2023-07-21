package params

type CreateTransaction struct {
	NamaMahasiswa string `json:"full_name" validate:"required"`
	JudulBuku     string `json:"judul_buku" validate:"required"`
}
