package resolver

import (
	"github.com/rkunihiro/gormgql/repository"
)

type Resolver struct {
	userRepo repository.UserRepository
	postRepo repository.PostRepository
}

func NewResolver(
	userRepo repository.UserRepository,
	postRepo repository.PostRepository,
) *Resolver {
	return &Resolver{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}
