package models

import (
	"github.com/jinzhu/gorm"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model `json:"-"`
	UUID       string `gorm:"not null;unique" json:"uuid"`
	Name       string `gorm:"not null;unique" json:"username"`
	Pass       string `gorm:"not null" json:"-"`
	Salt       string `gorm:"not null" json:"-"`
}

// UserFactory struct
type UserFactory struct {
	db *DB
}

// NewUserManager - Create a new *UserManager that can be used for managing users.
func NewUserManager(db *DB) (*UserFactory, error) {
	db.AutoMigrate(&User{})
	return &UserFactory{db: db}, nil
}

// Has checks if the given user exists.
func (uf *UserFactory) Has(name string) bool {
	err := uf.db.Where("name=?", name).Find(&User{}).Error
	return err == nil
}

// Get user by name
func (uf *UserFactory) Get(name string) *User {
	u := User{}
	uf.db.Where("name=?", name).Find(&u)
	return &u
}

// GetByID user by ID
func (uf *UserFactory) GetByID(id string) *User {
	u := User{}
	uf.db.Where("uuid=?", id).Find(&u)
	return &u
}

// Add - Creates a user and hashes the password
func (uf *UserFactory) Add(name, pass string) *User {
	salt, _ := uuid.NewV4() // FIXME: is this a good salt?
	guid, _ := uuid.NewV4()
	passwordHash := uf.HashPassword(pass, salt.String())
	user := &User{
		UUID: guid.String(),
		Name: name,
		Pass: passwordHash,
		Salt: salt.String(),
	}
	uf.db.Create(&user)
	return user
}

// HashPassword hashes salt + password
func (uf *UserFactory) HashPassword(pass, salt string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(salt+pass), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

// CheckPassword compare a hashed password with a possible plaintext equivalent,
// fetching the salt for user
func (uf *UserFactory) CheckPassword(salt, hash, pass string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+pass)) != nil {
		return false
	}
	return true
}
