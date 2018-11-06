package models

// BounceRule - JSON for individual rule
type BounceRule struct {
	ID           int    `json:"id"`
	ResponseCode string `json:"response_code"`
	EnhancedCode string `json:"enhanced_code"`
	Regex        string `json:"regex"`
	Priority     int    `json:"priority"`
	Description  []byte `json:"description"`
	BounceAction string `json:"bounce_action"`
}

// RulesObject - JSON for db
type RulesObject struct {
	Rules    []BounceRule `json:"rules"`
	NumRules int          `json:"numRules"`
}
