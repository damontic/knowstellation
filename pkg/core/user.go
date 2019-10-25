package core

// User represents a User of the system
type User struct {
	Name string `json:"name"`
}

// NewUser creates a User
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
