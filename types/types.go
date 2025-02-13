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

type Shop struct {
	ID            int        `json:"id"`
	UserID        int        `json:"user_id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	CategoryID    int        `json:"category_id"`
	Opens_at      *time.Time `json:"opens_at"`
	Closes_at     *time.Time `json:"closes_at"`
	Address       string     `json:"address"`
	Image         string     `json:"image"`
	BaseTimeModel
}

type ShopCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseTimeModel
}

type UserStore interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
}

type ShopStore interface {
	GetShopByID(shopID int) (*Shop, error)
	CreateShop(shop CreateShopPayload) error
	UpdateShop(shopID int, shop UpdateShopPayload) error
	DeleteShop(shopID int) (int64, error)
}

type ShopCategoryStore interface {
	GetShopCategoryByID(shopCategoryId int) (*ShopCategory, error)
	CreateShopCategory(shopcategory CreateShopCategoryPayload) error
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

type CreateShopCategoryPayload struct {
	Name string `json:"name"`
}

type CreateShopPayload struct {
	UserID        int        `json:"user_id" validate:"required"`
	Name          string     `json:"name" validate:"required"`
	Description   string     `json:"description,omitempty"`
	CategoryID    int        `json:"category_id" validate:"required"`
	Opens_at      *time.Time `json:"opens_at,omitempty"`
	Closes_at     *time.Time `json:"closes_at,omitempty"`
	Address       string     `json:"address,omitempty" validate:"omitempty,min=5"`
	Image         string     `json:"image,omitempty" validate:"omitempty,url"`
}

type UpdateShopPayload struct {
	Name          *string     `json:"name,omitempty"`
	Description   *string     `json:"description,omitempty"`
	CategoryID    *int        `json:"category_id,omitempty"`
	Opens_at      *time.Time `json:"opens_at,omitempty"`
	Closes_at     *time.Time `json:"closes_at,omitempty"`
	Address       *string     `json:"address,omitempty" validate:"omitempty,min=5"`
	Image         *string     `json:"image,omitempty" validate:"omitempty,url"`
}