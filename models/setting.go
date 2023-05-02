package models

type Setting struct { //由于图片在form表中没有配置name，因此不能用shouldbind解析
	Id              int    `form:"id"`
	SiteTitle       string `form:"site_title"`
	SiteLogo        string
	SiteKeywords    string `form:"site_keywords"`
	SiteDescription string `form:"site_description"`
	NoPicture       string
	SiteIcp         string `form:"site_icp"`
	SiteTel         string `form:"site_tel"`
	SearchKeywords  string `form:"search_keywords"`
	TongjiCode      string `form:"tongji_code"`
	Appid           string `form:"appid"`
	AppSecret       string `form:"app_secret"`
	EndPoint        string `form:"end_point"`
	BucketName      string `form:"bucket_name"`
	OssStatus       int    `form:"oss_status"`
	OssDomain       string `form:"oss_domain"`
	ThumbnailSize   string `form:"thumbnail_size"`
}

func (Setting) TableName() string {
	return "setting"
}
