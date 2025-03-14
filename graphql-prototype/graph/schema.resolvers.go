package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/bayu-wibowo-mf/personal-go-graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", randNumber),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
	}

	_, err := r.DB.Model(&movie).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new movie %v", err)
	}

	return &movie, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie

	err := r.DB.Model(&movies).Select()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// EmployeeDemography is the resolver for the employeeDemography field.
func (r *queryResolver) EmployeeDemography(ctx context.Context) ([]*model.EmployeeDemography, error) {
	var employeeDemography []*model.EmployeeDemography

	counter, err := r.DB.Model(&employeeDemography).Table("dm_employee_demography").Count()
	if err != nil {
		return nil, err
	}
	print(counter)

	err = r.DB.Model(&employeeDemography).Table("dm_employee_demography").Select()
	if err != nil {
		return nil, err
	}
	return employeeDemography, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "User " + obj.UserID}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
