package Models

import "xorm.io/xorm"

type DB struct {
	db *xorm.Engine
}