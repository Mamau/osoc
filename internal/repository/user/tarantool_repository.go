package user

import (
	"context"
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
	res, err := t.connection.Call("box.func.get_user", args)
	if err != nil {
		return entity.User{}, err
	}

	result := res.Tuples()[0]

	// тут конечно должны быть какие-то проверки
	return entity.User{
		ID:        int(result[0].(uint64)),
		FirstName: result[1].(string),
		LastName:  result[2].(string),
		Age:       int(result[3].(uint64)),
		Sex:       result[4].(string),
		Interests: result[5].(string),
	}, nil
}

func (t *TarantoolRepository) MultiCreateUser(ctx context.Context, users []entity.SecureUser) error {
	for _, v := range users {
		if err := t.CreateUser(ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func (t *TarantoolRepository) CreateUser(ctx context.Context, user entity.SecureUser) error {
	args := []interface{}{user.FirstName, user.LastName, user.Age, user.Sex, user.Interests, user.Password}
	if _, err := t.connection.Call("box.func.create_user", args); err != nil {
		return err
	}

	return nil
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
