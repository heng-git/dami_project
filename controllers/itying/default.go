package itying

import (
	"fmt"
	"xiaomi_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
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

	//3、获取分类的数据
	goodsCateList := []models.GoodsCate{}
	if hasGoodsCateList := models.CacheDb.Get("goodsCateList", &topNavList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		//只有pid=0并且status等于1的goodscatelist才会被显示在网页上
		models.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC") //自定义预加载函数  只有goodscate为1的goodsCateItems才会被降序显示在网页上
		}).Find(&goodsCateList)
		models.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}

	// fmt.Println(focusList)
	//4、获取中间导航
	middleNavList := []models.Nav{}
	if hasmiddleNavList := models.CacheDb.Get("middleNavList", &middleNavList); !hasmiddleNavList {
		models.DB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") //21，22,23,24
			relationIds := strings.Split(relation, ",")
			goodsList := []models.Goods{}
			models.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}
		models.CacheDb.Set("middleNavList", middleNavList, 60*60)
	}

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
	c.HTML(http.StatusOK, "itying/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})

}
