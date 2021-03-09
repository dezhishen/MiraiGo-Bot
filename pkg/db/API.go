package db

import "github.com/dezhiShen/MiraiGo-Bot/pkg/entity"

// GetAPI GetAPI
func GetAPI(ID string) *entity.API {
	if ID == ".hitokoto" {
		return &entity.API{
			ID:               ".hitokoto",
			Method:           "get",
			URI:              "https://v1.hitokoto.cn/",
			ResponseTemplate: "${hitokoto}",
		}
	}
	if ID == ".fart" {
		return &entity.API{
			ID:               ".fart",
			Method:           "get",
			URI:              "https://api.uixsj.cn/fart/get",
			ResponseTemplate: "ALL",
		}
	}
	return nil
}
