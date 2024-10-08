package handlers

import (
	"encoding/json"
	"gordan/internal/dto"
	"gordan/internal/entity"
	"gordan/internal/infra/database"
	entityPKG "gordan/pkg/entity"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary      Create Product
// @Description  Create Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      401  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products [post]
// @Security	ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// List Product godoc
// @Summary      List Product by ID
// @Description  List Product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request   path      string  true  "product id"
// @Success      201	{object}	entity.Product
// @Failure      400  {object}  Error
// @Failure      401  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary      Update Product by ID
// @Description  Update Product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request   path      string  true  "product id"
// @Success      200
// @Failure      400  {object}  Error
// @Failure      401  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [put]
// @Security	ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary      Delete Product by ID
// @Description  Delete Product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request   path      string  true  "product id"
// @Success      200
// @Failure      400  {object}  Error
// @Failure      401  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /products/{id} [delete]
// @Security	ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// List All Products godoc
// @Summary      List All Products
// @Description  List All Products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page    query      string  false  "page number"
// @Param        limit   query      string  false  "limit number"
// @Param        sort    query      string  false  "sort string"
// @Success      200	{array}	entity.Product
// @Failure      400 	{object}  Error
// @Failure      401  	{object}  Error
// @Failure      404 	{object}  Error
// @Failure      500 	{object}  Error
// @Router       /products [get]
// @Security	ApiKeyAuth
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	if sort == "" {
		sort = "desc"
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
