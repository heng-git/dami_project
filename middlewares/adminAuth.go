package middlewares

import (
	"encoding/json"
	"fmt"
	"xiaomi_project/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//进行权限判断 没有登录的用户 不能进入后台管理中心
	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	//strings.split将字符串拆分成了string数组 ？前面的元素是数组的第一个元素
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	//2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok { //如果传入的是字符串
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") { //如果传入的结构体数组为空即该用户不存在
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login") //如果当前页面不在登录界面  则跳转到登录界面
			}
		} else { //用户登录成功  权限判断
			//1.根据角色获取当前角色的权限列表 然后把权限id放在一个map类型的对象里面
			urlPath := strings.Replace(pathname, "/admin/", "", 1)               //用“”取代pathname里面的“/admin"  将结果赋值到urlPath中
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) { //当前角色不是超级管理员且访问的路径不在排除的访问路径里面才会进行权限判断
				accessList := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&accessList)
				roleAccess := make(map[int]bool)
				for i := 0; i < len(accessList); i++ {
					roleAccess[accessList[i].AccessId] = true
				}
				access := models.Access{}
				models.DB.Where("url=?", urlPath).Find(&access)

				if _, ok := roleAccess[access.Id]; ok == false { //如果不在用户的权限列表里面则返回没有权限
					c.String(200, "没有权限")
					c.Abort()
				}
			}

		}
	} else { //不是string类型代表无法获取用户信息 即没有用户登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}

// 排除权限判断的方法  即访问哪些路径的时候不用进行权限判断
func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	//fmt.Println(excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
