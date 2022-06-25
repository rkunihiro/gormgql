package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/rkunihiro/gormgql/entity"
	"github.com/rkunihiro/gormgql/generated"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *generated.CreatePostInput) (*entity.Post, error) {
	var posted *time.Time = nil
	if input.Posted != nil {
		t := time.Time(*input.Posted)
		posted = &t
	}
	return r.postRepo.Create(input.AuthorID, input.Title, posted)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
