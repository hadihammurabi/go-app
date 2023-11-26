package repository

// Repository struct
type Repository struct {
}

var r *Repository

// NewRepository func
func NewRepository() *Repository {
	return &Repository{}
}

func Get() *Repository {
	if r != nil {
		return r
	}

	r = NewRepository()
	return r
}
