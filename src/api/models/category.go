package models

// Category stores information about a category
type Category struct {
	ID          int
	Name        string
	Description string
}

// NewCategory creates a new instance of a category
func NewCategory(name, description string) *Category {
	return &Category{
		ID:          1,
		Name:        name,
		Description: description,
	}
}
