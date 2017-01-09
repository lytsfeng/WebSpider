package analzyer

import (
	"WebSpider/base"
)


type Analyzer interface {
	Id() uint32
	//根据规则分析响应并放回请求和条目
	Analyze(respParsers []ParsePesPonse, resp base.Response) ([]base.Data, []error)
}
