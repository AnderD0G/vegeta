package db

import (
	"github.com/olivere/elastic/v7"
	"log"
)

const (
	//todo:配置化
	secret = `Lht86052516`
	user   = `elastic`
	url    = `es-cn-i7m2ofsnx000da6gn.public.elasticsearch.aliyuncs.com:9200`
)

var es ESPro

func init() {
	es = ESPro{Info: struct {
		User    string
		Secret  string
		Address string
	}{User: user, Secret: secret, Address: url}}
	es.initial()
}

type ESPro struct {
	m    map[string]*elastic.Client
	Info struct {
		User    string
		Secret  string
		Address string
	}
}

func (e *ESPro) initial() {
	sniffOpt := elastic.SetSniff(false)

	options := []elastic.ClientOptionFunc{sniffOpt, elastic.SetURL("http://es-cn-i7m2ofsnx000da6gn.public.elasticsearch.aliyuncs.com:9200"), elastic.SetBasicAuth("elastic", "Lht86052516")}

	client, err := elastic.NewClient(options...)
	if err != nil {
		log.Fatal(err)
	}

	if e.m == nil {
		e.m = make(map[string]*elastic.Client)
		e.m["1"] = client
	}

}

func GetES() *elastic.Client {
	return es.m["1"]
}

type ES interface {
	Index() string
}
