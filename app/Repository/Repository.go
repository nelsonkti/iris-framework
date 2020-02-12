package Repository

import (
	"IrisFramework/config"
)

// NewUser get a user Repository
func NewUser() UserRepository {
	return UserRepository{
		db: config.DB,
	}
}