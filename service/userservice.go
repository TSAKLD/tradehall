package service

import (
	"errors"
	"steamsale/entity"
)

type UserService interface {
	EditUser(user entity.User) error
	RemoveUser(user entity.User, removeID int) error
	RegisterUser(user entity.User) (entity.User, error)

	FindUser(user entity.UserForFind) (userList entity.User, err error)
}

func (s *service) EditUser(user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) RemoveUser(user entity.User, removeID int) error {
	if user.ID != removeID || user.Status != entity.StatusAdmin {
		return ErrNotAllowed
	}

	_, err := s.repo.FindWithID(removeID)
	if err != nil {
		return err
	}

	err = s.repo.RemoveUser(removeID)
	return err
}

func (s *service) RegisterUser(user entity.User) (entity.User, error) {
	_, err := s.repo.FindWithNickname(user.Nickname)
	if err == nil {
		return user, ErrNicknameTaken
	} else if !errors.Is(err, ErrNotFound) {
		return user, err
	}

	_, err = s.repo.FindWithEmail(user.Email)
	if err == nil {
		return user, ErrEmailTaken
	} else if !errors.Is(err, ErrNotFound) {
		return user, err
	}

	user, err = s.repo.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) FindUser(user entity.UserForFind) (result entity.User, err error) {
	if user.ID != nil {
		result, err := s.repo.FindWithID(*user.ID)
		if err != nil {
			return result, err
		}
	}

	if user.Nickname != nil {
		result, err := s.repo.FindWithNickname(*user.Nickname)
		if err != nil {
			return result, err
		}

	}

	if user.Email != nil {
		result, err := s.repo.FindWithEmail(*user.Email)
		if err != nil {
			return result, err
		}

	}

	return result, nil
}
