package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/dezhiShen/MiraiGo-Bot/pkg/db"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
	"github.com/dezhiShen/MiraiGo-Bot/tools"
)

// RunRule 执行规则
func RunRule(rule *entity.Rule, context map[string]interface{}) (string, error) {
	if rule == nil {
		return "", nil
	}
	if rule.Type == "randomMath" {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(rule.Max-rule.Min) + rule.Min
		return fmt.Sprint(v), nil
	}
	if rule.Type == "randomItem" {
		setKey := context["1"]
		set := db.GetDraw(tools.ToString(setKey))
		if set == nil {
			return "", errors.New("卡池不存在")
		}
		countStr, ok := context["2"]
		var count int
		if !ok || countStr == "" {
			count = 1
		} else {
			var err error
			count, err = strconv.Atoi(tools.ToString(countStr))
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
	if rule.Type == "api" {
		api := db.GetAPI(rule.APIID)
		if api == nil {
			return "", nil
		}
		if api.Method == "get" {
			if api.RequestTemplate == "" {
				res, err := http.Get(api.URI)
				if err != nil {
					return "", err
				}
				robots, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					return "", err
				}
				resp := string(robots)
				return processApiResp(resp, api)
			}
		}

	}

	return "", nil
}

func processApiResp(resp string, api *entity.API) (string, error) {
	if resp == "" {
		return "", nil
	}
	if api.ResponseTemplate == "" {
		return "", nil
	}
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(resp), &mapResult)
	if err != nil {
		return "", err
	}
	template := api.ResponseTemplate
	template = tools.ParseTpl(template, mapResult)
	return template, nil
}
