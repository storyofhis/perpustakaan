package params

type CreateBook struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Place    string `json:"place"`
}

type UpdateBook struct {
	Quantity int `json:"quantity"`
}
