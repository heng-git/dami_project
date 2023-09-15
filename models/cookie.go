package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// 定义结构体  缓存结构体 私有
type ginCookie struct{}

// 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {

	bytes, _ := json.Marshal(value)
	//des加密
	desKey := []byte("shopping.c")          //注意：密钥key必须是8位
	encData, _ := DesEncrypt(bytes, desKey) //加密  最后得到的是一堆乱码数据
	c.SetCookie(key, string(encData), 3600*24*30, "/", c.Request.Host, false, true)
}

// 获取数据的方法A
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		//des解密
		desKey := []byte("shopping.c") //注意：key必须是8位 密钥相同才能进行解密
		decData, e := DesDecrypt([]byte(valueStr), desKey)
		if e != nil {
			return false
		} else {
			err2 := json.Unmarshal([]byte(decData), obj)
			return err2 == nil
		}

	}
	return false
}
func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	//maxAge：cookie 的最大有效期，单位是秒。如果设置为 -1，
	//表示该 cookie 会在浏览器关闭时过期；如果设置为 0，表示该 cookie 会立即过期。
	//path：cookie 的作用范围，指定可以访问该 cookie 的 URL 路径。
	//默认值是 "/"，表示访问网站任意页面都可以访问该 cookie。
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

// 实例化结构体
var Cookie = &ginCookie{}
