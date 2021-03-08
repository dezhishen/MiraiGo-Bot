package engine

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dezhiShen/MiraiGo-Bot/pkg/db"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
)

// RunRule 执行规则
func RunRule(rule *entity.Rule, context map[string]string) (string, error) {
	if rule == nil {
		return "", nil
	}
	if rule.Type == "randomMath" {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(rule.Max-rule.Min) + rule.Min
		return fmt.Sprint(v), nil
	}
	if rule.Type == "randomItem" {
		setKey := context["$1"]
		set := db.GetDraw(setKey)
		if set == nil {
			return "", errors.New("卡池不存在")
		}
		countStr, ok := context["$2"]
		var count int
		if !ok || countStr == "" {
			count = 1
		} else {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("参数类型错误,必须是正整数")
			}
		}
		if count > set.Size() {
			return "", errors.New("卡池中没有这么多的卡")
		}
		out := ""
		for i := 0; i < count; i++ {
			v := set.Pop()
			out += (fmt.Sprint(v) + ",")
		}
		return out[:len(out)-1], nil

	}

	return "", nil
}
