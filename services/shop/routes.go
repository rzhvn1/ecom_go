package shop

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
	store         types.ShopStore
	categoryStore types.ShopCategoryStore
	userStore     types.UserStore
}

func NewHandler(store types.ShopStore, categoryStore types.ShopCategoryStore, userStore types.UserStore) *Handler {
	return &Handler{
		store: store,
		categoryStore: categoryStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", auth.WithJWTAuth(h.handleCreateShop, h.userStore)).Methods(http.MethodPost)

	router.HandleFunc("/category", auth.WithAdminJWTAuth(h.handleCreateShopCategory, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCreateShopCategory(w http.ResponseWriter, r *http.Request) {
	var shopCategory types.CreateShopCategoryPayload
	if err := utils.ParseJSON(r, &shopCategory); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(shopCategory); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.categoryStore.CreateShopCategory(shopCategory)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, shopCategory)
}

func (h *Handler) handleCreateShop(w http.ResponseWriter, r *http.Request) {
	var shop types.CreateShopPayload
	if err := utils.ParseJSON(r, &shop); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(shop); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if _, err := h.categoryStore.GetShopCategoryByID(shop.CategoryID); err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("invalid payload: %v", err))
		return
	}
	
	err := h.store.CreateShop(shop)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, shop)
}