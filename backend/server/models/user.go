package models

// Defines a user struct to recieve user objects from the database
type User struct {
	// Note: golang requires struct fields to be uppercase to be exported
	// (lowercase field names are unexported and can't be used outside of this file )
	User_id  int
	Email    string
	Username string
	Password string
}
