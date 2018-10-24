package models

type BounceRule struct {
	Id 				int 	`json:"id"`;
	Response_code 	string 	`json:"response_code"`;
	Enhanced_code 	string 	`json:"enhanced_code"`;
	Regex 			string 	`json:"regex"`;
	Priority 		int 	`json:"priority"`;
	Description 	[]byte 	`json:"description"`;
	Bounce_action 	string 	`json:"bounce_action"`;
}

type RulesObject struct {
	Rules 			[]BounceRule 	`json:"rules"`;
	NumRules 		int				`json:"numRules"`
}