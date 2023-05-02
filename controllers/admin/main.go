package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
	"net/http"
	"xiaomi_project/models"

	"github.com/gin-gonic/gin"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//1.获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		//2、获取所有的权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC") //对获取到的accessitem按照数据库access表的sort字段进行降序排序
		}).Order("sort DESC").Find(&accessList) //对获取到的accessslist中的数据也降序排序

		//3. 获取当前角色对应权限
		role_access := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&role_access)
		roleAccessMap := make(map[int]int)
		for _, v := range role_access {
			roleAccessMap[v.AccessId] = 1
		}

		for i := 0; i < len(accessList); i++ { //此处不能使用range遍历  会导致没法修改accessList
			if _, ok := roleAccessMap[accessList[i].Id]; ok { //遍历顶级权限数据
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ { //遍历次级权限数据
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改状态的方法
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// status = ABS(0-1)   1

	// status = ABS(1-1)  0
	//执行原生sql语句
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 公共修改状态的方法
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}

}

// 清除缓存
func (con MainController) FlushAll(c *gin.Context) {
	models.CacheDb.FlushAll()
	con.Success(c, "清除Redis缓存数据成功", "/admin")
}
