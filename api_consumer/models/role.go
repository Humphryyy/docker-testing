package models

import (
	"fmt"
)

type RolePermission uint

const (
	RemoveUsers RolePermission = 1 << iota
	BanUsers
	EditWorkspace
)

type Role struct {
	Model

	Name        string         `json:"name"`
	Permissions RolePermission `json:"permissions"`

	WorkspaceID uint `json:"workspace_id"`
}

func (r Role) String() string {
	return fmt.Sprintf(`
	ID: %v
	Name: %v
	Permissions: %v
	WorkspaceID: %v
	`, r.ID, r.Name, r.Permissions, r.WorkspaceID)
}
