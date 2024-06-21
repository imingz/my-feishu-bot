package cron

import (
	"reflect"
	"runtime"
	"strings"
)

// getFunctionName 使用 reflect 和 runtime 获取函数名称
func getFunctionName(i interface{}) string {
	// 使用 reflect.ValueOf 获取函数的值
	value := reflect.ValueOf(i)

	// 检查传入的是否为函数
	if value.Kind() != reflect.Func {
		return "not a function"
	}

	// 使用 runtime.FuncForPC 和 Value.Pointer 获取函数名称
	pc := value.Pointer()
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	str := strings.Split(fn.Name(), ".")

	return str[len(str)-1]
}
