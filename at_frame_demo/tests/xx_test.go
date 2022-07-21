package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wh-data/wenhaialgo/at_frame_demo/biz"
)

func Test_get(t *testing.T) {
	if Cases == nil || Cases.GetCases == nil {
		return
	}
	fmt.Println("test: Cases.GetCases", Cases.GetCases)
	for _, c := range Cases.GetCases {
		fmt.Println(c.CaseType)
		fmt.Println(c.Name)
		res, err := biz.Getfunction(HttpClient, c.Req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		assert.Equal(t, c.Res.ErrorCode, res.ErrorCode)
	}

}
