package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityUserList(t *testing.T) {
	content, err := ioutil.ReadFile("cityuserlist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityUserList(content)
	const resultSize = 20

	expectedUrls := []string{
		"http://album.zhenai.com/u/1921187162",
		"http://album.zhenai.com/u/1002520123",
		"http://album.zhenai.com/u/1133630423",
	}

	expectedUsers := []string{}

	//fmt.Println(result)

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}

	for i, user := range expectedUsers {
		if result.Items[i].(string) != user {
			t.Errorf("expected url #%d: %s; but was %s", i, user, result.Items[i].(string))
		}
	}
}
