package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"
	_ "time/tzdata"

	"github.com/rkunihiro/gormgql/entity"
	"github.com/rkunihiro/gormgql/generated"
	"github.com/rkunihiro/gormgql/scalar"
)

func (r *queryResolver) Message(ctx context.Context) (string, error) {
	return "Hello,World!", nil
}

func (r *queryResolver) Now(ctx context.Context, timezone string) (*scalar.DateTime, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	dt := scalar.DateTime(time.Now().In(loc))
	return &dt, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*entity.Post, error) {
	return r.Resolver.postRepo.FindByID(id)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*entity.Post, error) {
	return r.Resolver.postRepo.Find()
}

func (r *queryResolver) User(ctx context.Context, id int) (*entity.User, error) {
	return r.Resolver.userRepo.FindByID(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*entity.User, error) {
	return r.Resolver.userRepo.Find()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
