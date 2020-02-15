package persist

import (
	"context"
	"encoding/json"
	"function/crawler-concurrency-queue-refactoring/engine"
	"function/crawler-concurrency-queue-refactoring/zhenai/model"
	"github.com/olivere/elastic"
	"testing"
)

func TestItemSave(t *testing.T) {
	profile := model.Profile{
		Name:             "有缘会相识",
		Gender:           "女士",
		Age:              27,
		Height:           "170cm",
		Weight:           0,
		Income:           "12001-20000元",
		Marriage:         "未婚",
		Education:        "大学本科",
		Occupation:       "北京海淀区",
		HuKou:            "北京",
		BasicInfo:        []string{"未婚", "27岁", "天秤座(09.23-10.22)", "170cm", "工作地:北京海淀区", "月收入:1.2-2万", "舞蹈 大学本科"},
		DetailInfo:       []string{"汉族", "籍贯:浙江杭州", "体型:苗条", "不吸烟", "不喝酒", "何时结婚:时机成熟就结婚"},
		House:            "",
		Car:              "",
		IntroduceContent: "籍贯浙江杭州，却生长在山东，一副江南女子的皮囊，却生就了山东人的豪爽性格，从小喜爱舞蹈，12岁来到北京，大学学的古典舞，毕业后进入艺术团，又落户北京，因不安于现状，去年自己又开了舞蹈教育机构，目前正处于发展阶段……觅另一半大学本科以上，有事业心，对生活有所追求，身高180-190以内，月收入过2万，三观一致的男士，非诚勿扰",
	}
	expected := engine.Item{
		Url:     "https://album.zhenai.com/u/1111145895",
		Type:    "zhenai",
		Id:      "1111145895",
		Payload: profile,
	}

	// todo: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	// save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// fetch saved item
	do, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", do.Source)

	var actual engine.Item
	err = json.Unmarshal(do.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// verity result
	if actual.Url != expected.Url {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
