package parser

import (
	"encoding/json"
	"function/crawler-concurrency-distributed/config"
	"function/crawler-concurrency-distributed/engine"
	"function/crawler-concurrency-distributed/zhenai/model"
	"log"
	"regexp"
)

var infoRe = regexp.MustCompile(`window.__INITIAL_STATE__=({.*});`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

type UserInfo struct {
	ObjectInfo        ObjectInfo   `json:"objectInfo"`
	Interest          []Interest   `json:"interest"`
	MemberList        []MemberList `json:"memberList"`
	SeoInfo           SeoInfo      `json:"seoInfo"`
	IsSearchEntryFlag bool         `json:"isSearchEntryFlag"`
}
type Photos struct {
	CreateTime  string `json:"createTime"`
	IsAvatar    bool   `json:"isAvatar"`
	PhotoID     int    `json:"photoID"`
	PhotoType   int    `json:"photoType"`
	PhotoURL    string `json:"photoURL"`
	PraiseCount int    `json:"praiseCount"`
	Praised     bool   `json:"praised"`
	Verified    bool   `json:"verified"`
}
type ObjectInfo struct {
	Age                      int      `json:"age"`
	AvatarPhotoID            int      `json:"avatarPhotoID"`
	AvatarPraiseCount        int      `json:"avatarPraiseCount"`
	AvatarPraised            bool     `json:"avatarPraised"`
	AvatarURL                string   `json:"avatarURL"`
	BasicInfo                []string `json:"basicInfo"`
	DetailInfo               []string `json:"detailInfo"`
	EducationString          string   `json:"educationString"`
	EmotionStatus            int      `json:"emotionStatus"`
	Gender                   int      `json:"gender"`
	GenderString             string   `json:"genderString"`
	HasIntroduce             bool     `json:"hasIntroduce"`
	HeightString             string   `json:"heightString"`
	HideVerifyModule         bool     `json:"hideVerifyModule"`
	IntroduceContent         string   `json:"introduceContent"`
	IntroducePraiseCount     int      `json:"introducePraiseCount"`
	IsActive                 bool     `json:"isActive"`
	IsFollowing              bool     `json:"isFollowing"`
	IsInBlackList            bool     `json:"isInBlackList"`
	IsStar                   bool     `json:"isStar"`
	IsZhenaiMail             bool     `json:"isZhenaiMail"`
	LastLoginTimeString      string   `json:"lastLoginTimeString"`
	LiveAudienceCount        int      `json:"liveAudienceCount"`
	LiveType                 int      `json:"liveType"`
	MarriageString           string   `json:"marriageString"`
	MemberID                 int      `json:"memberID"`
	MomentCount              int      `json:"momentCount"`
	Nickname                 string   `json:"nickname"`
	ObjectAgeString          string   `json:"objectAgeString"`
	ObjectBodyString         string   `json:"objectBodyString"`
	ObjectChildrenString     string   `json:"objectChildrenString"`
	ObjectEducationString    string   `json:"objectEducationString"`
	ObjectHeightString       string   `json:"objectHeightString"`
	ObjectInfo               []string `json:"objectInfo"`
	ObjectMarriageString     string   `json:"objectMarriageString"`
	ObjectSalaryString       string   `json:"objectSalaryString"`
	ObjectWantChildrenString string   `json:"objectWantChildrenString"`
	ObjectWorkCityString     string   `json:"objectWorkCityString"`
	Onlive                   int      `json:"onlive"`
	PhotoCount               int      `json:"photoCount"`
	Photos                   []Photos `json:"photos"`
	PraisedIntroduce         bool     `json:"praisedIntroduce"`
	PreviewPhotoURL          string   `json:"previewPhotoURL"`
	PycreditCertify          bool     `json:"pycreditCertify"`
	SalaryString             string   `json:"salaryString"`
	ShowValidateIDCardFlag   bool     `json:"showValidateIDCardFlag"`
	TotalPhotoCount          int      `json:"totalPhotoCount"`
	ValidateEducation        bool     `json:"validateEducation"`
	ValidateFace             bool     `json:"validateFace"`
	ValidateIDCard           bool     `json:"validateIDCard"`
	VideoCount               int      `json:"videoCount"`
	VideoID                  int      `json:"videoID"`
	WorkCity                 int      `json:"workCity"`
	WorkCityString           string   `json:"workCityString"`
	WorkProvinceCityString   string   `json:"workProvinceCityString"`
}
type AnswerContentDetailRecords struct {
	AnswerContentDetail string `json:"answerContentDetail"`
	Status              int    `json:"status"`
	TagID               int    `json:"tagId"`
}
type Interest struct {
	AnswerContent              string                       `json:"answerContent"`
	AnswerContentDetail        string                       `json:"answerContentDetail"`
	AnswerContentDetailStatus  int                          `json:"answerContentDetailStatus"`
	AnswerGuideWord            string                       `json:"answerGuideWord"`
	AnswerID                   int                          `json:"answerID"`
	AnswerOrder                int                          `json:"answerOrder"`
	AnswerRecordID             int                          `json:"answerRecordID"`
	AnswerWriteRule            int                          `json:"answerWriteRule"`
	IconURL                    string                       `json:"iconURL"`
	NewInterest                bool                         `json:"newInterest"`
	PraiseCount                int                          `json:"praiseCount"`
	QuestionGuideWord          string                       `json:"questionGuideWord"`
	QuestionID                 int                          `json:"questionID"`
	QuestionName               string                       `json:"questionName"`
	QuestionType               int                          `json:"questionType"`
	VerifyStatus               int                          `json:"verifyStatus"`
	AnswerContentDetailRecords []AnswerContentDetailRecords `json:"answerContentDetailRecords,omitempty"`
}
type FlagList struct {
	Status int `json:"status"`
	Type   int `json:"type"`
}
type NatureTags struct {
	TagColor int    `json:"tagColor"`
	TagDesc  string `json:"tagDesc"`
	TagFlag  int    `json:"tagFlag"`
	TagName  string `json:"tagName"`
	TagType  int    `json:"tagType"`
}
type MemberList struct {
	AdvantageMsgList       []interface{} `json:"advantageMsgList"`
	AdvantageNatureTagList []interface{} `json:"advantageNatureTagList"`
	AvatarURL              string        `json:"avatarURL"`
	EmotionStatus          int           `json:"emotionStatus"`
	FlagList               []FlagList    `json:"flagList"`
	Gender                 int           `json:"gender"`
	HasShortVideo          bool          `json:"hasShortVideo"`
	Height                 int           `json:"height"`
	IntroduceContent       string        `json:"introduceContent"`
	LastLoginTime          string        `json:"lastLoginTime"`
	LiveType               int           `json:"liveType"`
	NatureTags             []NatureTags  `json:"natureTags"`
	Nickname               string        `json:"nickname"`
	ObjectID               int           `json:"objectID"`
	Online                 int           `json:"online"`
	Onlive                 int           `json:"onlive"`
	ShowStarsFlag          bool          `json:"showStarsFlag"`
	TrueName               string        `json:"trueName"`
	UserAge                int           `json:"userAge"`
}
type Location struct {
	CurrLocation bool   `json:"currLocation"`
	LinkContent  string `json:"linkContent"`
	LinkURL      string `json:"linkURL"`
}
type NearbyCitys struct {
	LinkContent string `json:"linkContent"`
	LinkURL     string `json:"linkURL"`
}
type SeoInfo struct {
	Location    []Location    `json:"location"`
	NearbyCitys []NearbyCitys `json:"nearbyCitys"`
}

func parseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	ret := engine.ParseResult{
		Requests: nil,
		Items:    nil,
	}

	match := infoRe.FindSubmatch(contents)
	result := UserInfo{}
	if len(match) < 2 {
		log.Printf("find json string splice lowwer 2")
		return ret
	}
	err := json.Unmarshal(match[1], &result)
	if err != nil {
		log.Printf("userinfo unmarshal error : %v %v", err, match)
		return ret
	}

	profile.Income = result.ObjectInfo.SalaryString
	profile.Gender = result.ObjectInfo.GenderString
	profile.BasicInfo = result.ObjectInfo.BasicInfo
	profile.DetailInfo = result.ObjectInfo.DetailInfo
	profile.Marriage = result.ObjectInfo.MarriageString
	profile.Education = result.ObjectInfo.ObjectEducationString
	profile.Occupation = result.ObjectInfo.WorkProvinceCityString
	profile.HuKou = result.ObjectInfo.WorkCityString
	profile.IntroduceContent = result.ObjectInfo.IntroduceContent
	profile.Height = result.ObjectInfo.HeightString
	profile.Age = result.ObjectInfo.Age

	ret.Items = []engine.Item{
		{
			Url:     url,
			Type:    "zhenai",
			Id:      extractString([]byte(url), idUrlRe),
			Payload: profile,
		},
	}
	return ret
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParse struct {
	userName string
}

func (p *ProfileParse) Parser(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParse) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParse(name string) *ProfileParse {
	return &ProfileParse{
		userName: name,
	}
}
