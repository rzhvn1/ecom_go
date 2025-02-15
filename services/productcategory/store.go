package productcategory

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

func (s *Store) CreateShopCategory(productCategory types.CreateUpdateProductCategoryPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO productcategories (name) VALUES (?)", productCategory.Name)
	if err != nil {
		return err
	}

	return nil
}