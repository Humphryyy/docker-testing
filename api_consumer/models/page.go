package models

type Page struct {
	Model

	Title   string   `json:"title"`
	Widgets []Widget `json:"widgets,omitempty" gorm:"foreignKey:PageID"`

	WorkspaceID uint `json:"workspace_id"`
}
