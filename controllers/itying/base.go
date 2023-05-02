package itying

import (
	"net/http"
	"strings"
	"xiaomi_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct{}

func (con BaseController) Render(c *gin.Context, tpl string, data map[string]interface{}) {

	//1、获取顶部导航
	topNavList := []models.Nav{}
	if hasTopNavList := models.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
		models.DB.Where("status=1 AND position=1").Find(&topNavList)
		models.CacheDb.Set("topNavList", topNavList, 60*60)
	}

	//2、获取分类的数据
	goodsCateList := []models.GoodsCate{}

	if hasGoodsCateList := models.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		//https://gorm.io/zh_CN/docs/preload.html
		models.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		models.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}

	//3、获取中间导航
	middleNavList := []models.Nav{}
	if hasMiddleNavList := models.CacheDb.Get("middleNavList", &middleNavList); !hasMiddleNavList {
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

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.HTML(http.StatusOK, tpl, renderData)

}
