package public

type paginationQueryStruct struct {
	Page    int `json:"page" query:"page" example:"1"`
	PerPage int `json:"per_page" query:"per_page" example:"10"`
}

type paginationStruct struct {
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"0"`
}

type fetchBookByISBNParam struct {
	ISBN string `json:"isbn" param:"isbn" validate:"required"`
}

type createBookRequest struct {
	ISBN        string  `json:"isbn" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"author" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0,decimals2"`
	Description string  `json:"description" validate:"required"`
	Genre       string  `json:"genre" validate:"required"`
	Quantity    *int    `json:"quantity" validate:"required,gte=0"`
}

type updateBookRequest struct {
	ISBN        string  `json:"isbn" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"author" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0,decimals2"`
	Description string  `json:"description" validate:"required"`
	Genre       string  `json:"genre" validate:"required"`
	Quantity    *int    `json:"quantity" validate:"required,gte=0"`
}
