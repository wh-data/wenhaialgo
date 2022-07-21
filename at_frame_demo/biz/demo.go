package biz

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
	ID     int32 `json:"id"`
	Offset int32 `json:"offset"`
}

type Response struct {
	ErrorCode  int32 `json:"error_code"`
	TotalCount int32 `json:"total_count"`
}

type GetCase struct {
	CaseType string   `json:"type"`
	Name     string   `json:"name"`
	Req      Request  `json:"request"`
	Res      Response `json:"response"`
}

type UpdateCase struct {
	CaseType string   `json:"type"`
	Name     string   `json:"name"`
	Req      Request  `json:"request"`
	Res      Response `json:"response"`
}

type CasesStruct struct {
	GetCases []GetCase `json:"get_cases"`
}

func (c *CasesStruct) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

func InitHttpClient() (cli *http.Client) {
	return &http.Client{}
}

var (
	url    = "test"
	header = map[string][]string{
		"Host":          []string{"www.host.com"},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer Token"},
	}
)

func Getfunction(cli *http.Client, req Request) (res *Response, err error) {
	return &Response{
		ErrorCode:  12,
		TotalCount: 0,
	}, nil
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBufferString(string(payload)))
	//set header
	request.Header = header
	request.Header.Set("key", "value") // since each key may have multiple values, set will replace all old values, add will add to existing values
	response, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, res)
	return
}
