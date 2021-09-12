package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/megajon/gqlgen-todos2/graph/model"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?").First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := u.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (t *UsersRepo) CreateUser(user *model.User) (*model.User, error) {
	_, err := t.DB.Model(user).Returning("*").Insert()
	return user, err
}
