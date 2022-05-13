package model

import (
	"vegeta/pkg"
)

type (
	UserPub struct {
		Id        int    `json:"id" gorm:"column:id"`
		NickName  string `json:"nick_name" gorm:"column:nick_name"`   // 用户昵称
		AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"` // 用户头像图片的 URL。URL 最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640x640 的正方形头像，46 表示 46x46 的正方形头像，剩余数值以此类推。默认132），用户没有头像时该项为空。若用户更换头像，原有头像 URL 将失效。
		Gender    int    `json:"gender" gorm:"column:gender"`
	}
	User struct {
		UserPub
		Openid     string `json:"openid" gorm:"column:openid"`           // 用户唯一标识
		SessionKey string `json:"session_key" gorm:"column:session_key"` // 会话密钥
		TypeId     string `json:"type_id" gorm:"column:typeid"`          //  对应的某个小程序

	}
)

func (m *User) TableName() string {
	return "user"
}

func GetUsers(s *pkg.Inquirer[*User]) []User {
	i := make([]User, 0)
	s.Query(new(User).TableName(), &i)

	return i
}
