package params

type Register struct {
	FullName string `json:"full_name" validate:"required"`
	NIM      string `json:"nim" validate:"required"`
	Jurusan  string `json:"jurusan" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	NIM      string `json:"nim" validate:"required"`
	Password string `json:"password" validate:"required"`
}
