package provider

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vegeta/db"
	"vegeta/model"
	"vegeta/pkg"
)

const openId = "openId"

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
	redis := db.GetRedis()

	var (
		resp   pkg.WxResp
		result string
		err    error
	)

	if resp, err = pkg.Code2Session(code); err != nil {
		return WxToken{}, err
	}

	if result, err = redis.HGet(context.TODO(), openId, resp.Openid).Result(); err != nil {
		return WxToken{}, err
	}

	if result == "" {
		_, err := redis.HSet(context.TODO(), openId, resp.Openid).Result()
		if err != nil {
			return WxToken{}, err
		}
	}

	return WxToken{OpenID: resp.Openid, StandardClaims: jwt.StandardClaims{}}, nil
}

func (w WxTokenGen) save(x WxToken) error {
	u := model.User{}
	k := func(db *gorm.DB) {
		db.Debug().Where(model.User{Openid: x.OpenID}).FirstOrCreate(&u)
	}

	w.I.Query(new(model.User).TableName(), nil, k)

	if _, err := db.GetRedis().HSet(context.TODO(), openId, u.Openid).Result(); err != nil {
		return err
	}

	return nil
}

func (w WxTokenGen) get(x WxToken) (string, error) {
	return pkg.GenerateToken(pkg.Claims{Code: x.OpenID, StandardClaims: jwt.StandardClaims{}})
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
