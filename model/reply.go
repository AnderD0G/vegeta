package model

import (
	"gorm.io/gorm"
	"time"
	db2 "vegeta/db"
)

type Reply struct {
	Universal
	ReplyTo   int    `json:"-" gorm:"column:reply_to"`                         // 回复user_id
	ReplyUser User   `gorm:"foreignKey:Id;references:ReplyTo" json:"reply_to"` //关联User
	CommentID string `json:"comment_id" `
	User      User   `gorm:"foreignKey:Id;references:UserID" json:"user"` //关联User
}

func (m *Reply) TableName() string {
	return "reply"
}

func (m *Reply) Insert() error {

	db := db2.GetMysql("1")

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Debug().Model(m).Create(map[string]interface{}{
		"user_id":      m.User.Id,
		"comment_id":   m.CommentID,
		"content":      m.Content,
		"publish_time": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Model(new(Comment)).Where("id = ?", m.CommentID).Update("reply_counts", gorm.Expr("reply_counts  + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//type Reply struct {
//	CommentID int    `json:"comment_id" gorm:"column:comment_id"` // 评论id
//	ID        int    `json:"id" gorm:"column:id"`
//	UserID    string `json:"user_id" gorm:"column:user_id"`
//	Like      int    `json:"like" gorm:"column:like"`
//	Unlike    int    `json:"unlike" gorm:"column:unlike"`
//	Content   string `json:"content" gorm:"column:content"`
//	ReplyTo   string `json:"reply_to" gorm:"column:reply_to"` // 回复user_id
//}
//
//func (m *Reply) TableName() string {
//	return "reply"
//}
