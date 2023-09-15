package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	. "github.com/hunterhug/go_image"
	"gopkg.in/ini.v1"
	"html/template"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 { //由于go执行速度过快 使用纳秒记录才能保存多张图片
	return time.Now().Unix()
}

// 获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
func MD5(str string) string {
	h := md5.New()         //创建一个md5实例
	io.WriteString(h, str) //编码成数字
	//io.WriteString(h, "And Leon's getting laaarger!")
	return fmt.Sprintf("%x", h.Sum(nil)) //完成加密并输出
}
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把string转换成Float64
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}
func String(n int) string {
	s := strconv.Itoa(n)
	return s
}

// 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	if GetOssStatus() == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}

	//return LocalUploadImg(c, picName)

}

// 上传图片到本地服务器里面
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件  对文件名进行处理  否则文件名可能重复
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}
	// 3、创建图片保存目录  static/upload/20210624
	day := GetDay()
	dir := "./static/upload/" + day
	//使用MkdirAll创建本地目录 0666表示权限：任何人都可以对这个目录进行读写
	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}
	// 4、生成文件名称和文件保存的目录(将文件以时间戳的形式命名)   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName
	// 5、将目录和文件名进行拼接执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst) //保存文件到dst目录中
	return dst, nil

}

// 上传图片到本地服务器里面
// 上传图片到Oss
func OssUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、定义图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "static/upload/" + day //注意不能随便更改dir地址  这是存放在云中的地址

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)

	OssUplod(file, dst)
	return dst, nil

}

// Oss上传
func OssUplod(file *multipart.FileHeader, dst string) (string, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI5t95jCFtkC5PV4AGXFNV", "6Ki43tzFXhA7AZU5U5wa9ZjyAF18eH")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("xiaomi-go2")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//首先将fileheader格式的文件转换为file格式
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	//将文件流上传至exampledir目录下的exampleobject.txt文件。
	//由于此处putObject的第二个文件参数是io.reader格式  与传入的fileheader格式的文件不符 需要先用file.open()转换成file才能传入
	// 上传文件流。
	err = bucket.PutObject(dst, f)
	if err != nil {
		return "", err
	}
	return dst, nil
}

// 通过列获取值
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	//fmt.Println(v)
	val := v.FieldByName(columnName).String() //不能直接用c.columnName获取  认不出来columnName这个属性
	return val
}

// 获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := Int(config.Section("oss").Key("status").String())
	return ossStatus
}
func FormatImg(str string) string {
	if GetOssStatus() == 1 {
		return GetSettingFromColumn("OssDomain") + str
	} else {
		return "/" + str
	}
}
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// 生成商品缩略图
func ResizeGoodsImage(filename string) {
	extname := path.Ext(filename)
	thumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",") //21，22,23,24
	thumbnailSizeSlice := strings.Split(thumbnailSize, ",")
	fmt.Println(thumbnailSizeSlice)
	//static/upload/tao_400.png
	//static/upload/tao_400.png_100x100.png
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		savepath := filename + "_" + thumbnailSizeSlice[i] + "x" + thumbnailSizeSlice[i] + extname
		w, _ := Int(thumbnailSizeSlice[i])
		err := ThumbnailF2F(filename, savepath, w, w)
		if err != nil {
			fmt.Println(err) //写个日志模块  处理日志
		}
	}

}
func Sub(a int, b int) int {
	return a - b
}

//生成随机数

func GetRandomNum() string {
	var str string

	for i := 0; i < 4; i++ {
		current := rand.Intn(10)

		str += String(current)
	}
	return str
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}

	return string(rs[start:end])

}

/*

str就是markdown语法

### 我是一个三级标题

<h3>我是一个三级标题</h3>


**我是一个加粗**

<strong>我是一个加粗</strong>


*/

func FormatAttr(str string) string {
	//fmt.Println(str)
	tempSlice := strings.Split(str, "&") //将字符串用&分开  按行进行解析
	var tempStr string
	for _, v := range tempSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		tempStr += string(output)
	}
	//fmt.Println(tempStr)
	return tempStr
}

func GetOrderId() string {
	// 2022020312233
	template := "20060102150405"
	return time.Now().Format(template) + GetRandomNum()
}
