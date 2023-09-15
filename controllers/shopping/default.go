package shopping

import (
	"fmt"
	"xiaomi_project/models"

	"github.com/gin-gonic/gin"
	"time"
)

type DefaultController struct {
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	////设置cookie
	//models.Cookie.Set(c, "username", "李四")
	//var username string
	//models.Cookie.Get(c, "username", &username)

	timeStart := time.Now().UnixNano()
	//1、获取顶部导航
	topNavList := []models.Nav{}
	if hastopNavList := models.CacheDb.Get("topNavList", &topNavList); !hastopNavList {
		models.DB.Where("status=1 AND position=1").Find(&topNavList) //redis里面没有数据再去数据库查询
		models.CacheDb.Set("topNavList", topNavList, 60*60)
		fmt.Println("数据库里面读取数据")
	}

	//2、获取轮播图数据
	focusList := []models.Focus{}
	if hasFocusList := models.CacheDb.Get("focusList", &focusList); !hasFocusList {
		models.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		models.CacheDb.Set("focusList", focusList, 60*60)
	}

	//3、获取分类的数据 挪到了base.go里面

	//4、获取中间导航 挪到了base.go里面

	//手机
	phoneList := []models.Goods{}
	if hasPhoneList := models.CacheDb.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList = models.GetGoodsByCategory(1, "best", 8)
		models.CacheDb.Set("phoneList", phoneList, 60*60)
	}
	otherList := []models.Goods{}
	if hasOtherList := models.CacheDb.Get("otherList", &otherList); !hasOtherList {
		otherList = models.GetGoodsByCategory(9, "all", 1)
		models.CacheDb.Set("otherList", otherList, 60*60)
	}
	timeEnd := time.Now().UnixNano()
	fmt.Printf("执行时间:%vms", (timeEnd-timeStart)/1000000)
	con.Render(c, "shopping/index/index.html", gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
		"otherList": otherList,
	})

}
