package sqlite3

import (
	"context"
	"fmt"
	"testing"

	"github.com/Breeze0806/go-etl/config"
	"github.com/Breeze0806/go-etl/element"
	"github.com/Breeze0806/go-etl/storage/database"
)

func testJSONFromString(s string) *config.JSON {
	json, err := config.NewJSONFromString(s)
	if err != nil {
		panic(err)
	}
	return json
}

type TableParam struct {
	*database.BaseParam
}

func NewTableParam() *TableParam {
	return &TableParam{
		BaseParam: database.NewBaseParam(NewTable(database.NewBaseTable("", "", "test")), nil),
	}
}

func (t *TableParam) Query(_ []element.Record) (string, error) {
	return "select * from test", nil
}

func (t *TableParam) Agrs(_ []element.Record) ([]interface{}, error) {
	return nil, nil
}

type FetchHandler struct {
}

func (f *FetchHandler) OnRecord(r element.Record) error {
	fmt.Println(r)
	return nil
}

func (f *FetchHandler) CreateRecord() (element.Record, error) {
	return element.NewDefaultRecord(), nil
}

func Test_Sqlite3(t *testing.T) {
	t.Log("strat")
	db, err := database.Open("sqlite3", testJSONFromString(`{"url":"E:\\projects\\sqlite3\\test.db"}`))
	if err != nil {
		t.Errorf("open fail. err: %v", err)
	}
	defer db.Close()
	err = db.FetchRecord(context.TODO(), NewTableParam(), &FetchHandler{})
	if err != nil {
		t.Errorf("fetchRecord fail. err: %v", err)
	}
}