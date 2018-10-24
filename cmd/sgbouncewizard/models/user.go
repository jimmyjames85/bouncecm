package models

// DB Schema of the user
type User struct {
	Id 			int 	`json:"id"`;
	First_name 	string 	`json:"first_name"`;
	Last_name 	string 	`json:"last_name"`;
	Email 		string 	`json:"email"`;
	Role 		string 	`json:"role"`;
	Hash 		string 	`json:"hash"`;
	Created_at 	string 	`json:"created_at"`;
}

// Used for storing credentials from the HTTP request
type UserCredentials struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

// Used to marshal the JSON object for HTTP response
type UserObject struct {
	Id 			int 	`json:"id"`;
	First_name 	string 	`json:"first_name"`
	Last_name 	string 	`json:"last_name"`
	Role 		string 	`json:"role"`
}