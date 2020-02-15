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

	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	if len(list.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests; but have %d", resultSize, len(list.Requests))
	}

	for i, url := range expectedUrls {
		if list.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but"+" was %s", i, url, list.Requests[i].Url)
		}
	}

	if len(list.Items) != resultSize {
		t.Errorf("result should have %d "+"items; but have %d", resultSize, len(list.Requests))
	}

	for i, city := range expectedCities {
		if list.Items[i].(string) != city {
			t.Errorf("expected city#%d: %s; but"+" was %s", i, city, list.Items[i].(string))
		}
	}
}
