package db

import (
	"github.com/olivere/elastic/v7"
	"log"
)

const (
	secret = `gpv0cIFM7ES2quWOht50KXMa`
	user   = `elastic`
	url    = `https://my-deployment-d794e9.es.us-central1.gcp.cloud.es.io:9243`
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

	options := []elastic.ClientOptionFunc{sniffOpt, elastic.SetURL(e.Info.Address), elastic.SetBasicAuth(e.Info.User, e.Info.Secret)}

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
