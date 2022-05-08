package provider

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	ctxLogger "github.com/luizsuper/ctxLoggers"
	"go.uber.org/zap"
	"net/http"
	"vegeta/pkg"
)

const (
	Mysql   = "mysql"
	Es      = "es"
	Normal  = "normal"
	DetailC = "detail"
)

type (
	Provider[MODEL any] interface {
		FindByID(context *gin.Context) (MODEL, error)
		// List @typ:指定查询的typ，目前支持mysql与es
		List(context *gin.Context, db, typ string) ([]MODEL, error)
		Update(context *gin.Context, id string, model MODEL) error
		Insert(context *gin.Context, model MODEL) error
		Delete(context *gin.Context, id string) error
	}

	APIHandler[MODEL any] struct {
		Provider   Provider[MODEL]
		ListStruct func(new *[]MODEL) (error, interface{})
		OneStruct  func(new *MODEL) (error, interface{})
	}
)

func (h *APIHandler[MODEL]) List(db, typ string) gin.HandlerFunc {

	return func(context *gin.Context) {
		if r, err := h.Provider.List(context, db, typ); err != nil {
			ctxLogger.Error(nil, "500", zap.String("err", err.Error()))
			context.JSON(http.StatusInternalServerError, nil)
			return
		} else {
			//仅当为normal的时候才返回详情
			if h.ListStruct != nil && typ == pkg.Normal {
				err, i := h.ListStruct(&r)
				if err != nil {
					context.JSON(http.StatusInternalServerError, nil)
					return
				}
				context.JSON(http.StatusOK, i)
				return
			}
			context.JSON(http.StatusOK, r)
			return
		}
	}
}

func (h *APIHandler[MODEL]) FindByID() gin.HandlerFunc {
	return func(context *gin.Context) {
		if r, err := h.Provider.FindByID(context); err != nil {
			ctxLogger.Error(nil, "500", zap.String("err", err.Error()))
			context.JSON(http.StatusInternalServerError, nil)
			return
		} else {
			if h.OneStruct != nil {
				err, i := h.OneStruct(&r)
				if err != nil {
					context.JSON(http.StatusInternalServerError, nil)
					return
				}
				context.JSON(http.StatusOK, i)
				return
			}
			context.JSON(http.StatusOK, r)
			return
		}
	}

}

type (
	LoginHandler[token jwt.Claims] struct {
		JWTGenerator[token]
	}
)

var (
	validate = validator.New()
)

// WxMiniLogin 微信小程序登录
func (h *LoginHandler[token]) WxMiniLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		code := c.Query("code")

		if err := validate.Var(code, "required"); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var (
			token     token
			jwtString string
			err       error
		)

		//first step 根据openId 生成jwt.claims
		if token, err = h.JWTGenerator.generate(c); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//second step 根据jwt.claims 生成jwtstring
		if jwtString, err = h.JWTGenerator.resp(token); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Header("token", jwtString)
		c.Status(http.StatusOK)
		return

	}

}

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
