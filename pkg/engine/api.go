package engine

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
	"github.com/dezhiShen/MiraiGo-Bot/tools"
)

// CallAPI api访问
func CallAPI(api *entity.API, context map[string]interface{}) (string, error) {
	var r *http.Response
	var err error
	var resp string
	url := tools.ParseTpl(api.URI, context)
	if entity.HTTPGet == api.Method {
		r, err = http.DefaultClient.Get(url)
	} else if entity.HTTPPost == api.Method {
		body := tools.ParseTpl(api.RequestTemplate, context)
		r, err = http.DefaultClient.Post(url, "application/json", strings.NewReader(body))
	} else {
		return "", errors.New("不支持的请求格式,当前只支持get和post")
	}
	if err != nil {
		return "", err
	}
	robots, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return "", err
	}
	resp = string(robots)
	return processAPIResp(resp, api, context)
}

func processAPIResp(resp string, api *entity.API, context map[string]interface{}) (string, error) {
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
	for k, v := range context {
		mapResult[k] = v
	}
	template = tools.ParseTpl(template, mapResult)
	return template, nil
}
