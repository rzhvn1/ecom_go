package product

import "ecom_go/types"

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