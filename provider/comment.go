package provider

import (
	"github.com/gin-gonic/gin"
	"vegeta/model"
	"vegeta/pkg"
)

type Comment struct {
	Query *pkg.Query
	I     *pkg.Inquirer
}

func (s *Comment) FindByID(c *gin.Context) (model.Comment, error) {

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	query := c.DefaultQuery("query", "")

	s.Query.Condition = query
	s.Query.Page = pkg.Ati(page)
	s.Query.Size = pkg.Ati(size)

	s.I.InjectParam(s.Query)
	s.I.ParseStruct()

	if err := s.I.ParseQuery(); err != nil {
		return model.Comment{}, err
	}

	comment := model.GetComment(s.I)
	return comment, nil
}

func (s *Comment) List(c *gin.Context, db, typ string) ([]model.Comment, error) {

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

	comments := model.GetComments(s.I)
	return comments, nil
}

func (t *Comment) Update(context *gin.Context, id string, model model.Comment) error {

	//TODO implement me
	panic("implement me")
}

func (t *Comment) Insert(c *gin.Context, model model.Comment) error {

	//reply := c.PostForm("reply")

	//TODO implement me
	panic("implement me")
}

func (t *Comment) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
