package provider

import (
	"github.com/gin-gonic/gin"
	"vegeta/model"
	"vegeta/pkg"
)

type Journey struct {
	Query *pkg.Query
	I     *pkg.Inquirer
}

type Detail struct{}

func (t *Journey) FindByID(context *gin.Context) (model.JourneyDis, error) {

	panic("implement me")
	//return model.Script{Name: "luiz"}, nil
}

func (s *Journey) List(c *gin.Context, db, typ string) ([]model.JourneyDis, error) {
	//todo:set default
	long := c.DefaultQuery("long", "")
	lat := c.DefaultQuery("lat", "")

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
	//journey := model.GetJourney(long, lat)

	n := model.GetJourneyN(long, lat, s.I)
	return n, nil
}

func (t *Journey) Update(context *gin.Context, id string, model model.JourneyDis) error {
	//TODO implement me
	panic("implement me")
}

func (t *Journey) Insert(context *gin.Context, model model.JourneyDis) error {
	//TODO implement me
	panic("implement me")
}

func (t *Journey) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t *Detail) FindByID(context *gin.Context) (model.JourneyPerson, error) {
	id := context.DefaultQuery("id", "")
	m, err := model.GetJourneyDetailM(pkg.Ati(id))
	if err != nil {
		return model.JourneyPerson{}, err
	}
	return *m, nil
}

func (s *Detail) List(c *gin.Context, db, typ string) ([]model.JourneyPerson, error) {

	panic("implement me")
}

func (t *Detail) Update(context *gin.Context, id string, model model.JourneyPerson) error {
	//TODO implement me
	panic("implement me")
}

func (t *Detail) Insert(context *gin.Context, model model.JourneyPerson) error {
	//TODO implement me
	panic("implement me")
}

func (t *Detail) Delete(context *gin.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
