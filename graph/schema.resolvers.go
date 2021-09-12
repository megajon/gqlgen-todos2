package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/megajon/gqlgen-todos2/graph/generated"
	"github.com/megajon/gqlgen-todos2/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   input.UserID,
		Text: input.Text,
	}
	return r.TodosRepo.CreateTodo(todo)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   input.ID,
		Name: input.Name,
	}
	return r.UsersRepo.CreateUser(user)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodosRepo.GetTodos()
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UsersRepo.GetUsers()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
