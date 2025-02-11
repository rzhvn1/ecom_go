package shop

import (
	"database/sql"
	"ecom_go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateShop(shop types.CreateShopPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO shops (user_id, name, description, category_id, opens_at, closes_at, address, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", 
		shop.UserID, shop.Name, shop.Description, shop.CategoryID, shop.Opens_at, shop.Closes_at, shop.Address, shop.Image)
	if err != nil {
		return err
	}

	return nil
}
