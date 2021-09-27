package user

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get user by id", func(t *testing.T) {
		userStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		userStoreMock.
			EXPECT().
			GetUserByID(id).
			Return(User{
				ID: id,
			}, nil)

		userService := New(userStoreMock)
		user, err := userService.GetUserByID(
			context.Background(),
			id,
		)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", user.ID)
	})

	t.Run("tests insert user", func(t *testing.T) {
		userStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		userStoreMock.
			EXPECT().
			CreateUser(User{
				ID: id,
			}).
			Return(User{
				ID: id,
			}, nil)

		userService := New(userStoreMock)
		user, err := userService.CreateUser(
			context.Background(),
			User{
				ID: id,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", user.ID)
	})

	t.Run("tests delete user", func(t *testing.T) {
		userStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		userStoreMock.
			EXPECT().
			DeleteUser(id).
			Return(nil)

		userService := New(userStoreMock)
		err := userService.DeleteUser(context.Background(), id)
		assert.NoError(t, err)
	})
}
