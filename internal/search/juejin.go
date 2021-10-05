package search

import (
	"encoding/json"
	"fmt"
	"go-cli/types"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

type Juejin struct {
	rtype string
}

// 返回当前实现类的类型
func (j Juejin) getType() string {
	return j.rtype
}

func (j Juejin) Run(name string) string {
	fmt.Println("----- start invoke jujin search processing -----")

	url := "https://api.juejin.cn/search_api/v1/search?aid=2608&uuid=6920203267013019139"
	method := "POST"

	if len(name) == 0 {
		panic("无搜索参数")
	}

	search := url2.PathEscape(name)
	payload := strings.NewReader(fmt.Sprintf(`{ "key_word": "%s", "id_type": 0, "cursor": "0", "limit": 20, "search_type": 0, "sort_type": 0, "version": 1, "uuid": "6920203267013019139" }`, search))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "api.juejin.cn")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://juejin.cn")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://juejin.cn/")
	req.Header.Add("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("cookie", "_ga=GA1.2.75892462.1611235385; n_mh=MVzRqTFwJq1fajMuA549e9J1YOAsuIcvxYlYO9ELep8; MONITOR_WEB_ID=9253c471-8173-4ccc-8f3b-9921eeca4180; passport_csrf_token_default=2ac163214f485a184ef120ca46095b4a; passport_csrf_token=2ac163214f485a184ef120ca46095b4a; passport_auth_status=c2f5e41a68719e99d2f82dd5dd7dbd43%2C; passport_auth_status_ss=c2f5e41a68719e99d2f82dd5dd7dbd43%2C; sid_guard=1a9bf9c908311a25bb3cf29f2b0ff9c9%7C1632465201%7C5184000%7CTue%2C+23-Nov-2021+06%3A33%3A21+GMT; uid_tt=b554271bf6ab7ea599cf177d9a9cad34; uid_tt_ss=b554271bf6ab7ea599cf177d9a9cad34; sid_tt=1a9bf9c908311a25bb3cf29f2b0ff9c9; sessionid=1a9bf9c908311a25bb3cf29f2b0ff9c9; sessionid_ss=1a9bf9c908311a25bb3cf29f2b0ff9c9; sid_ucp_v1=1.0.0-KDFlNTAxYWI4NThhZTcwM2EzOGYwNTYyZGQ2Y2QxMTdmNzgwZmU1NjkKFgj4j8C__fVpELHitYoGGLAUOAJA8QcaAmxmIiAxYTliZjljOTA4MzExYTI1YmIzY2YyOWYyYjBmZjljOQ; ssid_ucp_v1=1.0.0-KDFlNTAxYWI4NThhZTcwM2EzOGYwNTYyZGQ2Y2QxMTdmNzgwZmU1NjkKFgj4j8C__fVpELHitYoGGLAUOAJA8QcaAmxmIiAxYTliZjljOTA4MzExYTI1YmIzY2YyOWYyYjBmZjljOQ; _tea_utm_cache_2608={%22utm_source%22:%22web_feed5%22%2C%22utm_medium%22:%22banner%22%2C%22utm_campaign%22:%22daka_xiaozhishi_20210924%22}; _gid=GA1.2.714645301.1632802616")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	mdSTR := " ## juejin \r\n"
	result := types.AutoGenerated{}
	json.Unmarshal(body, &result)
	for _, item := range result.Data {

		title := item.ResultModel.ArticleInfo.Title
		if title == "" {
			continue
		}

		link := fmt.Sprintf("https://juejin.cn/%s", item.ResultModel.ArticleID)
		mdSTR += fmt.Sprintf("[%s](%s) \r\n", title, link)
	}

	fmt.Println("----- start invoke jujin search successful -----")

	return mdSTR
}
