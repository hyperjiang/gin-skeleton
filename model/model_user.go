package model

import "time"

// User the user model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName for gorm
func (User) TableName() string {
	return "users"
}

// GetFirstByID gets the user by his ID
func (u *User) GetFirstByID(id string) error {
	db := DB().Where("id=?", id).First(u)

	if db.RecordNotFound() {
		return ErrDataNotFound
	} else if db.Error != nil {
		return db.Error
	}

	return nil
}

// Create a new user
func (u *User) Create() error {
	db := DB().Create(u)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}
