package pkg

import (
	"github.com/dgrijalva/jwt-go"
)

//todo:密钥配置化，或者从redis中取出来
var jwtSecret interface{} = []byte("luiz999")

type WxToken struct {
	OpenID string `json:"openID"`
	jwt.StandardClaims
}

func GenerateToken(openId string) (string, error) {

	claims := WxToken{
		OpenID: openId,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*WxToken, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &WxToken{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*WxToken); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
