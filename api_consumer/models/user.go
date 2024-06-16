package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Model

	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`

	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`

	WorkspaceID uint   `json:"workspace_id"`
	Roles       []Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

func (u User) JSON() ([]byte, error) {
	return json.MarshalIndent(u, "", "  ")
}

func (u User) String() string {
	return fmt.Sprintf(`
	ID: %v
	CreatedAt : %v
	UpdatedAt : %v
	Email: %v
	Username: %v
	FirstName: %v
	LastName: %v
	AvatarURL: %v
	Roles: %v
	`, u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.Username, u.FirstName, u.LastName, u.AvatarURL, u.Roles)
}
