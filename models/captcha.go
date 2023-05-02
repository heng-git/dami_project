package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// var Store = base64Captcha.DefaultMemStore //获取存储位置
// 让redisstore实现base64captcha.store这个接口
// 创建store
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverChinese{ //设置验证码参数
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "sadadadadaccxzc", //随机的验证码里的字符串就来自于这些字符
		//Source: "生成的就是中文验证码这里面的文字是配置文字源的",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		}, //创建一个字符串验证码
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()         //根据结构体参数转换字体
	c := base64Captcha.NewCaptcha(driver, store) // 根据字体驱动和存储位置创建一个验证码实例
	id, b64s, err := c.Generate()                //c.generate底层调用了store.Set方法写入数据库
	return id, b64s, err
}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	//id即创建时得到的id，verifyvalue即客户端传来的value 服务器会根据id找到服务器保存的value，
	//验证服务器保存的value和客户端传来的value是否一样
	if store.Verify(id, VerifyValue, true) { //true代表验证成功是否删除store里面的数据
		return true
	} else {
		return false
	}
}
