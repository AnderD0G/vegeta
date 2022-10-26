package main

import (
	"github.com/gin-contrib/gzip"
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

	//外网
	//d.Provider = &db.MysqlPro{Address: "super:Caoxinguan2022@tcp(rm-bp1r6329yn2fo03902o.mysql.rds.aliyuncs.com:3306)/taihe"}

	//内网
	d.Provider = &db.MysqlPro{Address: "super:Caoxinguan2022@tcp(rm-bp1r6329yn2fo0390.mysql.rds.aliyuncs.com:3306)/taihe"}

	d.Initial()

	db.SetMysql(&d)
}
func main() {
	tai := db.GetMysql("1")

	t := new(provider.APIHandler[model.Script])
	t.Provider = &provider.Scripts{QueryMap: new(pkg.Query), S: &pkg.Inquirer{
		M:  new(model.Script),
		Db: tai,
	}}
	t.ListStruct = model.ScriptPub

	p := new(provider.APIHandler[model.Reply])
	p.Provider = &provider.Reply{
		Query: new(pkg.Query),
		I: &pkg.Inquirer{
			M:  new(model.User),
			Db: tai,
		},
	}

	j := new(provider.APIHandler[model.JourneyDis])
	j.Provider = &provider.Journey{
		Query: new(pkg.Query),
		I: &pkg.Inquirer{
			M:  new(model.Journey),
			Db: tai,
		},
	}

	d := new(provider.APIHandler[model.JourneyPerson])
	d.Provider = &provider.Detail{}

	c := new(provider.APIHandler[model.Comment])
	c.Provider = &provider.Comment{Query: new(pkg.Query), I: &pkg.Inquirer{
		M:  new(model.Comment),
		Db: tai,
	},
	}
	c.ListStruct = model.CommentsPub

	l := new(provider.LoginHandler[pkg.WxToken])
	l.JWTGenerator = provider.WxTokenGen{
		Query: new(pkg.Query), I: &pkg.Inquirer{
			M:  new(model.User),
			Db: tai,
		},
	}

	ca := new(provider.APIHandler[model.Category])
	ca.Provider = &provider.Category{
		Query: new(pkg.Query),
		I: &pkg.Inquirer{
			M:  new(model.Category),
			Db: tai,
		},
	}

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/script", t.List(provider.Mysql, provider.Normal))
	router.GET("/script/detail", t.List(provider.Mysql, provider.DetailC))
	router.GET("/js", j.List(provider.Mysql, provider.Normal))
	router.GET("/js/detail", d.FindByID())
	router.GET("/comment", c.List(provider.Mysql, provider.Normal))
	router.GET("/comment/detail", c.FindByID())
	router.GET("/script/vague", t.List(provider.Es, provider.Normal))
	router.GET("/login", l.WxMiniLogin())
	router.GET("/register", l.WxMiniRegister())
	router.GET("/category", ca.List(provider.Mysql, provider.Normal))
	router.POST("/reply", p.Insert(model.Reply{}))

	log.Fatal(router.Run(":8081"))

}
