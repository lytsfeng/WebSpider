package analzyer

import (
	"net/http"
	"WebSpider/base"
)
//被用于解析HTTP请求效应的函数类型
type ParsePesPonse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)