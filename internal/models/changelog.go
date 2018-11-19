package models

// ChangelogEntry - JSON for individual rule from changlog table
type ChangelogEntry struct {
	ID           int    `json:"id"`
	ResponseCode string `json:"response_code"`
	EnhancedCode string `json:"enhanced_code"`
	Regex        string `json:"regex"`
	Priority     int    `json:"priority"`
	Description  string `json:"description"`
	BounceAction string `json:"bounce_action"`
	UserID       string `json:"userid"`
	CreatedAt    string `json:"createdat"`
	Comment      string `json:"comment"`
	
}


type ChangelogTable struct {
	Rules    []ChangelogEntry `json:"rules"`
	NumRules int          `json:"numRules"`
}