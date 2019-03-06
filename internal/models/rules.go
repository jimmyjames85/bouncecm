package models

// BounceRule - JSON for individual rule
type BounceRule struct {
	ID           int    `json:"id"`
	ResponseCode int    `json:"response_code"`
	EnhancedCode string `json:"enhanced_code"`
	Regex        string `json:"regex"`
	Priority     int    `json:"priority"`
	Description  string `json:"description"`
	BounceAction string `json:"bounce_action"`
}

type ChangelogEntry struct {
	BounceRule
	UserID    int    `json:"user_id"`
	CreatedAt int32    `json:"created_at"`
	Comment   string `json:"comment"`
	Operation string `json:"operation"`
}
