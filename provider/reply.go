package provider

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"vegeta/model"
	"vegeta/pkg"
)

type Reply struct {
	Query *pkg.Query
	I     *pkg.Inquirer[*model.User]
}

func (r Reply) FindByID(context *gin.Context) (model.Reply, error) {
	panic("implement me")
}

func (r Reply) List(context *gin.Context, db, typ string) ([]model.Reply, error) {
	//TODO implement me
	panic("implement me")
}

func (r Reply) Update(context *gin.Context, id string, s model.Reply) error {
	user := model.Reply{}
	err := context.BindJSON(&user)
	if err != nil {
		return err
	}
	return nil
}

func (r Reply) Insert(context *gin.Context, rp model.Reply) error {

	if err := context.BindJSON(&rp); err != nil {
		return err
	}
	//todo:将验证移到中间层
	token := context.GetHeader("token")

	if parseToken, err := pkg.ParseToken(token); err != nil {
		return err
	} else {
		r.Query.Condition = fmt.Sprintf("(openid=%v)", parseToken.OpenID)
		r.I.InjectParam(r.Query)
		r.I.ParseStruct()
		if err = r.I.ParseQuery(); err != nil {
			return err
		}
	}

	if s := model.GetUsers(r.I); len(s) != 1 {
		return errors.New("no such user")
	} else {
		rp.User = s[0]
	}

	return rp.Insert()
}

func (r Reply) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
