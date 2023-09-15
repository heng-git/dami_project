package main

import (
	"context"
	"fmt"
	"xiaomi_project/models"

	"time"
)

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// 实现设置captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := models.RedisDb.Set(ctx, key, value, time.Minute*2).Err()

	return err
}

// 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := models.RedisDb.Get(ctx, key).Result() //在redis数据库里面查找key值对应的value
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := models.RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
