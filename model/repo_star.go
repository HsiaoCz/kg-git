package model

import "gorm.io/gorm"

type RepoStar struct {
	gorm.Model
	Rid int `gorm:"column:rid;type:int(11);" json:"rid"` //仓库的id
	Uid int `gorm:"column:uid;type:int(11);" json:"uid"` //用户id
}

func (rs RepoStar) TableName() string {
	return "repo_star"
}
