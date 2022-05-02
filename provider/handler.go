package provider

import (
	"github.com/gin-gonic/gin"
	ctxLogger "github.com/luizsuper/ctxLoggers"
	"go.uber.org/zap"
	"net/http"
)

const (
	Mysql = "mysql"
	Es    = "es"
)

type (
	Provider[MODEL any] interface {
		FindByID(context *gin.Context) (MODEL, error)
		// List @typ:指定查询的typ，目前支持mysql与es
		List(context *gin.Context, typ string) ([]MODEL, error)
		Update(id string, model MODEL) error
		Insert(model MODEL) error
		Delete(id string) error
	}

	HTTPHandler[MODEL any] struct {
		Provider   Provider[MODEL]
		ListStruct func(new *[]MODEL) (error, interface{})
		OneStruct  func(new *MODEL) (error, interface{})
	}
)

func (h *HTTPHandler[MODEL]) List(typ string) gin.HandlerFunc {

	return func(context *gin.Context) {
		if r, err := h.Provider.List(context, typ); err != nil {
			ctxLogger.Error(nil, "500", zap.String("err", err.Error()))
			context.JSON(http.StatusInternalServerError, nil)
			return
		} else {
			if h.ListStruct != nil {
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

func (h *HTTPHandler[MODEL]) FindByID() gin.HandlerFunc {
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
		}
	}

}
