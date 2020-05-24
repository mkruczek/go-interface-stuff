package user

type User struct {
	ID    int
	Login string
}

type UserStorage interface {
	User(id int) (*User, error)
	Save(u *User) (*User, error)
	Delete(id int) error
	Amount() int
}
