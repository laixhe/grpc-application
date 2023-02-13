package data

import (
	"context"
)

type User struct {
}

func (r *DB) SaveUser(ctx context.Context, u *User) error {
	return nil
}

func (r *DB) UpdateUser(ctx context.Context, u *User) error {
	return nil
}

func (r *DB) GetUser(ctx context.Context, id int64) (*User, error) {
	return nil, nil
}

func (r *DB) ListAllUser(ctx context.Context) ([]*User, error) {
	return nil, nil
}
