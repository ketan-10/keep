package tmp

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/ketan-10/keep/entities"
	graphql1 "github.com/ketan-10/keep/graphql"
)

type Resolver struct{}

// // foo
func (r *mutationResolver) CreateTodo(ctx context.Context, input entities.NewTodo) (*entities.Todo, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Todos(ctx context.Context) ([]*entities.Todo, error) {
	panic("not implemented")
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
