package provider

import (
	"github.com/gin-gonic/gin"
	"vegeta/model"
	"vegeta/pkg"
)

type Category struct {
	Query *pkg.Query
	I     *pkg.Inquirer
}

func (c Category) FindByID(context *gin.Context) (model.Category, error) {
	return model.Category{}, nil
}

func (s Category) List(c *gin.Context, db, typ string) ([]model.Category, error) {

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	query := c.DefaultQuery("query", "")

	s.Query.Condition = query
	s.Query.Page = pkg.Ati(page)
	s.Query.Size = pkg.Ati(size)

	s.I.InjectParam(s.Query)
	s.I.ParseStruct()

	if err := s.I.ParseQuery(); err != nil {
		return nil, err
	}

	category := model.GetCategory(s.I)
	return category, nil
}

func (c Category) Update(context *gin.Context, id string, model model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c Category) Insert(context *gin.Context, model model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c Category) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
