package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ketan-10/keep/entities"
)

type Resolver struct {
	notes []*entities.Note
}

// func (r *mutationResolver) CreateTodo(ctx context.Context, input entities.NewTodo) (*entities.Todo, error) {
// 	panic("not implemented")
// }

// func (r *queryResolver) Todos(ctx context.Context) ([]*entities.Todo, error) {
// 	panic("not implemented")
// }

// implement resolvers
func (r *mutationResolver) AddNote(ctx context.Context, input *entities.NewNote) (*entities.Note, error) {
	newNote := &entities.Note{
		// ID:          fmt.Sprintf("T%d", rand.Int()),
		ID:          rand.Int(),
		Text:        input.Text,
		Description: input.Description,
	}

	r.notes = append(r.notes, newNote)
	return newNote, nil
}

func (r *queryResolver) GetAllNotes(ctx context.Context) ([]*entities.Note, error) {
	return r.notes, nil
}

// func (r *mutationResolver) AddSubNote(ctx context.Context, text string, description string, IDParent string) (*entities.SubNote, error) {
func (r *mutationResolver) AddSubNote(ctx context.Context, text string, description string, IDParent int) (*entities.SubNote, error) {
	var note *entities.Note
	for _, n := range r.notes {
		if n.ID == IDParent {
			note = n
			break
		}
	}

	if note == nil {
		panic("Parent note not found")
	}

	newSubNote := &entities.SubNote{
		ID: fmt.Sprintf("T%d", rand.Int()),
		// ID:          rand.Int(),
		Text:        text,
		Description: description,
	}

	note.SubNote = append(note.SubNote, newSubNote)
	return newSubNote, nil
}

// // Mutation returns MutationResolver implementation.
// func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// // Query returns QueryResolver implementation.
// func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{
		r,
	}
}

// why this is not acceptable
// it it recommended for a function to return Interface?

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{
		r,
	}
}

type queryResolver struct {
	*Resolver
}

type mutationResolver struct {
	*Resolver
}

// expect a RootResolver As interface, with 2 functions
// we pass the struct
