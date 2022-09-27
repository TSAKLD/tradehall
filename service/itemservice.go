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

	FindItem(item entity.ItemFinder) ([]entity.Item, error)
}

func (s *service) AddItem(item entity.Item) (entity.Item, error) {
	result, err := s.repo.FindItemWithName(item.Name)
	if err == nil {
		err = fmt.Errorf("item already exists")
		return result, err
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

func (s *service) FindItem(item entity.ItemFinder) ([]entity.Item, error) {
	var result []entity.Item

	if item.Name != nil {
		res, err := s.repo.FindItemWithName(*item.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
		return result, nil
	}

	if item.Rarity != nil {
		result, err := s.repo.FindItemWithRarity(*item.Rarity)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	if item.ID != nil {
		res, err := s.repo.FindItemWithID(*item.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
		return result, nil
	}

	if item.LowestPrice != nil && item.HighestPrice != nil {
		result, err := s.repo.FindItemWithPrice(*item.LowestPrice, *item.HighestPrice)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, nil
}
