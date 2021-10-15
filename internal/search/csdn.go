package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Csdn struct {
	rtype string
}

func (c Csdn) getType() string {
	return c.rtype
}

func (c Csdn) Run(name string) string {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://so.csdn.net/api/v3/search?q=%s&t=blog&p=1&s=hot&tm=0&lv=-1&ft=0&l=&u=&ct=-1&pnt=-1&ry=-1&ss=-1&dct=-1&vco=-1&cc=-1&sc=-1&akt=-1&art=-1&ca=-1&prs=&pre=&ecc=-1&ebc=-1&urw=&ia=1&platform=pc", name), nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "so.csdn.net")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://so.csdn.net/so/search?q=vim&t=blog&u=&s=hot")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cookie", "uuid_tt_dd=10_30631841840-1611208208172-705312; _ga=GA1.2.680810348.1611208210; UserName=UNIQUE_XIAOXI; UserInfo=37030658c6324cfda54d9ead7d9faab3; UserToken=37030658c6324cfda54d9ead7d9faab3; UserNick=atom2018; AU=43D; UN=UNIQUE_XIAOXI; BT=1625629818802; p_uid=U010000; ssxmod_itna=7q+xgD0CqYqWq0dD=eYEg8KDttTL+rbR4GN+YGRDCqAPGfDI4fa8qx+xgxOqBIbtiD7KdKn1f5+TwFSq2DIr5bDvxYD84i7DKqibLDiH3DjxOqeDxxDUcmTxDeje3291fCb3o44tKAOAmhKRv583xd97Gxm7D597ALKRIoOA0xiL4C9eDAff4DOiD===; ssxmod_itna2=7q+xgD0CqYqWq0dD=eYEg8KDttTL+rb4A=Rq4eD/SYDFhOu0PBIxGOQWx8rRZZ4XkUilkDXHPoiTpb7YCnmYj7fbFuSXsL7Qp0KauSKmGXhd6t9A35zO9aoBUMD3rc0BfKBPQ=X5kTEwK+UsDu5xXup4rd0clOv2zMXj3Rx4flvAGu0QrqXcD=YrxTnwBe6L44aIDSvoa+E+bnER0j0+nob6iky2t8HGmpHGX+IAinPQDu79Gg5Qa1BhmivH9BaalhGu7zLcjQM4VPoz5UAk7AlUgEL6xUdokypu7f0u39p9SZpZ48vhjNMPN025El4olGREmkWtGGEOOmaCpw7Itg5loPKOGk=vUAL9D0jam1BXxUAQFAGOXHnwcnR6xKbUXwDurUYTFU1BvV=5UoUFI+cjdSm5b3W4epwSN0ARRQWstwzxoKYQzDEpDcPZuz=vCbWE4NC3tX/3bDG2U2=CGSefxFdQADZFkks2kkwYjkWCL=G2KQuDiDujkswhNAxu3Rs54qbGRK0zXIe7Q51hxjFUCDSAhsBls6+zyCvExDFqD+64eQ4O7PqYLK+D8Aw7k=MmWIohWNDLrDnDPpbGf8pK7P5YL=YdHkpfOpwceUx8em8q0DxD; Hm_up_6bcd52f51e9b3dce32bec4a3997715ac=%7B%22islogin%22%3A%7B%22value%22%3A%221%22%2C%22scope%22%3A1%7D%2C%22isonline%22%3A%7B%22value%22%3A%221%22%2C%22scope%22%3A1%7D%2C%22isvip%22%3A%7B%22value%22%3A%220%22%2C%22scope%22%3A1%7D%2C%22uid_%22%3A%7B%22value%22%3A%22UNIQUE_XIAOXI%22%2C%22scope%22%3A1%7D%7D; Hm_ct_6bcd52f51e9b3dce32bec4a3997715ac=6525*1*10_30631841840-1611208208172-705312\\u00215744*1*UNIQUE_XIAOXI; c_first_ref=www.google.com; c_segment=0; dc_sid=ae9d5406c87627901977cdf2fc0d878c; _gid=GA1.2.1870584063.1633107793; dc_session_id=10_1633196616921.518591; c_first_page=https%3A//www.csdn.net/; c_page_id=default; Hm_lvt_6bcd52f51e9b3dce32bec4a3997715ac=1632989748,1633107793,1633139742,1633196639; c_pref=https%3A//www.google.com/; c_ref=https%3A//www.csdn.net/; __51uvsct__JQTDiOVZ2pRjGa1K=1; __51vcke__JQTDiOVZ2pRjGa1K=0fbb2ad0-a4a2-5c35-bc6b-d4abb24382c8; __51vuft__JQTDiOVZ2pRjGa1K=1633196644531; showTip=false; Hm_lpvt_6bcd52f51e9b3dce32bec4a3997715ac=1633196665; __vtins__JQTDiOVZ2pRjGa1K=%7B%22sid%22%3A%20%2225fefa9a-aa53-5a39-a0b0-7dfdec582825%22%2C%20%22vd%22%3A%202%2C%20%22stt%22%3A%2022621%2C%20%22dr%22%3A%2022621%2C%20%22expires%22%3A%201633198467144%2C%20%22ct%22%3A%201633196667144%7D; searchHistoryArray-new=%5B%22vim%22%5D; dc_tos=r0d2n5; log_Id_pv=681; log_Id_view=1107; log_Id_click=192")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}

	mdSTR := "## csdn \r\n"
	result := Result{}
	json.Unmarshal(body, &result)

	for _, item := range result.ResultVos {

		var (
			title = ""
			link = ""
			isBlog = item.Type == "blog"
		)

		if isBlog {
			link = item.URL
			title = item.Title
		}

		mdSTR += fmt.Sprintf("[%v](%v) \r\n", title, link)
	}

	return mdSTR
}

type Result struct {
	ShowLiveStatus  bool        `json:"show_live_status"`
	IsNavwordResult interface{} `json:"is_navword_result"`
	AbTestExt       struct {
		ReplaceNew bool `json:"replace_new"`
	} `json:"ab_test_ext"`
	JsInsertCount  int         `json:"js_insert_count"`
	MakeFlag       interface{} `json:"make_flag"`
	QueryRewrite   string      `json:"query_rewrite"`
	SplitWords     []string    `json:"split_words"`
	IsContainBaidu string      `json:"is_contain_baidu"`
	IsCsdnURL      interface{} `json:"is_csdn_url"`
	IsNavword      string      `json:"is_navword"`
	NavwordScene   bool        `json:"navword_scene"`
	Total          int         `json:"total"`
	Experiment     string      `json:"experiment"`
	SensitiveCode  bool        `json:"sensitive_code"`
	ResultVos      []struct {
		Digg           string        `json:"digg"`
		Language       string        `json:"language"`
		Pic            string        `json:"pic"`
		Type           string        `json:"type"`
		Title          string        `json:"title"`
		View           string        `json:"view"`
		EventClick     bool          `json:"eventClick"`
		Ext1           string        `json:"ext_1"`
		Ext10          string        `json:"ext_10"`
		Nickname       string        `json:"nickname"`
		Digest         string        `json:"digest"`
		ID             string        `json:"id"`
		Ext6           string        `json:"ext_6"`
		OpsRequestMisc string        `json:"ops_request_misc"`
		CreateTime     string        `json:"create_time"`
		IsNew          bool          `json:"is_new"`
		Author         string        `json:"author"`
		VipViewAuth    string        `json:"vip_view_auth"`
		Ext15          string        `json:"ext_15"`
		Index          string        `json:"index"`
		URL            string        `json:"url"`
		Tags           []interface{} `json:"tags"`
		Popu           string        `json:"popu"`
		UtmTerm        string        `json:"utm_term"`
		CreateTimeStr  string        `json:"create_time_str"`
		Ext23          string        `json:"ext_23"`
		Ext20          string        `json:"ext_20"`
		Ext21          string        `json:"ext_21"`
		SoType         string        `json:"so_type"`
		Comment        string        `json:"comment"`
		Style          string        `json:"style"`
		BizID          string        `json:"biz_id"`
		Strategy       string        `json:"strategy"`
		RequestID      string        `json:"request_id"`
		URLLocation    string        `json:"url_location"`
		EventView      bool          `json:"eventView"`
		UtmSource      string        `json:"utm_source"`
		SearchTag      []string      `json:"search_tag,omitempty"`
		PicList        []string      `json:"pic_list,omitempty"`
	} `json:"result_vos"`
	Isc           bool    `json:"isc"`
	TotalPage     float64 `json:"total_page"`
	JsInsertFirst bool    `json:"js_insert_first"`
}
