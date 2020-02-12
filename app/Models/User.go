package Models

import "time"

// User 用户实体，对应表user
type User struct {
	ID        int64 `xorm:"bigint(20) unsigned NOT NULL pk autoincr id"`
	Username  string `xorm:"not null default '' VARCHAR(32)   comment('姓名')"`
	Passwd    string `xorm:"not null  default '' VARCHAR(64)  comment('密码')"`
	Gender    string `xorm:"enum('男','女') NOT NULL DEFAULT '男' comment('性别') "`
	Age       int64 `xorm:"tinyint(4) NOT NULL DEFAULT 0  comment('年龄')"`
	CreatedAt time.Time `xorm:"timestamp created_at"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt time.Time `xorm:"timestamp updated_at"` // 这个Field将在Insert或Update时自动赋值为当前时间
	DeletedAt time.Time `xorm:"timestamp deleted_at"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}

func (DB *DB) GetInfoByUserName(username string) (User, error)  {
	var where = User{
		Username:  username,
	}
	userInfo, err := DB.db.Get(&where)

	if err != nil {
		return User{}, err
	}
	if !userInfo {
		return User{}, nil
	}
	return where, nil
}