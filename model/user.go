package model

import (
	"log"

	"github.com/eikendev/pushbits/authentication/credentials"
)

// User holds information like the name, the secret, and the applications of a user.
type User struct {
	ID           uint   `gorm:"AUTO_INCREMENT;primary_key"`
	Name         string `gorm:"type:string;size:128;unique"`
	PasswordHash []byte
	IsAdmin      bool
	MatrixID     string `gorm:"type:string"`
	Applications []Application
}

// ExternalUser represents a user for external purposes.
type ExternalUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" form:"name" query:"name" binding:"required"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin" query:"is_admin"`
	MatrixID string `json:"matrix_id" form:"matrix_id" query:"matrix_id" binding:"required"`
}

// UserCredentials holds information for authenticating a user.
type UserCredentials struct {
	Password string `json:"password,omitempty" form:"password" query:"password" binding:"required"`
}

// ExternalUserWithCredentials represents a user for external purposes and includes the user's credentials in plaintext.
type ExternalUserWithCredentials struct {
	ExternalUser
	UserCredentials
}

// NewUser creates a new user.
func NewUser(name, password string, isAdmin bool, matrixID string) *User {
	log.Printf("Creating user %s.\n", name)

	user := User{
		Name:         name,
		PasswordHash: credentials.CreatePassword(password),
		IsAdmin:      isAdmin,
		MatrixID:     matrixID,
	}

	return &user
}

// IntoInternalUser converts a ExternalUserWithCredentials into a User.
func (u *ExternalUserWithCredentials) IntoInternalUser() *User {
	return &User{
		Name:         u.Name,
		PasswordHash: credentials.CreatePassword(u.Password),
		IsAdmin:      u.IsAdmin,
		MatrixID:     u.MatrixID,
	}
}

// IntoExternalUser converts a User into a ExternalUser.
func (u *User) IntoExternalUser() *ExternalUser {
	return &ExternalUser{
		ID:       u.ID,
		Name:     u.Name,
		IsAdmin:  u.IsAdmin,
		MatrixID: u.MatrixID,
	}
}
