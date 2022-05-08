package provider

import (
	"github.com/gin-gonic/gin"
	"vegeta/model"
	"vegeta/pkg"
)

type Journey struct{}

type Detail struct{}

func (t *Journey) FindByID(context *gin.Context) (model.JourneyDis, error) {

	panic("implement me")
	//return model.Script{Name: "luiz"}, nil
}

func (s *Journey) List(c *gin.Context, db, typ string) ([]model.JourneyDis, error) {
	//todo:set default
	long := c.DefaultQuery("long", "")
	lat := c.DefaultQuery("lat", "")

	journey := model.GetJourney(long, lat)

	return journey, nil
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
