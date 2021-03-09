package engine

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
	"github.com/dezhiShen/MiraiGo-Bot/tools"
)

// CallAPI api访问
func CallAPI(api *entity.API, context map[string]interface{}) (string, error) {
	if api.Method == entity.HTTPGet {
		res, err := http.Get(tools.ParseTpl(api.URI, context))
		if err != nil {
			return "", err
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return "", err
		}
		resp := string(robots)
		return processAPIResp(resp, api)
	}
	return "", nil
}

func processAPIResp(resp string, api *entity.API) (string, error) {
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
