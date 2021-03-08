package db

import "gopkg.in/fatih/set.v0"

// GetDraw 获取选项
func GetDraw(key string) set.Interface {
	if key == "标点符号" {
		result := set.New(set.ThreadSafe)
		result.Add("!")
		result.Add(",")
		result.Add("。")
		result.Add("...")
		return result
	}
	if key == "明日方舟" {
		result := set.New(set.ThreadSafe)
		result.Add("推进之王")
		result.Add("星熊")
		result.Add("伊芙利特")
		result.Add("德克萨斯")
		result.Add("芙兰卡")
		result.Add("普罗旺斯")
		result.Add("雷蛇")
		result.Add("深海色")
		result.Add("能天使")
		result.Add("闪灵")
		result.Add("幽灵鲨")
		result.Add("蓝毒")
		result.Add("临光")
		result.Add("赫默")
		result.Add("阿米娅")
		result.Add("红")
		result.Add("凛冬")
		result.Add("流星")
		result.Add("蛇屠箱")
		result.Add("白面鸮")
		return result
	}
	return nil
}
