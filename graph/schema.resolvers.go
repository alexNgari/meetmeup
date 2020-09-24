package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/alexNgari/meetmeup/graph/generated"
	"github.com/alexNgari/meetmeup/graph/models"
	customMiddleware "github.com/alexNgari/meetmeup/middleware"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := customMiddleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	meetup := &models.Meetup{
		UserID:      currentUser.ID,
		Name:        input.Name,
		Description: input.Description,
	}
	return r.MeetupsRepo.CreateMeetup(meetup)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
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

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
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

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	_, err = r.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already in use")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, err
	}

	// @TODO: create verifivation code

	tx, err := r.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("Something went wrong")
	}
	defer tx.Rollback()

	if _, err := r.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while committing: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrBadCredentials  = errors.New("email/password combination do not match")
	ErrUnauthenticated = errors.New("not authorized to view this content")
)
