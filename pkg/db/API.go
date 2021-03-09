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
			URI:              "https://api.openweathermap.org/data/2.5/weather?units=metric&lang=zh_cn&q=${1},chn&appid={API key}",
			ResponseTemplate: "${name}今天${main.temp}摄氏度,风速是${wind.speed}",
		}
	}
	return nil
}
