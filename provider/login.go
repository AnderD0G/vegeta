package provider

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"vegeta/model"
	"vegeta/pkg"
)

const openId = "openId"

// JWTGenerator jwt 生成器
type JWTGenerator[token jwt.Claims] interface {
	generate(openId, name, url string) (token, error)
	resp(token) (string, error)
	register(code string) (error, string)
}

// WxTokenGen WxJWT 根据code 生成jwt
type WxTokenGen struct {
	Query *pkg.Query
	I     *pkg.Inquirer[*model.User]
}

func (w WxTokenGen) generate(openId, name, url string) (pkg.WxToken, error) {
	//如果结果为空就去数据库中生成
	//u := model.User{}
	//k := func(db *gorm.DB) {
	//	user := model.User{
	//		UserPub: model.UserPub{NickName: name, AvatarUrl: url},
	//		Openid:  openId,
	//	}
	//	db.Debug().Where(user).FirstOrCreate(&u)
	//}
	s := func(db *gorm.DB) {
		user := model.User{
			UserPub: model.UserPub{NickName: name, AvatarUrl: url},
			Openid:  openId,
		}
		db.Debug().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "openid"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"avatar_url": url, "nick_name": name, "openid": openId}),
		}).Create(&user)
	}
	w.I.Query(new(model.User).TableName(), nil, s)

	return pkg.WxToken{OpenID: openId, StandardClaims: jwt.StandardClaims{}}, nil
}

func (w WxTokenGen) resp(x pkg.WxToken) (string, error) {
	return pkg.GenerateToken(x.OpenID)
}

func (w WxTokenGen) register(code string) (err error, openId string) {
	if session, err := pkg.Code2Session(code); err != nil {
		return err, ""
	} else {
		return nil, session.Openid
	}

}

//func SendToken(code, now string) error {
//
//	currentYear := time.Now().Year()   //当前年
//	currentMonth := time.Now().Month() //当前月
//	currentDay := time.Now().Day()     //当前日
//	zero := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.Local).Unix()
//
//	if dv := int64(pkg.Ati(now)) - zero; dv < 0 {
//
//	}
//}
