package product

import (
	"ecom_go/services/auth"
	"ecom_go/types"
	"ecom_go/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
	categoryStore types.ProductCategoryStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, categoryStore types.ProductCategoryStore, userStore types.UserStore) *Handler {
	return &Handler{
		store: store,
		categoryStore: categoryStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/category", auth.WithAdminJWTAuth(h.handleCreateProductCategory, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var productCategory types.CreateUpdateProductCategoryPayload

	if err := utils.ParseJSON(r, &productCategory); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(productCategory); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.categoryStore.CreateShopCategory(productCategory); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	utils.WriteJSON(w, http.StatusOK, productCategory)
}