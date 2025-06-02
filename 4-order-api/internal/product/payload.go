package product

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"min=1"`
	Quantity    int    `json:"quantity" validate:"min=1"`
	Image       string `json:"image"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"min=1"`
	Quantity    int    `json:"quantity" validate:"min=1"`
	Image       string `json:"image"`
}
