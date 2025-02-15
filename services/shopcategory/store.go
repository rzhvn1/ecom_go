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

func (s *Store) CreateShopCategory(shopCategory types.CreateUpdateShopCategoryPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO shopcategories (name) VALUES (?)", shopCategory.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateShopCategory(categoryID int, shopCategory types.CreateUpdateShopCategoryPayload) error {
	_, err := s.db.Exec("UPDATE shopcategories SET name = ? WHERE id = ?", shopCategory.Name, categoryID)

	return err
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

func (s *Store) DeleteShopCategory(categoryID int) (int64, error) {
	result, err := s.db.Exec("DELETE FROM shopcategories WHERE id = ?", categoryID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rowsAffected, nil
}
