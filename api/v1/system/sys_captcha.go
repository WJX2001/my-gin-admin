package system

import (
	"github.com/WJX2001/gin-vue-admin-server/utils/captcha"
	"github.com/mojocn/base64Captcha"
	"strconv"
)

var store base64Captcha.Store = captcha.NewCacheStore()

type BaseApi struct{}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	case int64:
		i = int(v)
	case string:
		// redis 后端 GET 返回字符串
		if n, err := strconv.Atoi(v); err == nil {
			i = n
		}
	default:
		i = 0
	}
	return
}
