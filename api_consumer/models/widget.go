package models

import "github.com/lib/pq"

type WidgetType string

const (
	WidgetTypeDocument WidgetType = "document"
	WidgetTypeCalendar WidgetType = "calendar"
	WidgetTypeGraph    WidgetType = "graph"
	WidgetTypeMap      WidgetType = "map"

	WidgetTypeTextChat  WidgetType = "text_chat"
	WidgetTypeVoiceChat WidgetType = "voice_chat"
	WidgetTypeVideoChat WidgetType = "video_chat"

	WidgetTypeKanban WidgetType = "kanban"
	WidgetTypeForum  WidgetType = "forum"
	WidgetTypePoll   WidgetType = "poll"

	WidgetTypeFile  WidgetType = "file"
	WidgetTypeImage WidgetType = "image"
	WidgetTypeVideo WidgetType = "video"
)

type Widget struct {
	Model

	Type WidgetType `json:"type"`

	WidgetDataID uint       `json:"widget_data_id"`
	WidgetData   WidgetData `json:"widget_data" gorm:"-"`

	LocationX int `json:"location_x"`
	LocationY int `json:"location_y"`
	Width     int `json:"width"`
	Height    int `json:"height"`

	PageID      uint `json:"page_id"`
	WorkspaceID uint `json:"workspace_id"`
}

type WidgetPermissions struct {
	BlockedRoles pq.Int64Array `json:"blocked_roles" gorm:"type:integer[]"`
	BlockedUsers pq.Int64Array `json:"blocked_users" gorm:"type:integer[]"`
}

type WidgetData struct {
	Model

	Permissions WidgetPermissions `json:"permissions" gorm:"embedded"`

	PageID      uint `json:"page_id"`
	WidgetID    uint `json:"widget_id"`
	WorkspaceID uint `json:"workspace_id"`
}

type Document struct {
	WidgetData

	Title   string `json:"title"`
	Content string `json:"content"`
}

type Event struct {
	Model

	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`

	CalendarID uint `json:"calendar_id"`
}

type Calendar struct {
	WidgetData

	Title string `json:"title"`

	Events []Event `json:"events" gorm:"foreignKey:CalendarID"`
}

type GraphPoint struct {
	Model

	X float64 `json:"x"`
	Y float64 `json:"y"`

	GraphID uint `json:"graph_id"`
}

type Graph struct {
	WidgetData

	Title string `json:"title"`

	Points []GraphPoint `json:"points" gorm:"foreignKey:GraphID"`
}

type MapMarker struct {
	Model

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Label     string  `json:"label"`

	MapID uint `json:"map_id"`
}

type Map struct {
	WidgetData

	Title string `json:"title"`

	Markers []MapMarker `json:"markers" gorm:"foreignKey:MapID"`
}

type Message struct {
	Model

	Content string `json:"content"`

	TextChatID uint `json:"text_chat_id"`
}

type TextChat struct {
	WidgetData

	Title string `json:"title"`

	Messages []Message `json:"messages" gorm:"foreignKey:TextChatID"`

	WSUrl string `json:"ws_url"`
}

type VoiceChat struct {
	WidgetData

	Title string `json:"title"`

	WebRTCUrl string `json:"webrtc_url"`
}

type VideoChat struct {
	WidgetData

	Title string `json:"title"`

	WebRTCUrl string `json:"webrtc_url"`
}

type KanbanColumn struct {
	Model

	Title string       `json:"title"`
	Cards []KanbanCard `json:"cards" gorm:"foreignKey:KanbanColumnID"`

	KanbanID uint `json:"kanban_id"`
}

type KanbanCard struct {
	Model

	Title       string `json:"title"`
	Description string `json:"description"`

	KanbanColumnID uint `json:"kanban_column_id"`
}

type Kanban struct {
	WidgetData

	Title string `json:"title"`

	Columns []KanbanColumn `json:"columns" gorm:"foreignKey:KanbanID"`
}

type ForumThread struct {
	Model

	Title string      `json:"title"`
	Posts []ForumPost `json:"posts" gorm:"foreignKey:ForumThreadID"`

	ForumID uint `json:"forum_id"`
}

type ForumPost struct {
	Model

	Content string `json:"content"`

	ForumThreadID uint `json:"forum_thread_id"`
}

type Forum struct {
	WidgetData

	Title string `json:"title"`

	Threads []ForumThread `json:"threads" gorm:"foreignKey:ForumID"`
}
