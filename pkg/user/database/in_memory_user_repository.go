package database

import (
	"errors"
	"sync"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type InMemoryUserRepository struct {
	mutex sync.Mutex
	users map[int]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user

	return nil
}

func (r *InMemoryUserRepository) FindByID(id int) (*domain.User, error) {
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

func (r *InMemoryUserRepository) FindByEmail(email string) (*domain.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, domain.ErrUserNotFound
}

func (r *InMemoryUserRepository) Update(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return domain.ErrUserNotFound
	}

	r.users[user.ID] = user

	return nil
}

func (r *InMemoryUserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[id]; !ok {
		return domain.ErrUserNotFound
	}

	delete(r.users, id)

	return nil
}
