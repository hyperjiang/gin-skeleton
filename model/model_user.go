package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

// GetFirstByEmail gets the user by his email
func (u *User) GetFirstByEmail(email string) error {
	db := DB().Where("email=?", email).First(u)

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

// Signup a new user
func (u *User) Signup() error {
	var user User
	err := user.GetFirstByEmail(u.Email)

	if err == nil {
		return ErrUserExists
	} else if err != ErrDataNotFound {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// replace the plaintext password with ciphertext password
	u.Password = string(hash)

	return u.Create()
}

// Login a user
func (u *User) Login(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// LoginByEmailAndPassword login a user by his email and password
func LoginByEmailAndPassword(email, password string) (User, error) {
	var user User
	err := user.GetFirstByEmail(email)
	if err != nil {
		return user, err
	}

	return user, user.Login(password)
}
