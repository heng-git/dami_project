//func main() {
//	// 创建OSSClient实例。
//	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
//	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
//	client, err := oss.New("oss-cn-hongkong.aliyuncs.com", "LTAI5t95jCFtkC5PV4AGXFNV", "6Ki43tzFXhA7AZU5U5wa9ZjyAF18eH")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//
//	// 填写存储空间名称，例如examplebucket。
//	bucket, err := client.Bucket("xiaomi-go2")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//
//	// 依次填写要上传到云存储空间的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
//	err = bucket.PutObjectFromFile("手机2号.jpg", "C:\\Users\\86158\\Desktop\\手机.jpg")
//	if err != nil {
//		fmt.Println("Error:", err)
//		os.Exit(-1)
//	}
//}

package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
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

	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	fd, err := os.Open("C:\\Users\\86158\\Desktop\\test.py")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 将文件流上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject("test.py", fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
