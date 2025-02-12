package shopcategory

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

func (s *Store) CreateShopCategory(shopcategory types.CreateShopCategoryPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO shopcategories (name) VALUES (?)", shopcategory.Name)
	if err != nil {
		return err
	}

	return nil
}