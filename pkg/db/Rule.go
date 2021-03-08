package db

import "github.com/dezhiShen/MiraiGo-Bot/pkg/entity"

// GetRule 根据传入的值,获取规则
func GetRule(messageType string, id int64, command string) *entity.Rule {
	if command == ".r" {
		return &entity.Rule{
			Command:      ".r",
			ID:           ".r",
			Max:          100,
			Min:          0,
			Type:         "randomMath",
			RespTemplate: "[$nickName]掷出了:$answer",
		}
	}
	if command == ".draw" {
		return &entity.Rule{
			Command:      ".draw",
			ID:           ".draw",
			Max:          100,
			Min:          0,
			Type:         "randomItem",
			RespTemplate: "[$nickName]抽出[$answer]",
		}
	}
	return nil
}
