package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"xiaomi_project/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	pbLogin "xiaomi_project/proto/rbacLogin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	//fmt.Println(models.MD5("123456"))
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}
func (con LoginController) DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue") //通过标签的“name"属性获取传入的验证码的id和value值
	//1.验证验证码是否正确
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag { //如果验证码正确
		////2、查询数据库 判断用户以及密码是否存在
		//userinfoList := []models.Manager{}
		//password = models.MD5(password) //数据库中都是加密过的  因此对于用户上传的密码 加密完后才能进行比较
		//models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		//从rbac consul服务器中创建一个新的RbacLoginService
		rbacClient := pbLogin.NewRbacLoginService("rbac", models.RbacClient)
		res, _ := rbacClient.Login(context.Background(), &pbLogin.LoginRequest{
			Username: username,
			Password: models.MD5(password),
		})
		if res.IsLogin { //如果用户存在   保存用户信息并执行登录
			session := sessions.Default(c)
			//session.Set无法保存结构体对应的切片 因此要先把结构体转换成json字符串
			userinfoslice, _ := json.Marshal(res.Userlist) //返回的是byte  要进行类型转换
			session.Set("userinfo", string(userinfoslice)) //在session中创建userinfo信息
			session.Save()
			con.Success(c, "登录成功", "/admin/")
		} else {
			con.Error(c, "用户名或密码错误", "/admin/login")
		}

	} else {
		con.Error(c, "验证码验证失败", "/admin/login")
	}

}
func (con LoginController) Captcha(c *gin.Context) { //获取验证码
	id, b64s, err := models.MakeCaptcha(40, 100, 2)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("id", id)
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "退出登录成功", "/admin/login")
}

//func (con LoginController) Index(c *gin.Context) {
//	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
//
//}
//func (con LoginController) DoLogin(c *gin.Context) {
//	c.String(http.StatusOK, "-add--文章-")
//}
