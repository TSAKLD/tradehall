package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"steamsale/entity"
	"steamsale/service"
)

type Repository struct {
	db *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
		db: database,
	}
}

//USER FUNCTIONS

func (r *Repository) CreateUser(user entity.User) (entity.User, error) {
	q := "INSERT INTO users(nickname, email, age, status) values $1, $2, $3, $4 RETURNING id, nickname, email, age, status"

	err := r.db.QueryRow(q, user.Nickname, user.Nickname, user.Age, user.Status).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) FindWithNickname(nickname string) (entity.User, error) {
	var user entity.User

	q := "SELECT id, nickname, status FROM users WHERE nickname = $1"

	err := r.db.QueryRow(q, nickname).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("%w: user nickname %v", service.ErrNotFound, nickname)
		}

		return user, err
	}

	return user, nil
}

func (r *Repository) FindWithEmail(email string) (entity.User, error) {
	var user entity.User

	q := "SELECT id, nickname, status FROM users WHERE email = $1"

	err := r.db.QueryRow(q, email).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("%w: user email %v", service.ErrNotFound, email)
		}

		return user, err
	}

	return user, nil
}

func (r *Repository) FindWithID(id int) (entity.User, error) {
	var user entity.User

	q := "SELECT id, nickname, status FROM users WHERE id = $1"

	err := r.db.QueryRow(q, id).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("%w: user id %v", service.ErrNotFound, id)
		}

		return user, err
	}

	return user, nil
}

func (r *Repository) RemoveUser(id int) error {
	q := "DELETE * FROM users WHERE id = $1"

	_, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

//ITEM FUNCTIONS

func (r *Repository) AddItem(item entity.Item) (entity.Item, error) {
	q := "INSERT INTO items(name, rarity, hero, cost) values $1, $2, $3, $4 RETURNING id, rarity, hero, cost"

	err := r.db.QueryRow(q, item.Name, item.Rarity, item.ForHero, item.Cost).Scan(&item)
	if err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (r *Repository) FindItemWithID(id int) (entity.Item, error) {
	var item entity.Item

	q := "SELECT id, name, rarity, hero, cost FROM items WHERE id = $1"

	err := r.db.QueryRow(q, id).Scan(&item)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return item, fmt.Errorf("%w: item id: %v", service.ErrNotFound, id)
		}

		return entity.Item{}, err
	}

	return item, nil
}

func (r *Repository) FindItemWithName(name string) (entity.Item, error) {
	var item entity.Item

	q := "SELECT id, name, rarity, hero, cost FROM items WHERE name = $1"

	err := r.db.QueryRow(q, name).Scan(&item)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return item, fmt.Errorf("%w: item name: %v", service.ErrNotFound, name)
		}

		return entity.Item{}, err
	}

	return item, nil
}

func (r *Repository) FindItemWithRarity(rarity string) ([]entity.Item, error) {
	var item entity.Item
	var itemList []entity.Item

	q := "SELECT id, name, rarity, hero, cost FROM items WHERE rarity = $1"

	result, err := r.db.Query(q, rarity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: item rarity: %v", service.ErrNotFound, rarity)
		}

		return nil, err
	}
	defer result.Close()

	for result.Next() {
		itemList = append(itemList, item)
	}

	return itemList, nil
}

func (r *Repository) FindItemWithPrice(hprice int, lprice int) ([]entity.Item, error) {
	var item entity.Item
	var list []entity.Item

	q := "SELECT id, name, rarity, hero, cost FROM items WHERE cost > $1 AND cost < $2"

	rows, err := r.db.Query(q, lprice, hprice)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: item cost beetween %v and %v", service.ErrNotFound, lprice, hprice)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		list = append(list, item)
	}

	return list, nil
}

func (r *Repository) EditItem(changer entity.ItemChanger) (entity.Item, error) {
	return entity.Item{}, nil
}

func (r *Repository) DeleteItem(id int) error {
	q := "DELETE FROM items * WHERE id = $1"

	_, err := r.db.Exec(q, id)

	return err
}
