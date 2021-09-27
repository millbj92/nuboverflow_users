//go:generate mockgen -destination=user_mocks_test.go -package=user github.com/brandomillerio/nuboverflow_users/internal/user Store

package user

import "context"

// Data model of a User.
// type User struct {
// 	ID      string

// }

type User struct {
	ID string
    UserName string
    Email string
    Github string
    Linkedin string
    UserScore int
    Bio string
    Profession string
    WorkPlace string
    Interests []string
    UserAwards []Award
}

type AwardLevel int
const (
	Bronze = iota
	Silver
	Gold
)

type Award struct {
    AwardName string
    AwardDescription string
    AwardLevel AwardLevel
    AwardScore int
}

type Store interface {
	GetUserByID(id string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
}

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetUserByID(ctx context.Context, id string) (User, error) {
	user, err := s.Store.GetUserByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s Service) CreateUser(ctx context.Context, user User) (User, error) {
	user, err := s.Store.CreateUser(user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s Service) DeleteUser(id string) error {
	err := s.Store.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
