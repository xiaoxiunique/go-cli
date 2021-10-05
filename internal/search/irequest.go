package search

import "net/http"

// Req 请求接口
type Req interface {
	getType() string
	Run(interface{}) interface{}
}

type ReqConf struct {
	url string
	method string
	payload string
	headers http.Header
}

// ResItem 反馈抽象类型
type ResItem struct {
	title string
	url string
	image string
	likeCount int
	commentCount int
}