package store

var NewStore Store

// Result struct will contain the search result
type Result struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Store is an interface for any database
// that supports the below methods
type Store interface {
	Get(name, email string, page int) ([]Result, error)
	AddContact(name, email string) error
	DeleteContact(email string) error
	UpdateContact(name, email string) error
}
