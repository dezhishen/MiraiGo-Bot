package tools

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}
func intToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func boolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

// ToString toString
func ToString(arg interface{}) string {
	switch arg.(type) {
	case bool:
		return boolToString(arg.(bool))
	case float32:
		return floatToString(float64(arg.(float32)))
	case float64:
		return floatToString(arg.(float64))
		//case complex64:
		//  p.fmtComplex(complex128(f), 64, verb)
		//case complex128:
		//  p.fmtComplex(f, 128, verb)
	case int:
		return intToString(int64(arg.(int)))
	case int8:
		return intToString(int64(arg.(int8)))
	case int16:
		return intToString(int64(arg.(int16)))
	case int32:
		return intToString(int64(arg.(int32)))
	case int64:
		return intToString(int64(arg.(int64)))
	default:
		return fmt.Sprint(arg)
	}
}

func combinePath(pre string, path string) string {
	if pre != "" && path != "" {
		return pre + "." + path
	}
	return pre + path
}

// FlatMap 将一个map[string]interface打平
func FlatMap(prefix string, mapData map[string]interface{}) map[string]interface{} {
	v := reflect.ValueOf(mapData)
	res := make(map[string]interface{})
	foreachObj(prefix, v, res)
	return res
}

func foreachObj(pre string, v reflect.Value, res map[string]interface{}) {
	switch v.Kind() {
	case reflect.Ptr:
		foreachObj(pre, v.Elem(), res)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			foreachObj(combinePath(pre, strconv.Itoa(i)), v.Index(i), res)
		}
	case reflect.Struct:
		vType := v.Type()
		for i := 0; i < v.NumField(); i++ {
			foreachObj(combinePath(pre, vType.Field(i).Name), v.Field(i), res)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			foreachObj(combinePath(pre, key.String()), v.MapIndex(key), res)
		}
	case reflect.Interface:
		foreachObj(combinePath(pre, ""), v.Elem(), res)
	default: // float, complex, bool, chan, string,int,func, interface
		res[pre] = fmt.Sprintf("%v", v)
	}
}

func getTplExpressions(str string) []string {
	regStr := `\$\{\S*?\}`
	re, _ := regexp.Compile(regStr)
	all := re.FindAll([]byte(str), 1024)
	keyArrays := make([]string, 0)
	for _, item := range all {
		itemStr := string(item)
		if len(itemStr) > 3 {
			itemStr = itemStr[2 : len(itemStr)-1]
			keyArrays = append(keyArrays, itemStr)
		}

	}
	return keyArrays
}

// ParseTpl 将tpl中的占位符 替换为真实值 ${data.0.att1}
func ParseTpl(tpl string, data map[string]interface{}) string {
	if len(tpl) < 4 {
		return tpl
	}
	expressions := getTplExpressions(tpl)
	data = FlatMap("", data)
	for _, exp := range expressions {
		exp = strings.TrimSpace(exp)
		v, ok := data[exp]
		if !ok {
			v = ""
		}
		tpl = strings.Replace(tpl, "${"+exp+"}", ToString(v), -1)
	}
	return tpl
}
