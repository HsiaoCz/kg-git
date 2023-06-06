package model

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:id;type:int(11);" json:"uid"` // 用户id
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`
	Desc     string `gorm:"column:desc;type:varchar(255);" json:"desc"`
	Star     int    `gorm:"column:star;type:int(11);default:0;" json:"star"`
}

func (r RepoBasic) TableName() string {
	return "repo_basic"
}
