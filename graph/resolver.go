package graph

import (
	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	UserRepo domain.UserRepository
}

func NewResolver(UserRepo domain.UserRepository) *Resolver {
	return &Resolver{
		UserRepo: UserRepo,
	}
}
