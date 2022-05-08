package provider

import "C"
import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
	"vegeta/model"
	"vegeta/pkg"
)

type Scripts struct {
	QueryMap    *pkg.Query
	scriptModel *[]model.Script
	S           *pkg.Inquirer[*model.Script]
}

func (s *Scripts) List(c *gin.Context, db, typ string) ([]model.Script, error) {
	switch db {
	case Es:
		return s.esList(c)
	case Mysql:
		return s.msList(c)
	default:
		return nil, errors.New("UnExcepted type")
	}
}

//getListFromMysql
func (s *Scripts) esList(c *gin.Context) ([]model.Script, error) {
	keyword := c.DefaultQuery("key", "")
	from := c.DefaultQuery("page", "1")

	return model.GetMultiMatch(model.Script{}.Index(), keyword, pkg.Ati(from), []string{"script_name", "script_text_context"})
}

//getListFromMysql
func (s *Scripts) msList(c *gin.Context) ([]model.Script, error) {
	if c != nil {
		page := c.DefaultQuery("page", "1")
		size := c.DefaultQuery("size", "10")
		query := c.DefaultQuery("query", "")

		s.QueryMap.Condition = query
		s.QueryMap.Page = pkg.Ati(page)
		s.QueryMap.Size = pkg.Ati(size)
	}

	t := new(Tag)

	err := pkg.Run(10*time.Second, c, s, t)
	if err != nil {
		return nil, err
	}

	m := make(map[int]string)
	for _, v := range *t.tagModel {
		m[v.Uuid] = v.Value
	}

	scripts := make([]model.Script, len(*s.scriptModel))

	for k, v := range *s.scriptModel {
		tags := make([]string, len(v.ScriptTag))
		for k, v := range v.ScriptTag {
			if s, ok := m[v]; ok {
				tags[k] = s
			} else {
				return nil, errors.New(fmt.Sprintf("tag %v 没有对应value", k))
			}
		}
		v.Tags = tags
		scripts[k] = v
	}

	return scripts, err
}

func (s *Scripts) Work(ctx context.Context, finishChan chan<- pkg.Finish) {
	go pkg.Watcher(ctx, finishChan)
	i := new([]model.Script)

	s.S.InjectParam(s.QueryMap)
	s.S.ParseStruct()
	err := s.S.ParseQuery()
	if err != nil {
		pkg.SafeSend(finishChan, pkg.Finish{
			IsDone: false,
			Err:    err,
		})
	}

	s.S.Query(new(model.Script).TableName(), i)

	s.scriptModel = i
	pkg.SafeSend(finishChan, pkg.Finish{
		IsDone: true,
		Err:    nil,
	})
}

func (t *Scripts) FindByID(context *gin.Context) (model.Script, error) {
	panic("implement me")
	//return model.Script{Name: "luiz"}, nil
}

func (t *Scripts) Update(context *gin.Context, id string, model model.Script) error {
	//TODO implement me
	panic("implement me")
}

func (t *Scripts) Insert(context *gin.Context, model model.Script) error {
	//TODO implement me
	panic("implement me")
}

func (t *Scripts) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
