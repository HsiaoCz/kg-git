package model

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	Rid  int `gorm:"column:rid;type:int(11);" json:"rid"`      //仓库的id
	Uid  int `gorm:"column:uid;type:int(11);" json:"uid"`      //用户id
	Type int `gorm:"column:type;type:tinyint(1);" json:"type"` //类型 {1:所有者，2:被授权者}
}

func (r RepoUser) TableName() string {
	return "repo_user"
}
