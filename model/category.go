package model

import (
	"gorm.io/gorm"
	"vegeta/pkg"
)

type (
	Category struct {
		Value string `json:"value" gorm:"column:value"`
		Uuid  string `json:"uuid" gorm:"primary_key" valid:"no_empty"`
		IsDel int    `json:"-" gorm:"column:is_del"`
		Tags  *[]Tag `json:"tags,omitempty" gorm:"many2many:c_tag;foreignKey:Uuid;joinForeignKey:CategoryUid;References:Uuid;JoinReferences:TagUid" valid:"no_empty"`
	}
	CategoriesTagDto struct {
		TIds []string `json:"tag_ids" valid:"no_empty"`
		Cid  string   `json:"category_id" valid:"no_empty"`
	}
)

func (m *Category) TableName() string {
	return "category"
}

func GetCategory(s *pkg.Inquirer) []Category {
	i := make([]Category, 0)
	k := func(db *gorm.DB) {
		db.Debug().Preload("Tags").Find(&i)
	}
	s.Query(new(Category).TableName(), nil, k)
	return i
}
