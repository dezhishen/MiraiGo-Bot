package db

import "github.com/dezhiShen/MiraiGo-Bot/pkg/entity"

// GetRule 根据传入的值,获取规则
func GetRule(messageType string, id int64, command string) *entity.Rule {
	if command == ".r" {
		return &entity.Rule{
			Command: ".r",
			ID:      ".r",
			Max:     100,
			Min:     0,
			Type:    "randomMath",
			RespID:  ".r",
		}
	}
	if command == ".draw" {
		return &entity.Rule{
			Command: ".draw",
			ID:      ".draw",
			Max:     100,
			Min:     0,
			Type:    "randomItem",
			RespID:  ".draw",
		}
	}
	if command == ".hitokoto" {
		return &entity.Rule{
			Command: ".hitokoto",
			ID:      ".hitokoto",
			Max:     100,
			Min:     0,
			Type:    "api",
			RespID:  ".hitokoto",
			APIID:   ".hitokoto",
		}
	}
	if command == ".weather" {
		return &entity.Rule{
			Command: ".weather",
			ID:      ".weather",
			Max:     100,
			Min:     0,
			Type:    "api",
			RespID:  ".weather",
			APIID:   ".weather",
		}
	}
	return nil
}
