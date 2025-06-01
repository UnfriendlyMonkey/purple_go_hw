package product

import (
	"go/hw/4-order-api/pkg/req"
	"go/hw/4-order-api/pkg/resp"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("POST /product", handler.CreateProduct())
	router.HandleFunc("GET /product/{id}", handler.GetProductById())
	router.HandleFunc("PUT /product/{id}", handler.UpdateProduct())
	router.HandleFunc("DELETE /product/{id}", handler.DeleteProduct())
}

func (h *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}

		product := &Product{
			Name:        body.Name,
			Description: body.Description,
			Price:       body.Price,
			Quantity:    body.Quantity,
			Image:       body.Image,
		}

		createdProduct, err := h.ProductRepository.CreateProduct(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp.Json(w, createdProduct, http.StatusCreated)
	}
}

func (h *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := h.ProductRepository.GetProductById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp.Json(w, product, http.StatusOK)
	}
}

func (h *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}

		product := &Product{
			Model: gorm.Model{
				ID: uint(id),
			},
			Name:        body.Name,
			Description: body.Description,
			Price:       body.Price,
			Quantity:    body.Quantity,
			Image:       body.Image,
		}

		updatedProduct, err := h.ProductRepository.UpdateProduct(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp.Json(w, updatedProduct, http.StatusOK)
	}
}

func (h *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.ProductRepository.GetProductById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = h.ProductRepository.DeleteProduct(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Json(w, nil, http.StatusNoContent)
	}
}
