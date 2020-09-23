package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/alexNgari/meetmeup/graph/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := u.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}

func (u *UsersRepo) getUserByField(field, value string) (*models.User, error) {
	var user *models.User = new (models.User)
	err := u.DB.Model(user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return user, err 
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	return u.getUserByField("id", id)
}

// GetUserByEmail email => user
func (u *UsersRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.getUserByField("email", email)
}

func (u *UsersRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.getUserByField("username", username)
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
