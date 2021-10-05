package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Zhihu struct {
	rtype string
}

func (j Zhihu) getType() string {
	return j.rtype
}

func (j Zhihu) Run(name string) string {
	var req, err = http.NewRequest("GET", fmt.Sprintf("https://www.zhihu.com/api/v4/search_v3?t=general&q=%s&correction=1&offset=0&limit=20&filter_fields=&lc_idx=0&show_all_topics=0", name), nil)
	if err != nil {
		fmt.Println("Request Faild")
	}
	req.Header.Set("Authority", "www.zhihu.com")
	req.Header.Set("X-Zse-93", "101_3_2.0")
	req.Header.Set("X-Ab-Param", "tp_dingyue_video=0;top_test_4_liguangyi=1;tp_zrec=1;pf_noti_entry_num=2;qap_question_visitor= 0;zr_expslotpaid=1;zr_slotpaidexp=1;tp_topic_style=0;qap_question_author=0;se_ffzx_jushen1=0;tp_contents=2;pf_adjust=1")
	req.Header.Set("X-Ab-Pb", "CrgB2AJRBUAB4QRgC8oCMgOOA+AEogM3BT8AaQH0A1cETwNSC0MAtADWBDIFOwIzBNcC9gILBMcCtwPpBFUFAQsKBEUE9AsbAEcAbAMqBDQM4At9AgcM+APXC+gDZATMBBEFGQVWBVADNwyJDLQK7AoqAp8CNAQOBTMFzAJXA6YEmwsUBQ8LhALzA9wLjAJSBc8LVgzjBHQBcgPRBNoE9wMpBXUEEgVqAaADQgS1CxUF5Aq5AqEDbATBBBJcAAABAgAAAAEAAAAAAQABAAEVAAAAAAAAAAAAAAEAAAEBAAAAAQAAAAEBAQAAAQEAAAAAAQAAAQABAAACAAkAAgIBAAAAAAALAQIAAQUAAQAAAAAAAAMAABUAAAE=")
	req.Header.Set("X-Api-Version", "3.0.91")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Set("X-Requested-With", "fetch")
	req.Header.Set("X-Zse-96", "2.0_aMtBb0UBo0SxQ8NqKLS0FJr0r0xfoXt088xqeQUqkLSx")
	req.Header.Set("X-App-Za", "OS=Web")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", fmt.Sprintf("https://www.zhihu.com/search?type=content&q=%s", name))
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cookie", "_zap=d0cfd117-9151-4a99-903c-9b18f57358c7; d_c0=\"AHCY5gT4iBKPTsBQivYRPUIbDLBctuaZZzs=|1611209467\"; __snaker__id=EJHuxuWtL43TSuJY; _9755xjdesxxd_=32; YD00517437729195%3AWM_TID=49%2FqJQyMZ9lEAVRARRJqbx0z41JF8BjO; tshl=; gdxidpyhxdE=lZZmQcGIEMBMJJtcD8gdf2ScLEuto%5CubOcDgIw%2BADCV36rh92DK6WhjDfOj51EOhbOjJGCEIPjBThwhjQHWE5JB3W1XHtuPsMCUYxfU0BdY%5CVUxGfNSygdIwBn95evJVZ%5Cv0QYAOiAeUJbP%2BdelSlKjiaon%5CXBZS3Gr4DaxcnPpIEMez%3A1627021591797; YD00517437729195%3AWM_NI=A9OYDLdvRLod47%2BTiUKeAdSRPSUOaDm4Zfj7k09S6OYdGEkuAz4A0MLGRldmvGHXk7ZTAm9n64meYT%2F2tNmH5QIKfJtPOShHEIoCbCyyju8aOy3mw1pkaQD92moUHQV6WmU%3D; YD00517437729195%3AWM_NIKE=9ca17ae2e6ffcda170e2e6ee8ef346aaa6af98e73b9ab08eb3c14f969f8e85ae3a94939cccf86a85939798d32af0fea7c3b92a83e781bac748a8aa9eadee70f8e9e5b5ec5f89bb8f85d25a948d9bb0e450afec9687e93da9f0f7b4b5459bbc89a6ea7f83a6fca2b34790999e8de767b69988b9ea3d9a8c8f8ebc66f6eca386dc60edeebcaff044f78f87ace542b89081b4eb448797a1a4e149bba7a696bb6b8da7fd8ef63bb4b5afb2c733f2b3a1b6cb5b96ea998ed837e2a3; z_c0=\"2|1:0|10:1627020698|4:z_c0|92:Mi4xa2ZEVkF3QUFBQUFBY0pqbUJQaUlFaVlBQUFCZ0FsVk5tcXZuWVFBcEUxYkdWbm82LTZrZ0hBd0t6b2o4MF9neDhn|979d92d815466bde4d4589a34cc9d5a9047ee8775ae228f52e2ce77daac40c81\"; _xsrf=CLswPUA5uoD78xfdqgJpwfcKqBfkjhw7; q_c1=9a92be425083406cb7e101bdcb8b42b2|1628838667000|1611896095000; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1630949792,1632394840,1632878724,1632925558; tst=r; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1633138355; SESSIONID=ldpzxqyT85sQ6MgDCuEfBUrdtjqrqlv0FWhcbWaaPLD; JOID=VVsSBUKP4QTt5TAKdI67mvoGtolm54F1mJNFdyG50HO93HNpNDtSXY3vMA96K_jbERA4Uw8ifMkKLnAB7GQllQ4=; osd=UVsUAUOL4QLp5DQKcoq6nvoAsohi54dxmZdFcSW41HO72HJtND1WXInvNgt7L_jdFRE8Uwkmfc0KKHQA6GQjkQ8=; KLBRSID=4843ceb2c0de43091e0ff7c22eadca8c|1633138394|1633138353")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request Faild")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	mdSTR := "## zhihu \r\n"
	result := ResultObject{}
	json.Unmarshal(body, &result)

	for _, item := range result.Data {

		var (
			title = ""
			link = ""
			itemType = item.Type
			itemObjectType = item.Object.Type
			isZVideoType = "zvideo" == itemType
			isAnswer = "answer" == itemType
			isSearchAnswer = "search_result" == itemType && itemObjectType == "answer"
			isSearchArticle = "search_result" == itemType && itemObjectType == "article"
		)

		if isZVideoType {
			title = item.Object.Title
			link = item.Object.VideoUrl
		}

		if isAnswer || isSearchAnswer {
			title = item.Object.Question.Name
			link = item.Object.Question.Url
		}

		if isSearchArticle {
			title = item.Object.Title
		}

		if len(title) == 0 || len(link) == 0 {
			continue
		}

		fmt.Println(title, TrimHtml(title), link)
		mdSTR += fmt.Sprintf("[%v](%v) \r\n", title, link)
	}

	return mdSTR
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("<style[\\S\\s]+?</style>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("<script[\\S\\s]+?</script>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	src = strings.Replace(src, "\\n", "", -1)
	src = strings.Replace(src, "\\r", "", -1)
	return strings.TrimSpace(src)
}

type ResultObject struct {
	Paging struct {
		IsEnd bool   `json:"is_end"`
		Next  string `json:"next"`
	} `json:"paging"`
	Data []struct {
		Type      string `json:"type"`
		Highlight struct {
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"highlight,omitempty"`
		Object struct {
			ID             string        `json:"id"`
			Content           string        `json:"content"`
			Name           string        `json:"name"`
			Title           string        `json:"title"`
			URL            string        `json:"url"`
			Type           string        `json:"type"`
			Excerpt        string        `json:"excerpt"`
			Introduction   string        `json:"introduction"`
			Description    string        `json:"description"`
			AvatarURL      string        `json:"avatar_url"`
			IsFollowing    bool          `json:"is_following"`
			QuestionsCount int           `json:"questions_count"`
			FollowersCount int           `json:"followers_count"`
			FollowerCount  int           `json:"follower_count"`
			TopAnswerCount int           `json:"top_answer_count"`
			Score          float64       `json:"score"`
			Aliases        []interface{} `json:"aliases"`
			HasMeta        bool          `json:"has_meta"`
			VideoUrl        string          `json:"videoUrl"`
			Question struct{
				Name string `json:"name"`
				Id   string `json:"id"`
				Type   string `json:"type"`
				Url   string `json:"url"`
			}
			Meta           struct {
				ID              int64         `json:"id"`
				Type            string        `json:"type"`
				Category        string        `json:"category"`
				Description     string        `json:"description"`
				ImageURL        string        `json:"image_url"`
				IsFollowing     bool          `json:"is_following"`
				FollowerCount   int           `json:"follower_count"`
				QuestionCount   int           `json:"question_count"`
				DiscussionCount int           `json:"discussion_count"`
				Description1    string        `json:"description1"`
				Description2    string        `json:"description2"`
				Subtitle        string        `json:"subtitle"`
				Title           string        `json:"title"`
				Tags            []interface{} `json:"tags"`
				AvatarURL       string        `json:"avatar_url"`
				Template        struct {
					MetaTemplateType int `json:"meta_template_type"`
					AvatarType       int `json:"avatar_type"`
				} `json:"template"`
				WikiDescription []interface{} `json:"wiki_description"`
				Scores          []interface{} `json:"scores"`
			} `json:"meta"`
			TopicType         string `json:"topic_type"`
			AttachedInfoBytes string `json:"attached_info_bytes"`
			WikiType          string `json:"wiki_type"`
			DiscussionCount   int    `json:"discussion_count"`
			HasIndexSection   bool   `json:"has_index_section"`
			EssenceFeedCount  int    `json:"essence_feed_count"`
		} `json:"object,omitempty"`
		Index     int    `json:"index,omitempty"`
		ID        string `json:"id,omitempty"`
		QueryList []struct {
			Query             string `json:"query"`
			AttachedInfoBytes string `json:"attached_info_bytes"`
		} `json:"query_list,omitempty"`
		AttachedInfoBytes string `json:"attached_info_bytes,omitempty"`
	} `json:"data"`
	RelatedSearchResult []interface{} `json:"related_search_result"`
	SearchActionInfo    struct {
		AttachedInfoBytes string `json:"attached_info_bytes"`
		LcIdx             int    `json:"lc_idx"`
		SearchHashID      string `json:"search_hash_id"`
		Isfeed            bool   `json:"isfeed"`
	} `json:"search_action_info"`
	IsBrandWord    bool        `json:"is_brand_word"`
	Pendant        interface{} `json:"pendant"`
	SensitiveLevel int         `json:"sensitive_level"`
	Warning        string      `json:"warning"`
}
