package models

type Event struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Time        string   `json:"time"`
	Location    string   `json:"location"`
	Rules       []string `json:"rules"`
	Capacity    int      `json:"capacity"`
	CreatorID   int      `json:"creator_id"`
	HostName    string   `json:"host_name,omitempty"`
	Participants int64     `json:"participants,omitempty"`
}
