package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type JianShu struct {
	rtype string
}

func (j JianShu) getType() string {
	return j.rtype
}

func (j JianShu) Run(name string) string {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://www.jianshu.com/search/do?q=%s&type=note&page=1&order_by=default", name), nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Csrf-Token", "YE+WpV8Dkpb8lxrYtRlPqi2MeAG7KgUBLwyk/9V+K09htk9sETMHpNgg6OP7htfBcIMqtU17Fcsj4YZlMGPJuQ==")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Set("Origin", "https://www.jianshu.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.jianshu.com/search?q=vim&page=1&type=note")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cookie", "__yadk_uid=viibcAHXupjvJPX9DOXXUujITQsJRf1H; _ga=GA1.2.1213035727.1611208474; UM_distinctid=178b6ed6e40e51-07311056833b12-11114659-13c680-178b6ed6e416fe; read_mode=day; default_font=font2; locale=zh-CN; _gid=GA1.2.855198170.1633136228; remember_user_token=W1sxMjk2NjQ4NF0sIiQyYSQxMSQ1OWpDT3BET05ldjVIWU02M1VFZUpPIiwiMTYzMzE5NjAxMy4zODA4MTY3Il0%3D--024ca32b5d1d47ee1a4a305d35d4953f77feca9d; _m7e_session_core=b5a8c4bdc0e9618e50b6a06dbb6f8c04; Hm_lvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1632925175,1632972190,1633136228,1633196014; CNZZDATA1279807957=1039469270-1617971830-https%253A%252F%252Fwww.baidu.com%252F%7C1633190468; _gat=1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2212966484%22%2C%22first_id%22%3A%2217723823cfeddb-0ca79da8d350d8-6915227c-1296000-17723823cff1906%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_utm_source%22%3A%22desktop%22%2C%22%24latest_utm_medium%22%3A%22search-input%22%2C%22%24latest_utm_campaign%22%3A%22maleskine%22%2C%22%24latest_utm_content%22%3A%22note%22%7D%2C%22%24device_id%22%3A%2217723823cfeddb-0ca79da8d350d8-6915227c-1296000-17723823cff1906%22%7D; Hm_lpvt_0c0e9d9b1e7d617b3e6842e85b9fb068=1633199019")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	mdSTR := "## jianshu \r\n"
	result := JianShuResult{}
	json.Unmarshal(body, &result)

	for _, item := range result.Entries {

		var (
			title = ""
			link = ""
		)

		link = fmt.Sprintf("https://jianshu.com/s/%v", item.Slug)
		title = item.Title

		fmt.Println(title, TrimHtml(title), link)
		mdSTR += fmt.Sprintf("[%v](%v) \r\n", title, link)
	}

	return mdSTR
}

type JianShuResult struct {
	Q          string `json:"q"`
	Page       int    `json:"page"`
	Type       string `json:"type"`
	TotalCount int    `json:"total_count"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
	OrderBy    string `json:"order_by"`
	Entries    []struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Slug    string `json:"slug"`
		Content string `json:"content"`
		User    struct {
			ID        int    `json:"id"`
			Nickname  string `json:"nickname"`
			Slug      string `json:"slug"`
			AvatarURL string `json:"avatar_url"`
		} `json:"user"`
		Notebook struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"notebook"`
		Commentable         bool      `json:"commentable"`
		PublicCommentsCount int       `json:"public_comments_count"`
		LikesCount          int       `json:"likes_count"`
		ViewsCount          int       `json:"views_count"`
		TotalRewardsCount   int       `json:"total_rewards_count"`
		FirstSharedAt       time.Time `json:"first_shared_at"`
	} `json:"entries"`
	RelatedUsers []struct {
		ID              int    `json:"id"`
		AvatarURL       string `json:"avatar_url"`
		Nickname        string `json:"nickname"`
		Slug            string `json:"slug"`
		TotalWordage    int    `json:"total_wordage"`
		TotalLikesCount int    `json:"total_likes_count"`
	} `json:"related_users"`
	MoreRelatedUsers   bool `json:"more_related_users"`
	RelatedCollections []struct {
		ID               int    `json:"id"`
		Title            string `json:"title"`
		Slug             string `json:"slug"`
		ImageURL         string `json:"image_url"`
		PublicNotesCount int    `json:"public_notes_count"`
		LikesCount       int    `json:"likes_count"`
	} `json:"related_collections"`
	MoreRelatedCollections bool `json:"more_related_collections"`
}