package admin

import (
	"context"
	"fmt"
	"xiaomi_project/models"

	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
	pbRole "xiaomi_project/proto/rbacRole"
)

type RoleController struct { //直接用BaseController对RoleController做了一个初始化
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	//roleList := []models.Role{}
	//models.DB.Find(&roleList)
	rbacRole := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, _ := rbacRole.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": res.RoleList, //渲染index.html里面的roleList字段
	})

}
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) { //执行添加

	title := strings.Trim(c.PostForm("title"), " ") //用Trim去除提交的内容里的空格
	description := strings.Trim(c.PostForm("description"), " ")
	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}
	rbacRole := pbRole.NewRbacRoleService("rbac", models.RbacClient) //新建一个RbacRoleService微服务
	res, _ := rbacRole.RoleAdd(context.Background(), &pbRole.RoleAddRequest{ //调用微服务
		Title:       title,
		Description: description,
		AddTime:     models.GetUnix(),
		Status:      1,
	})

	if !res.Success {
		con.Error(c, "增加角色失败 请重试", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}

}
func (con RoleController) Edit(c *gin.Context) {

	id, err := models.Int(c.Query("id")) //获取查询到id的值
	fmt.Println("修改时的id", id)
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		rbacRole := pbRole.NewRbacRoleService("rbac", models.RbacClient)
		res, _ := rbacRole.RoleGet(context.Background(), &pbRole.RoleGetRequest{
			Id: int64(id),
		})

		fmt.Println("role", rbacRole)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": res.RoleList[0],
		})
	}

}
func (con RoleController) DoEdit(c *gin.Context) { //执行修改

	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/edit")
	}
	rbacRole := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, _ := rbacRole.RoleEdit(context.Background(), &pbRole.RoleEditRequest{
		Id:          int64(id),
		Title:       title,
		Description: description,
	})
	if !res.Success {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
	//role := models.Role{Id: id}
	//models.DB.Find(&role)
	//role.Title = title
	//role.Description = description
	//
	//err2 := models.DB.Save(&role).Error
	//if err2 != nil {
	//	con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	//} else {
	//	con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	//}

	//查询要修改的数据 然后 修改

	//c.String(http.StatusOK, "-执行修改")
}
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		//role := models.Role{Id: id}
		//models.DB.Delete(&role)
		rbacRole := pbRole.NewRbacRoleService("rbac", models.RbacClient)
		res, _ := rbacRole.RoleDelete(context.Background(), &pbRole.RoleDeleteRequest{
			Id: int64(id),
		})
		if res.Success {
			con.Success(c, "删除数据成功", "/admin/role")
		} else {
			con.Error(c, "删除数据失败", "/admin/role")
		}
	}

}

func (con RoleController) Auth(c *gin.Context) {
	//1、获取角色id
	roleId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	////2、获取所有的权限
	//accessList := []models.Access{}
	//models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)
	//
	////3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	//roleAccess := []models.RoleAccess{}
	//models.DB.Where("role_id=?", roleId).Find(&roleAccess)
	//roleAccessMap := make(map[int]int)
	//for _, v := range roleAccess {
	//	roleAccessMap[v.AccessId] = 1
	//}
	//
	////4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	//
	//for i := 0; i < len(accessList); i++ { //此处不能使用range遍历  会导致没法修改accessList
	//	if _, ok := roleAccessMap[accessList[i].Id]; ok { //遍历顶级权限数据
	//		accessList[i].Checked = true
	//	}
	//	for j := 0; j < len(accessList[i].AccessItem); j++ { //遍历次级权限数据
	//		if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
	//			accessList[i].AccessItem[j].Checked = true
	//		}
	//	}
	//}
	//调用微服务显示授权页面
	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, _ := rbacClient.RoleAuth(context.Background(), &pbRole.RoleAuthRequest{
		RoleId: int64(roleId),
	})

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": res.AccessList,
	})

}

func (con RoleController) DoAuth(c *gin.Context) {
	//获取角色id
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	////获取被选中的权限id  切片
	accessIds := c.PostFormArray("access_node[]") //使用了PostFormArray

	////删除当前角色对应的权限
	//roleAccess := models.RoleAccess{}
	//models.DB.Where("role_id=?", roleId).Delete(&roleAccess)
	//
	////增加当前角色对应的权限
	//for _, v := range accessIds {
	//	roleAccess.RoleId = roleId
	//	accessId, _ := models.Int(v)
	//	roleAccess.AccessId = accessId
	//	models.DB.Create(&roleAccess)
	//}
	////fmt.Println(roleId)
	////fmt.Println(accessIds)
	//
	//fmt.Println("/admin/role/auth?id=?" + models.String(roleId))
	//// c.String(200, "DoAuth")
	//// admin/role/auth?id=9
	//con.Success(c, "授权成功", "/admin/role/auth?id="+models.String(roleId))
	//调用微服务执行授权

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, _ := rbacClient.RoleDoAuth(context.Background(), &pbRole.RoleDoAuthRequest{
		RoleId:    int64(roleId),
		AccessIds: accessIds,
	})
	if res.Success {
		con.Success(c, "授权成功", "/admin/role/auth?id="+models.String(roleId))
	} else {
		con.Error(c, "授权失败", "/admin/role/auth?id="+models.String(roleId))
	}
}
