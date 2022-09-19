package service

import (
	"errors"
	"fmt"
	"steamsale/entity"
)

type ItemService interface {
	AddItem(item entity.Item) (entity.Item, error)
	EditItem(changer entity.ItemChanger) (entity.Item, error)
	DeleteItem(id int) error

	FindItemWithID(id int) (entity.Item, error)
	FindItemWithName(name string) (entity.Item, error)
	FindItemWithRarity(rarity string) ([]entity.Item, error)
	FindItemWithPrice(hprice int, lprice int) ([]entity.Item, error)
}

func (s *service) AddItem(item entity.Item) (entity.Item, error) {
	item, err := s.repo.FindItemWithName(item.Name)
	if err == nil {
		err = fmt.Errorf("item already exists")
		return item, err
	} else if !errors.Is(err, ErrNotFound) {
		return entity.Item{}, err
	}

	item, err = s.repo.AddItem(item)
	if err != nil {
		return entity.Item{}, err
	}

	return item, nil
}
func (s *service) EditItem(changer entity.ItemChanger) (entity.Item, error) {
	return entity.Item{}, nil
}
func (s *service) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

func (s *service) FindItemWithID(id int) (entity.Item, error) {
	return s.repo.FindItemWithID(id)
}
func (s *service) FindItemWithName(name string) (entity.Item, error) {
	return s.repo.FindItemWithName(name)

}
func (s *service) FindItemWithRarity(rarity string) ([]entity.Item, error) {
	return s.repo.FindItemWithRarity(rarity)

}
func (s *service) FindItemWithPrice(lprice int, hprice int) ([]entity.Item, error) {
	return s.repo.FindItemWithPrice(lprice, hprice)
}
