package config

const (
	ParseProfile      = "ParseProfile"
	ParseCityList     = "ParseCityList"
	ParseCityUserList = "ParseCityUserList"
	NilParser         = "NilParser"

	ElasticIndex    = "dating_profile"
	ItemSaverRPC    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
	ItemSaverPort   = 1234

	// Rate limiting
	Qps = 20
)
