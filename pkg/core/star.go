package core

// Star represents a set of skills
type Star struct {
	Name string `json:"name"`
}

// NewStar creates a Star
func NewStar(name string) *Star {
	return &Star{
		Name: name,
	}
}
