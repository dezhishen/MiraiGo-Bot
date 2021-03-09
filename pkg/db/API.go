package db

import "github.com/dezhiShen/MiraiGo-Bot/pkg/entity"

// GetAPI GetAPI
func GetAPI(ID string) *entity.API {
	if ID == ".hitokoto" {
		return &entity.API{
			ID:               ".hitokoto",
			Method:           entity.HTTPGet,
			URI:              "https://v1.hitokoto.cn/",
			ResponseTemplate: "${hitokoto}",
		}
	}
	if ID == ".fart" {
		return &entity.API{
			ID:               ".fart",
			Method:           entity.HTTPGet,
			URI:              "https://api.uixsj.cn/fart/get",
			ResponseTemplate: "ALL",
		}
	}
	if ID == ".weather" {
		return &entity.API{
			ID:               ".weather",
			Method:           entity.HTTPGet,
			URI:              "https://wis.qq.com/weather/common?source=pc&weather_type=observe&province=${1}&city=${2}&county=${3}",
			ResponseTemplate: "${1}${2}${3}现在${data.observe.weather_short},温度是${data.observe.degree}摄氏度",
		}
	}
	return nil
}
