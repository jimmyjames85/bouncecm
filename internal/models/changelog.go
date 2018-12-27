package models

// ChangelogEntry - JSON for individual rule from changlog table
type ChangelogEntry struct {
	ID           int    `json:"rule_id"`
	ResponseCode int `json:"response_code"`
	EnhancedCode string `json:"enhanced_code"`
	Regex        string `json:"regex"`
	Priority     int    `json:"priority"`
	Description  string `json:"description"`
	BounceAction string `json:"bounce_action"`
	UserID       int `json:"user_id"`
	CreatedAt    int `json:"created_at"`
	Comment      string `json:"comment"`
	
}