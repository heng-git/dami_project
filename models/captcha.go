package models

import (
	"context"
	pbCaptcha "xiaomi_project/proto/captcha"

	"github.com/mojocn/base64Captcha"
	"go-micro.dev/v4/util/log"
)

// var Store = base64Captcha.DefaultMemStore //获取存储位置
// 让redisstore实现base64captcha.store这个接口
// 创建store
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha(height int, width int, length int) (string, string, error) {
	// Create client
	captchaClient := pbCaptcha.NewCaptchaService("captcha", CaptchaClient)
	// Call service
	res, err := captchaClient.MakeCaptcha(context.Background(), &pbCaptcha.MakeCaptchaRequest{
		Height: int32(height),
		Width:  int32(width),
		Length: int32(length),
	})
	if err != nil {
		log.Fatal(err)
	}

	return res.Id, res.B64S, err
}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	//id即创建时得到的id，verifyvalue即客户端传来的value 服务器会根据id找到服务器保存的value，
	//验证服务器保存的value和客户端传来的value是否一样
	// Create client
	captchaClient := pbCaptcha.NewCaptchaService("captcha", CaptchaClient)
	// Call service
	res, err := captchaClient.VerifyCaptcha(context.Background(), &pbCaptcha.VerifyCaptchaRequest{
		Id:          id,
		VerifyValue: VerifyValue,
	})
	if err != nil {
		log.Fatal(err)
	}
	return res.VerifyResult
}
