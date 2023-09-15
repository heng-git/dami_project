package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

// 生成token的过程
// 1、首先需要自定义一个结构体，这个结构体需要继承 jwt.StandardClaims 结构体
type MyClaims struct {
	Uid int //自定义的属性 用于不同接口传值
	jwt.StandardClaims
}

// 2、定义key 和过期时间
var jwtKey = []byte("shopping.comxxx") //byte类型的切片
var expireTime = time.Now().Add(24 * time.Hour).Unix()

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "Api接口首页")
}

// 为了方便演示 这里使用Get请求
func (con ApiController) Login(c *gin.Context) {
	//3、实例化 存储token的结构体
	myClaimsObj := MyClaims{
		23,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "shopping",
		},
	}

	// 使用指定的签名方法创建签名对象
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaimsObj)
	// 使用指定的 jwtkey 签名并获得完整的编码后的字符串 token
	tokenStr, err := tokenObj.SignedString(jwtKey)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "生成token失败重试",
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取token成功",
		"token":   tokenStr,
		"success": true,
	})

}

func (con ApiController) AddressList(c *gin.Context) {
	//获取token  需要在请求头的Authorization字段传入
	tokenData := c.Request.Header.Get("Authorization")
	if len(tokenData) > 0 {
		tokenStr := strings.Split(tokenData, " ")[1]

		fmt.Println(tokenStr)
		token, myClaims, err := ParseToken(tokenStr)

		if err != nil || !token.Valid {
			c.JSON(200, gin.H{
				"message": "token传入错误",
				"success": false,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "验证token成功",
				"Uid":     myClaims.Uid,
				"success": true,
			})
		}
	}
}

type Address struct {
	Username string `form:"username" json:"username"`
	Address  string `form:"address" json:"address"`
	Tel      string `form:"tel" json:"tel"`
}

func (con ApiController) AddAddress(c *gin.Context) {
	//获取数据 post
	var address Address
	b, _ := c.GetRawData()
	json.Unmarshal(b, &address)
	fmt.Println(address)

	//获取token
	tokenData := c.Request.Header.Get("Authorization")
	if len(tokenData) > 0 {
		tokenStr := strings.Split(tokenData, " ")[1] //获取token

		fmt.Println(tokenStr)
		token, myClaims, err := ParseToken(tokenStr)

		if err != nil || !token.Valid {
			c.JSON(200, gin.H{
				"message": "token传入错误",
				"success": false,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "验证token成功",
				"Uid":     myClaims.Uid,
				"success": true,
			})
		}
	}
}

// 验证token是否合法
func ParseToken(tokenStr string) (*jwt.Token, *MyClaims, error) {
	myClaims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, myClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, myClaims, err
}
