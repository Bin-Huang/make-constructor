package repositories

// UserRepository the user repository for example
//go:generate go run github.com/Bin-Huang/make-constructor@v0.2.0
type UserRepository struct {
	db        *database
	TableName string
}

// FindByID find something by id
func (r *UserRepository) FindByID() {
}
