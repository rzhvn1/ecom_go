package shopcategory

import (
	"database/sql"
	"ecom_go/types"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetShopCategoryByID(shopCategoryId int) (*types.ShopCategory, error) {
	rows, err := s.db.Query("SELECT * FROM shopcategories WHERE id = ?", shopCategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	shopCategory := new(types.ShopCategory)
	if rows.Next() {
		shopCategory, err = scanRowsIntoShopCategory(rows)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("shop category not found")
	}

	return shopCategory, nil
}

func (s *Store) CreateShopCategory(shopCategory types.CreateShopCategoryPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO shopcategories (name) VALUES (?)", shopCategory.Name)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoShopCategory(rows *sql.Rows) (*types.ShopCategory, error) {
	shopCategory := new(types.ShopCategory)

	err := rows.Scan(
		&shopCategory.ID,
		&shopCategory.Name,
		&shopCategory.CreatedAt,
		&shopCategory.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return shopCategory, nil
}