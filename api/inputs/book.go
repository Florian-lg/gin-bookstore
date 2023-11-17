package inputs

type Book struct {
	Title  string  `json:"title" binding:"required"`
	Price  float64 `json:"price"`
	Author string  `json:"author"`
}
