package service

import (
	"context"
	"errors"
	"webapp/auth"
	"webapp/entity"
	"webapp/lib"
)

type UserService struct {
	repo         UserRepository
	recaptchaSrv lib.Recaptcha
}

func NewUserService(repo UserRepository, recaptchaCli lib.Recaptcha) *UserService {
	return &UserService{repo: repo, recaptchaSrv: recaptchaCli}
}

func (us *UserService) RegisterUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	user.Id = uid

	user, err := us.repo.Register(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserById(ctx context.Context) (*entity.User, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}

	user, err := us.repo.GetById(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	uid, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, errors.New("user id not found")
	}
	user.Id = uid

	updatedUser, err := us.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
