package provider

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"vegeta/model"
	"vegeta/pkg"
)

// JWTGenerator jwt 生成器
type JWTGenerator[token jwt.Claims] interface {
	generate(ctx *gin.Context) (token, error)
	save(token) error
	get(token) (string, error)
}

// WxTokenGen WxJWT 根据code 生成jwt
type WxTokenGen struct {
	Query *pkg.Query
	I     *pkg.Inquirer[*model.User]
}

type WxToken struct {
	OpenID string `json:"openID"`
	jwt.StandardClaims
}

func (w WxTokenGen) generate(c *gin.Context) (WxToken, error) {
	code := c.Query("code")
	nowTime := time.Now()                    //当前时间
	expireTime := nowTime.Add(3 * time.Hour) //有效时间

	if session, err := pkg.Code2Session(code); err != nil {
		return WxToken{}, err
	} else {
		return WxToken{OpenID: session.Openid, StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "its me"}}, nil
	}
}

func (w WxTokenGen) save(x WxToken) error {
	u := model.User{}
	k := func(db *gorm.DB) {
		db.Debug().Where(model.User{Openid: x.OpenID}).FirstOrCreate(&u)
	}

	w.I.Query(new(model.User).TableName(), nil, k)
	return nil
}

func (w WxTokenGen) get(x WxToken) (string, error) {
	//TODO implement me
	panic("implement me")
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
