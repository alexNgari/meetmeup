package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexNgari/meetmeup/graph/generated"
	"github.com/alexNgari/meetmeup/graph/model"
	"github.com/alexNgari/meetmeup/graph/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	// meetup := &models.Meetup{
	// 	UserID:      input.UserID,
	// 	Name:        input.Name,
	// 	Description: input.Description,
	// }
	// return r.MeetupsRepo.CreateMeetup(meetup)
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	// user := &models.User{
	// 	ID:       fmt.Sprintf("U%d", rand.Int()),
	// 	Username: input.Username,
	// 	Email:    input.Email,
	// 	Meetups:  []*models.Meetup{},
	// }
	// r.users = append(r.users, user)
	// return user, nil
	return nil, nil
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("Meetup does not exist")
	}

	didUpdate := false

	if input.Name != nil {
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("Error while updating meetup: %v", err)
	}
	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupsRepo.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("Meetup does not exist")
	}
	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}
	return true, nil
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserByID(id)
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.UsersRepo.GetMeetupsForUser(obj)
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
