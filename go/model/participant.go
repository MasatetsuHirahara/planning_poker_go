package model

type (

	// Participant : 参加者
	Participant struct {
		ID          string  `json:"id"`
		JoinedRoom  string  `json:"joined_room"`
		Name        string  `json:"name"`
		Points      []Point `json:"points"`
		IsModerator bool    `json:"is_moderator"`
	}

	// Point : ポイント
	Point struct {
		ID    string `json:"id"`
		Point int    `json:"point"`
	}
)
