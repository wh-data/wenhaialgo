package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/wh-data/wenhaialgo/at_frame_demo/biz"
)

var (
	cases_file_path = "../cases/test.json"
)

var (
	HttpClient *http.Client
	Cases      *biz.CasesStruct
)

func TestMain(m *testing.M) {
	fmt.Println("test sadf")
	HttpClient = biz.InitHttpClient()
	c := &biz.CasesStruct{}
	err := c.Load(cases_file_path)
	if err != nil {
		fmt.Println(err)
	}
	Cases = c
	exitVal := m.Run()
	os.Exit(exitVal)

}
