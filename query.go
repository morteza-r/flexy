package flexy

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Query struct {
	QTable   string      `json:"table"`
	QDoc     interface{} `json:"doc"`
	QModel   interface{} `json:"-"`
	QId      string      `json:"id"`
	QWheres  []Where     `json:"where"`
	QOrder   Order       `json:"order"`
	QType    string      `json:"type"`
	QAddress string      `json:"-"`
}
type Where struct {
	Field    string      `json:"field"`
	Value    interface{} `json:"value"`
	Operator string      `json:"operator"`
}

type Order struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

func (q *Query) Table(table string) *Query {
	q.QTable = table

	return q
}

func (q *Query) Model(item interface{}) *Query {
	if reflect.Indirect(reflect.ValueOf(item)).Kind() == reflect.Struct {
		q.QDoc = item
	}
	q.QModel = item

	return q
}

func (q *Query) OrderString(field string) *Query {
	q.QOrder.Type = "string"
	q.QOrder.Field = field

	return q
}

func (q *Query) OrderNumber(field string) *Query {
	q.QOrder.Type = "float_64"
	q.QOrder.Field = field

	return q
}

func (q *Query) Where(field string, operator string, value interface{}) *Query {
	var where = Where{
		Field:    field,
		Value:    value,
		Operator: operator,
	}
	q.QWheres = append(q.QWheres, where)

	return q
}

func (q *Query) CallApi() (err error) {
	b, err := json.Marshal(q)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST", q.QAddress+"/api/v1/db/query",
		bytes.NewBuffer(b),
	)
	if err != nil {
		//TODO better err
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res RestResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	if res.Status == false {
		err = errors.New(res.Message)
		return
	}

	err = mapstructure.Decode(res.Data, &q.QModel)
	if err != nil {
		return
	}

	err = resp.Body.Close()
	if err != nil {
		return
	}

	return
}
