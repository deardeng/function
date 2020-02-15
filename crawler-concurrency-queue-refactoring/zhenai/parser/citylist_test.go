package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	list := ParseCityList(contents)

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(list.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but have %d", resultSize, len(list.Requests))
	}

	for i, url := range expectedUrls {
		if list.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but"+" was %s", i, url, list.Requests[i].Url)
		}
	}

}
