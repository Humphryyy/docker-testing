package models

import (
	"encoding/json"
	"fmt"
)

type Workspace struct {
	Model

	Name string `json:"name"`

	Hosts []Host `json:"hosts,omitempty" gorm:"foreignKey:WorkspaceID"`

	Pages []Page `json:"pages,omitempty" gorm:"foreignKey:WorkspaceID"`

	Users []User `json:"users,omitempty" gorm:"foreignKey:WorkspaceID"`
	Roles []Role `json:"roles,omitempty" gorm:"foreignKey:WorkspaceID"`
}

type Host struct {
	Model

	WorkspaceID uint `json:"workspace_id"`

	Default  bool   `json:"default"`  // shoeplex.mystandard.co
	Hostname string `json:"hostname"` // platform.shoeplex.com
	IP       string `json:"ip"`
}

func (w Workspace) JSON() ([]byte, error) {
	return json.MarshalIndent(w, "", "  ")
}

func (w Workspace) String() string {
	return fmt.Sprintf(`
	ID: %v
	CreatedAt : %v
	UpdatedAt : %v
	Name: %v
	Hosts: %v
	Users: %v
	Roles: %v
	`, w.ID, w.CreatedAt, w.UpdatedAt, w.Name, w.Hosts, w.Users, w.Roles)
}

func (h Host) String() string {
	return fmt.Sprintf(`
	ID: %v
	CreatedAt : %v
	UpdatedAt : %v
	WorkspaceID: %v
	Default: %v
	Hostname: %v
	IP: %v
	`, h.ID, h.CreatedAt, h.UpdatedAt, h.WorkspaceID, h.Default, h.Hostname, h.IP)
}
