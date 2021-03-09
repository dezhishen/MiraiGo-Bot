package db

import "github.com/dezhiShen/MiraiGo-Bot/pkg/entity"

// GetResp 获取返回值类型
func GetResp(ID string) *entity.Resp {
	if ID == ".r" {
		return &entity.Resp{
			ID:       ".r",
			Type:     "content",
			Template: "[${nickName}]掷出了:${answer}",
		}
	}
	if ID == ".draw" {
		return &entity.Resp{
			ID:       ".draw",
			Type:     "content",
			Template: "[${nickName}]抽出[${answer}]",
		}
	}
	if ID == ".hitokoto" {
		return &entity.Resp{
			ID:       ".hitokoto",
			Type:     "content",
			Template: "For ${nickName}: ${answer}",
		}
	}
	if ID == ".weather" {
		return &entity.Resp{
			ID:       ".hitokoto",
			Type:     "content",
			Template: "${answer}",
		}
	}
	return nil
}
