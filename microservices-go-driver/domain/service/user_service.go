package service

// import (
// 	"context"
// 	"errors"
// 	"sync"

// 	"github.com/microservices/domain/model"
// )

// var (
// 	ErrInconsistentIDs = errors.New("inconsistent IDs")
// 	ErrAlreadyExists   = errors.New("already exists")
// 	ErrNotFound        = errors.New("not found")
// )

// type UserService interface {
// 	CreateUser(ctx context.Context, user *model.User) error
// 	GetUserById(ctx context.Context, id int) (model.User, error)
// 	DeleteUser(ctx context.Context, id int) error
// 	UpdateUser(ctx context.Context, id int, user *model.User) error
// }

// type userService struct {
// 	mtx sync.RWMutex
// 	m   map[int]model.User
// }

// func NewService() UserService {
// 	return &userService{
// 		m: map[int]model.User{},
// 	}
// }

// func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
// 	s.mtx.Lock()
// 	defer s.mtx.Unlock()
// 	if _, ok := s.m[user.Userid]; ok {
// 		return ErrAlreadyExists
// 	}
// 	s.m[user.Userid] = *user
// 	return nil

// 	// err := s.repository.CreateUser(ctx, user)

// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// return nil
// }

// func (s *userService) GetUserById(ctx context.Context, id int) (model.User, error) {

// 	s.mtx.RLock()
// 	defer s.mtx.RUnlock()
// 	p, ok := s.m[id]
// 	if !ok {
// 		return model.User{}, ErrNotFound
// 	}
// 	return p, nil

// 	// var data model.User
// 	// data, err := s.repository.GetUserById(ctx, id)

// 	// if err != nil {
// 	// 	return data, err
// 	// }

// 	// return data, nil
// }

// func (s *userService) DeleteUser(ctx context.Context, id int) error {

// 	// err := s.repository.DeleteUser(ctx, id)

// 	// if err != nil {
// 	// 	return err
// 	// }

// 	s.mtx.Lock()
// 	defer s.mtx.Unlock()
// 	if _, ok := s.m[id]; !ok {
// 		return ErrNotFound
// 	}
// 	delete(s.m, id)

// 	return nil
// }

// func (s *userService) UpdateUser(ctx context.Context, id int, user *model.User) error {

// 	if id != user.Userid {
// 		return ErrInconsistentIDs
// 	}
// 	s.mtx.Lock()
// 	defer s.mtx.Unlock()
// 	s.m[id] = *user
// 	return nil

// 	// err := s.repository.UpdateUser(ctx, id, user)

// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// return nil
// }
