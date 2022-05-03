package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"vegeta/db"
	"vegeta/model"
	"vegeta/pkg"
	"vegeta/provider"
)

func init() {
	d := db.DB[*gorm.DB]{}

	d.Provider = &db.MysqlPro{Address: "root:Caoxinguan2022@tcp(rm-bp1r6329yn2fo0390.mysql.rds.aliyuncs.com:3306)/taihe"}

	d.Initial()

	db.SetMysql(&d)
}
func main() {
	tai := db.GetMysql("1")

	t := new(provider.HTTPHandler[model.Script])
	t.Provider = &provider.Scripts{QueryMap: new(pkg.Query), S: &pkg.Inquirer[*model.Script]{
		M:  new(model.Script),
		Db: tai,
	}}

	j := new(provider.HTTPHandler[model.JourneyDis])
	j.Provider = &provider.Journey{}

	d := new(provider.HTTPHandler[model.JourneyPerson])
	d.Provider = &provider.Detail{}

	c := new(provider.HTTPHandler[model.Comment])
	c.Provider = &provider.Comment{Query: new(pkg.Query), I: &pkg.Inquirer[*model.Comment]{
		M:  new(model.Comment),
		Db: tai,
	},
	}
	c.ListStruct = model.CommentsPub
	c.OneStruct = model.CDetailPub

	router := gin.Default()
	router.GET("/script", t.List(provider.Mysql))
	router.GET("/js", j.List(provider.Mysql))
	router.GET("/js/detail", d.FindByID())
	router.GET("/comment", c.List(provider.Mysql))
	router.GET("/comment/detail", c.FindByID())
	router.GET("/script/vague", t.List(provider.Es))
	log.Fatal(router.Run(":8081"))

}
