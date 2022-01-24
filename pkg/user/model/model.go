package model

type User struct {
	ID    uint64
	Name  string
	Email string

	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt uint64
}

type UserDB struct {
	ID    uint64
	Name  string
	Email string

	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt uint64
}
