package store

import "time"

type Store struct {
	ID           uint      `db:"id" json:"id"`
	LogoImageUrl string    `db:"logo_image_url" json:"logo_image_url"`
	StoreName    string    `db:"store_name" json:"store_name"`
	StoreAddress string    `db:"store_address" json:"store_address"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	CreatedBy    string    `db:"created_by" json:"created_by"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    string    `db:"updated_by" json:"updated_by"`
}