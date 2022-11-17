package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rkunihiro/gormgql/entity"
	"github.com/rkunihiro/gormgql/generated"
	"github.com/rkunihiro/gormgql/scalar"
)

// Posted is the resolver for the posted field.
func (r *postResolver) Posted(ctx context.Context, obj *entity.Post) (*scalar.DateTime, error) {
	dt := scalar.DateTime(obj.Posted)
	return &dt, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
