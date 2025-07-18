package models

type Event struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Location    string `json:"location"`
	Rules       string `gorm:"type:text" json:"rules"` // Stored as JSON string
	Capacity    int    `json:"capacity"`
	CreatorID   uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"creator_id"`
	ImageURL    string `json:"image_url"`
	HostName    string `gorm:"-" json:"host_name,omitempty"` // Ignored by GORM
	Participants int64  `gorm:"-" json:"participants"` // Ignored by GORM
}
