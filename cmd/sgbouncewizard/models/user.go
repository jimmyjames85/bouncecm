package models

// User - DB Schema of the user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Hash      string `json:"hash"`
	CreatedAt string `json:"created_at"`
}

// UserCredentials - Used for storing credentials from the HTTP request
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserObject - Used to marshal the JSON object for HTTP response
type UserObject struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}
