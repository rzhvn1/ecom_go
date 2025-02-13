package shop

import (
	"ecom_go/services/auth"
	"ecom_go/types"
	"ecom_go/utils"
	"fmt"
	"net/http"
	"strconv"

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
	router.HandleFunc("/{shop_id}", auth.WithJWTAuth(h.handleGetShop, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/{shop_id}", auth.WithJWTAuth(h.handleUpdateShop, h.userStore)).Methods(http.MethodPut)
	router.HandleFunc("/{shop_id}", auth.WithJWTAuth(h.handleDeleteShop, h.userStore)).Methods(http.MethodDelete)

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

func (h *Handler) handleGetShop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["shop_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing shop id"))
		return
	}

	shopID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid shop ID"))
		return
	}

	shop, err := h.store.GetShopByID(shopID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, shop)
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

// ToDo: add owner permissions
func (h *Handler) handleUpdateShop(w http.ResponseWriter, r *http.Request) {
	var shop types.UpdateShopPayload
	
	vars := mux.Vars(r)
	str, ok := vars["shop_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing shop ID"))
		return
	}

	shopID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid shop ID"))
		return
	}

	existingShop, err := h.store.GetShopByID(shopID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	if err := utils.ParseJSON(r, &shop); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(shop); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if shop.CategoryID != nil {
		_, err := h.categoryStore.GetShopCategoryByID(*shop.CategoryID)
		if err != nil {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("shop category not found"))
			return
		}	
	}

	if shop.Name == nil {
		shop.Name = &existingShop.Name
	}
	if shop.Description == nil {
		shop.Description = &existingShop.Description
	}
	if shop.CategoryID == nil {
		shop.CategoryID = &existingShop.CategoryID
	}
	if shop.Opens_at == nil {
		shop.Opens_at = existingShop.Opens_at
	}
	if shop.Closes_at == nil {
		shop.Closes_at = existingShop.Closes_at
	}
	if shop.Address == nil {
		shop.Address = &existingShop.Address
	}
	if shop.Image == nil {
		shop.Image = &existingShop.Image
	}

	err = h.store.UpdateShop(shopID, shop)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	updatedShop, _ := h.store.GetShopByID(shopID)
	utils.WriteJSON(w, http.StatusOK, updatedShop)
}

// ToDo: add owner permissions
func (h *Handler) handleDeleteShop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["shop_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing shop ID"))
		return
	}

	shopID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid shop ID"))
		return
	}

	rowsAffected, err := h.store.DeleteShop(shopID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to delete shop: %v", err))
		return
	}

	if rowsAffected == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("shop not found"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}