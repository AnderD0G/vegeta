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
	resp(token) (string, error)
}

// WxTokenGen WxJWT 根据code 生成jwt
type WxTokenGen struct {
	Query *pkg.Query
	I     *pkg.Inquirer[*model.User]
}

func (w WxTokenGen) generate(c *gin.Context) (pkg.WxToken, error) {
	code := c.Query("code")
	name := c.Query("name")
	url := c.Query("url")

	//redis := db.GetRedis()

	var (
		resp   pkg.WxResp
		result string
		err    error
		rdb    = db.GetRedis()
	)

	if resp, err = pkg.Code2Session(code); err != nil {
		return pkg.WxToken{}, err
	}

	if result, err = rdb.HGet(context.TODO(), openId, resp.Openid).Result(); err != nil {
		return pkg.WxToken{}, err
	}
	//如果结果为空就去数据库中生成
	if result == "" {

		u := model.User{}
		k := func(db *gorm.DB) {
			user := model.User{
				UserPub: model.UserPub{NickName: name, AvatarUrl: url},
				Openid:  resp.Openid,
			}
			db.Debug().Where(user).FirstOrCreate(&u)
		}

		w.I.Query(new(model.User).TableName(), nil, k)

	}

	return pkg.WxToken{OpenID: resp.Openid, StandardClaims: jwt.StandardClaims{}}, nil
}

func (w WxTokenGen) resp(x pkg.WxToken) (string, error) {
	return pkg.GenerateToken(x.OpenID)
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
