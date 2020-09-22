package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"errors"

	"github.com/alexNgari/meetmeup/graph/generated"
	"github.com/alexNgari/meetmeup/graph/model"
	"github.com/alexNgari/meetmeup/graph/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == obj.UserID {
			return user, nil
		}
	}
	return nil, errors.New("User does not exist")
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	meetup := &models.Meetup{
		UserID:      input.UserID,
		Name:        input.Name,
		Description: input.Description,
		ID:          fmt.Sprintf("M%d", rand.Int()),
	}
	r.meetups = append(r.meetups, meetup)
	return meetup, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user := &models.User{
		ID: fmt.Sprintf("U%d", rand.Int()),
		Username: input.Username,
		Email: input.Email,
		Meetups: []*models.Meetup{},
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	userMeetups := []*models.Meetup{}
	for _, meetup := range r.meetups {
		if meetup.UserID == obj.ID {
			userMeetups = append(userMeetups, meetup)
		}
	}
	return userMeetups, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
