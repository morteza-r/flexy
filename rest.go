package flexy

import (
	"bytes"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

type RestResponse struct {
	Data   interface{} `json:"data"`
	Status bool        `json:"status"`
}

func CallApi(q *Query) (res RestResponse, err error) {
	b, err := json.Marshal(q)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", q.QAddress + "/api/v1/db/query", bytes.NewBuffer(b))
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

	err = json.Unmarshal(body, &res)
	if err != nil {
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
