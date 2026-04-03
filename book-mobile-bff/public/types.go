package public

type paginationQueryStruct struct {
	Page    int `json:"page" query:"page" example:"1"`
	PerPage int `json:"per_page" query:"per_page" example:"10"`
}

type paginationStruct struct {
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"0"`
}

type createUserRequest struct {
	UserID   string `json:"userId" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Address2 string `json:"address2"`
	City     string `json:"city" validate:"required"`
	State    string `json:"state" validate:"required,len=2"`
	Zipcode  string `json:"zipcode" validate:"required"`
}

type fetchUserByIdParam struct {
	ID uint `json:"id" param:"id" validate:"required"`
}

type fetchUserByUserIdQuery struct {
	UserID string `json:"userId" query:"userId" validate:"required,email"`
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

type getUserResponse struct {
	ID     uint   `json:"id" `
	UserID string `json:"userId" validate:"required,email"`
	Name   string `json:"name" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

type userResponse struct {
	ID       uint   `json:"id" `
	UserID   string `json:"userId" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Address2 string `json:"address2"`
	City     string `json:"city" validate:"required"`
	State    string `json:"state" validate:"required,len=2"`
	Zipcode  string `json:"zipcode" validate:"required"`
}

type bookResponse struct {
	ID          uint    `json:"-"`
	ISBN        string  `json:"ISBN" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"Author" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0,decimals2"`
	Description string  `json:"description" validate:"required"`
	Genre       string  `json:"genre" validate:"required"`
	Quantity    *int    `json:"quantity" validate:"required,gte=0"`
	Summary     *string `json:"summary,omitempty"`
}

type getbookResponse struct {
	ID          uint    `json:"-"`
	ISBN        string  `json:"ISBN" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"Author" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0,decimals2"`
	Description string  `json:"description" validate:"required"`
	Genre       int     `json:"genre" validate:"required"`
	Quantity    *int    `json:"quantity" validate:"required,gte=0"`
	Summary     *string `json:"summary,omitempty"`
}
