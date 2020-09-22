package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/alexNgari/meetmeup/graph/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user *models.User = new (models.User)
	err := u.DB.Model(user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}