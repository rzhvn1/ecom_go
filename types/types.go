package types

import "time"

type BaseTimeModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Password    string `json:"password"`
	BaseTimeModel
}

type UserProfile struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	BaseTimeModel
}

type ShopCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseTimeModel
}

type Shop struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
	Opens_at    string `json:"opens_at"`
	Closes_at   string `json:"closes_at"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	BaseTimeModel
}

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseTimeModel
}

type Product struct {
	ID          int    `json:"id"`
	ShopID      int    `json:"shop_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
	Quantity    int    `json:"quantity"`
	Image       string `json:"image"`
	BaseTimeModel
}

type UserStore interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
}

type ShopCategoryStore interface {
	GetShopCategoryByID(shopCategoryId int) (*ShopCategory, error)
	CreateShopCategory(shopcategory CreateUpdateShopCategoryPayload) error
	UpdateShopCategory(categoryID int, shopCategory CreateUpdateShopCategoryPayload) error
	DeleteShopCategory(categoryID int) (int64, error)
}

type ShopStore interface {
	GetShopByID(shopID int) (*Shop, error)
	CreateShop(shop CreateShopPayload) error
	UpdateShop(shopID int, shop UpdateShopPayload) error
	DeleteShop(shopID int) (int64, error)
}

type ProductCategoryStore interface {
	CreateShopCategory(productCategory CreateUpdateProductCategoryPayload) error
}

type ProductStore interface {
}

type RegisterUserPayload struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}

type CreateUpdateShopCategoryPayload struct {
	Name string `json:"name"`
}

type CreateShopPayload struct {
	UserID      int    `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Opens_at    string `json:"opens_at,omitempty"`
	Closes_at   string `json:"closes_at,omitempty"`
	Address     string `json:"address,omitempty" validate:"omitempty,min=5"`
	Image       string `json:"image,omitempty" validate:"omitempty,url"`
}

type UpdateShopPayload struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
	Opens_at    *string `json:"opens_at,omitempty"`
	Closes_at   *string `json:"closes_at,omitempty"`
	Address     *string `json:"address,omitempty" validate:"omitempty,min=5"`
	Image       *string `json:"image,omitempty" validate:"omitempty,url"`
}

type CreateUpdateProductCategoryPayload struct {
	Name string `json:"name"`
}
