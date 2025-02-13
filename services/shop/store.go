package shop

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

func (s *Store) GetShopByID(shopID int) (*types.Shop, error) {
	rows, err := s.db.Query("SELECT * FROM shops WHERE id = ?", shopID)
	if err != nil {
		return nil, err
	}

	shop := new(types.Shop)
	for rows.Next() {
		shop, err = scanRowsIntoShop(rows)
		if err != nil {
			return nil, err
		}
	}

	if shop.ID == 0 {
		return nil, fmt.Errorf("shop not found")
	}
	
	return shop, nil
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

func (s *Store) UpdateShop(shopID int, shop types.UpdateShopPayload) error {
	_, err := s.db.Exec(
		"UPDATE shops SET name = ?, description = ?, category_id = ?, opens_at = ?, closes_at = ?, address = ?, image = ? WHERE id = ?",
		shop.Name, shop.Description, shop.CategoryID, shop.Opens_at, shop.Closes_at, shop.Address, shop.Image, shopID)
	
	return err
}


func scanRowsIntoShop(rows *sql.Rows) (*types.Shop, error) {
	shop := new(types.Shop)

	err := rows.Scan(
		&shop.ID,
		&shop.UserID,
		&shop.Name,
		&shop.Description,
		&shop.CategoryID,
		&shop.Opens_at,
		&shop.Closes_at,
		&shop.Address,
		&shop.Image,
		&shop.CreatedAt,
		&shop.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *Store) DeleteShop(shopID int) (int64, error) {
	result, err := s.db.Exec("DELETE FROM shops WHERE id = ?", shopID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rowsAffected, nil
}