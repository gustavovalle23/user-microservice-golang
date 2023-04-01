package database

import (
	"errors"
	"sync"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type InMemoryUserRepository struct {
	mutex sync.Mutex
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID.Hex()]; ok {
		return errors.New("user already exists")
	}

	r.users[user.ID.Hex()] = user

	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user, ok := r.users[id]
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepository) FindByDocumentNo(documentNo string) (*domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, user := range r.users {
		if user.DocumentNo == documentNo {
			return user, nil
		}
	}

	return nil, domain.ErrUserNotFound
}

func (r *InMemoryUserRepository) Update(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID.Hex()]; !ok {
		return domain.ErrUserNotFound
	}

	r.users[user.ID.Hex()] = user

	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return domain.ErrUserNotFound
	}

	delete(r.users, id)

	return nil
}
