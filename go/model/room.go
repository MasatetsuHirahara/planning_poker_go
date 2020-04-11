package model

type (
	// Room : room
	Room struct {
		ID                 string   `json:"id"`
		No                 string   `json:"no"`
		Title              string   `json:"title"`
		Requirement        string   `json:"requirement"`
		Participants       []string `json:"participants"`
		CurrentRequirement string   `json:"current_requirement"`
		ModeratorID        string   `json:"moderator_id"`
	}

	// Requirement : 要件
	Requirement struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		IsFinish    string `json:"is_finish"`
	}
)
