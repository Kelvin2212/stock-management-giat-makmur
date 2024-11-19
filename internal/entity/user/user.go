package user

import "time"

type User struct {
	ID            uint      `db:"id" json:"id"`
	FirstName     string   `db:"first_name" json:"first_name"`
	LastName      string   `db:"last_name" json:"last_name"`
	Email         string   `db:"email" json:"email"`
	Password      string   `db:"password" json:"password`
	PhoneNumber   uint    `db:"phone_number json:phone_number`
	Active_status bool     `db:"active_status" json:"active_status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}