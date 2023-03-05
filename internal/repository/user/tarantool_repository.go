package user

import (
	"context"
	"fmt"
	"osoc/internal/api/http/v1/request"
	"osoc/internal/entity"
	"osoc/pkg/tarantool"
)

type TarantoolRepository struct {
	connection *tarantool.Connection
}

func NewTarantool(c *tarantool.Connection) *TarantoolRepository {
	return &TarantoolRepository{
		connection: c,
	}
}
func (t *TarantoolRepository) GetFriends(ctx context.Context, userID int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) GetUser(ctx context.Context, id int) (entity.User, error) {
	args := []interface{}{id}
	resp, err := t.connection.Call("get_user", args)
	if err != nil {
		return entity.User{}, err
	}
	// Получаем возвращаемое значение процедуры
	result := resp.Data[0]

	// Выводим результат
	fmt.Println(result)

	return entity.User{}, nil
}

func (t *TarantoolRepository) MultiCreateUser(ctx context.Context, users []entity.SecureUser) error {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) CreateUser(ctx context.Context, user entity.SecureUser) error {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) SearchUsers(ctx context.Context, query *request.UserSearch) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) UpdateUser(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolRepository) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
