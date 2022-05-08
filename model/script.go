package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"vegeta/db"
)

type (
	intSlice []int
	strSlice []string
)

type (
	// ScriptBasic journey需要的元素
	ScriptBasic struct {
		Uuid           string `json:"uuid" gorm:"column:uuid" valid:"no_empty"`
		ScriptCoverUrl string `json:"script_cover_url" gorm:"column:script_cover_url"`
		ScriptName     string `json:"script_name" gorm:"column:script_name"`
	}
	// ScriptPic scriptList里面的元素
	ScriptPic struct {
		ScriptBasic
		ScriptIntro string   `json:"script_intro" gorm:"column:script_intro"`
		Tags        strSlice `json:"script_tags" gorm:"-"`
	}

	Script struct {
		ScriptPic
		ScriptTag             intSlice `json:"-" gorm:"column:script_tag" type:"JsonArray"`
		ScriptScore           float64  `json:"script_score" gorm:"column:script_score"`
		GroupDuration         int      `json:"group_duration" gorm:"column:group_duration"`
		ScriptTextContext     string   `json:"script_text_context" gorm:"column:script_text_context"`
		ScriptPlotScore       float64  `json:"script_plot_score" gorm:"column:script_plot_score"`
		ScriptImageContent    strSlice `json:"script_image_content" gorm:"column:script_image_content"`
		ScriptMalePlayer      int      `json:"script_male_player" gorm:"column:script_male_player"`
		ScriptFemalePlayer    int      `json:"script_female_player" gorm:"column:script_female_player"`
		ScriptDifficultDegree string   `json:"script_difficult_degree" gorm:"column:script_difficult_degree"`
		ScriptPlayerLimit     int      `json:"script_player_limit" gorm:"column:script_player_limit"`
		Uuid                  string   `json:"uuid" gorm:"column:uuid" valid:"no_empty"`
		ScriptComplexScore    float64  `json:"script_complex_score" gorm:"column:script_complex_score"`
		Qid                   string   `json:"qid" gorm:"column:qid"`
		IsDel                 int      `json:"-" gorm:"column:is_del"`
	}
)

func (m Script) Index() string {
	return "scripts"
}

func (m *Script) TableName() string {
	return "scripts"
}

func (p intSlice) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *intSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &p)
}

func (p strSlice) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *strSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &p)
}

func (m *Script) Search() string {
	return "scripts"
}

// GetMultiMatch esMultiMatch
func GetMultiMatch(index, keyword string, page int, fields []string) ([]Script, error) {
	size := 20
	query := elastic.NewMultiMatchQuery(keyword, fields...)
	resp, err := db.GetES().Search().Index(index).Query(query).From(page).Size(size).Do(context.TODO())
	if err != nil {
		return nil, err
	}
	scripts := make([]Script, 0)
	script := Script{}
	for _, v := range resp.Each(reflect.TypeOf(script)) {
		s := v.(Script)
		scripts = append(scripts, s)
	}
	return scripts, nil
}

func ScriptPub(from *[]Script) (error, interface{}) {

	pubs := make([]ScriptPic, 0)
	err := copier.Copy(&pubs, from)
	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError)), nil
	}
	return nil, pubs
}
