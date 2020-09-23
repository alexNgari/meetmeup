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

func (u *UsersRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := u.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}