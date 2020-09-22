package graph

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/alexNgari/meetmeup/graph/models"
)

type Resolver struct{
	meetups []*models.Meetup
	// {
	// 	ID: 1,
	// 	Name: "A meetup",
	// 	Description: "A description",
	// 	UserID: "1"
	// },
	// {
	// 	ID: "2",
	// 	Name: "A second meetup",
	// 	Description: "A description",
	// 	UserID: "2"
	// }

	users []*models.User
	// {
	// 	ID: 1,
	// 	Username: "Bob",
	// 	Email: "bob@gmail.com"
	// },
	// {
	// 	ID: "2",
	// 	Username: "Jon",
	// 	Email: "jon@gmail.com"
	// }
}
