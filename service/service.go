package service

import (
	"github.com/golang-jwt/jwt/v4"
	"steamsale/entity"
	"time"
)

type service struct {
	repo Repository
}

// New returns 2 services to start Server.
func New(r Repository) (UserService, ItemService) {
	return &service{repo: r}, &service{repo: r}
}

type Repository interface {
	CreateUser(user entity.User) (entity.User, error)

	FindWithNickname(nickname string) (entity.User, error)
	FindWithEmail(email string) (entity.User, error)
	FindWithID(id int) (entity.User, error)

	RemoveUser(id int) error

	///////////////////////////////////////////

	AddItem(item entity.Item) (entity.Item, error)
	EditItem(changer entity.ItemChanger) (entity.Item, error)
	DeleteItem(id int) error

	FindItemWithID(id int) (entity.Item, error)
	FindItemWithName(name string) (entity.Item, error)
	FindItemWithRarity(rarity string) ([]entity.Item, error)
	FindItemWithPrice(hprice int, lprice int) ([]entity.Item, error)
}

func JWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	return "qq", nil
}
