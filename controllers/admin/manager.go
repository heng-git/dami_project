package admin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xiaomi_project/models"
	pbManager "xiaomi_project/proto/rbacManager"
	pbRole "xiaomi_project/proto/rbacRole"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	//managerList := []models.Manager{}
	//models.DB.Preload("Role").Find(&managerList) //找到Role.id与manager.id相等的manager
	rbacClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	res, _ := rbacClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{})
	//fmt.Printf("%#v", managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": res.ManagerList,
	})

}
func (con ManagerController) Add(c *gin.Context) {
	//rolelist := []models.Role{}
	//models.DB.Find(&rolelist)
	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, _ := rbacClient.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": res.RoleList,
	})
}
func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	//用户名和密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}
	////判断管理是否存在
	//managerList := []models.Manager{}
	//models.DB.Where("username=?", username).Find(&managerList)
	//if len(managerList) > 0 {
	//	con.Error(c, "此管理员已存在", "/admin/manager/add")
	//	return
	//}
	////执行增加管理员
	//manager := models.Manager{
	//	Username: username,
	//	Password: models.MD5(password),
	//	Email:    email,
	//	Mobile:   mobile,
	//	RoleId:   roleId,
	//	Status:   1,
	//	AddTime:  int(models.GetUnix()),
	//}
	//err2 := models.DB.Create(&manager).Error //给数据库增加对应的数据
	//if err2 != nil {
	//	con.Error(c, "增加管理员失败", "/admin/manager/add")
	//	return
	//}
	//判断管理是否存在
	rbacClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	res, _ := rbacClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Username: username,
	})

	if len(res.ManagerList) > 0 {
		con.Error(c, "此管理员已存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	addResult, _ := rbacClient.ManagerAdd(context.Background(), &pbManager.ManagerAddRequest{
		Username: username,
		Password: models.MD5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   int64(roleId),
		Status:   1,
		AddTime:  int64(models.GetUnix()),
	})

	if !addResult.Success {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "增加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {

	//获取管理员
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	//manager := models.Manager{Id: id}
	//models.DB.Find(&manager)
	//
	////获取所有的角色
	//roleList := []models.Role{}
	//models.DB.Find(&roleList)
	//获取管理员
	managerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	resManager, _ := managerClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Id: int64(id),
	})

	//获取所有的角色
	roleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	resRole, _ := roleClient.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  resManager.ManagerList[0],
		"roleList": resRole.RoleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := models.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ") //username是只读
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(mobile) > 11 {
		con.Error(c, "mobile长度不合法", "/admin/manager/edit?id="+models.String(id))
		return
	}
	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			con.Error(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
			return
		}
		password = models.MD5(password)
	}
	////执行修改
	//manager := models.Manager{Id: id}
	//models.DB.Find(&manager)
	//manager.Username = username
	//manager.Email = email
	//manager.Mobile = mobile
	//manager.RoleId = roleId

	//err3 := models.DB.Save(&manager).Error
	//if err3 != nil {
	//	con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
	//	return
	//}
	managerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	editResult, _ := managerClient.ManagerEdit(context.Background(), &pbManager.ManagerEditRequest{
		Id:       int64(id),
		Username: username,
		Password: password,
		Email:    email,
		Mobile:   mobile,
		RoleId:   int64(roleId),
	})
	if !editResult.Success {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	//注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码

	con.Success(c, "修改数据成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
	} else {
		//manager := models.Manager{Id: id}
		//models.DB.Delete(&manager)
		managerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
		managerRes, _ := managerClient.ManagerDelete(context.Background(), &pbManager.ManagerDeleteRequest{
			Id: int64(id),
		})
		if managerRes.Success {
			con.Success(c, "删除数据成功", "/admin/manager")
			return
		}
		con.Success(c, "删除数据成功", "/admin/manager")
	}
}
