package Repository

import (
	"IrisFramework/app/Models"
	"xorm.io/xorm"
)

// Query表示访问者和操作查询。
type Query func(Models.User) bool

type UserRepository struct {
	db *xorm.Engine
}

func (userService *UserRepository) QueryByUsername(username string) (Models.User, error) {
	var user = Models.User{
		Username: username,
	}
	has, err := userService.db.Get(&user)
	if err != nil {
		return Models.User{}, err
	}
	if !has {
		return Models.User{}, nil
	}
	return user, nil
}