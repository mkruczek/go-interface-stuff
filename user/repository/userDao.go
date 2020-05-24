package repository

import (
	"interface-stuff/user"
	"math/rand"
)

var (
	UserDao user.UserStorage
	repo    = map[int]user.User{}
)

func init() {
	UserDao = &userDao{}
}

type userDao struct {
}

func (ud userDao) User(id int) (*user.User, error) {
	result := repo[id]
	return &result, nil
}
func (ud userDao) Save(u *user.User) (*user.User, error) {
	if u.ID == 0 {
		u.ID = rand.Int()
	}
	repo[u.ID] = *u
	return u, nil
}
func (ud userDao) Delete(id int) error {
	delete(repo, id)
	return nil
}
func (ud userDao) Amount() int {
	return len(repo)
}
