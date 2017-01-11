package analzyer

import (
	"WebSpider/base"
	"WebSpider/middleware"
	"errors"
	"net/url"
	"logging"
	"fmt"
)



type Analyzer interface {
	Id() uint32
	//根据规则分析响应并放回请求和条目
	Analyze(respParsers []ParsePesPonse, resp base.Response) ([]base.Data, []error)
}
type myAnalyzer struct {
	id uint32
}

var analyzerIdGenerator middleware.IdGenerator = middleware.NewIdGenerator()
// 日志记录器。
var logger logging.Logger = base.NewLogger()

func genAnalyzerId() uint32{
	return analyzerIdGenerator.GetUint32()
}

func NewAnalyzer() Analyzer{
	return &myAnalyzer{id:genAnalyzerId()}
}
func (analyzer *myAnalyzer) Id() uint32 {
	return analyzer.id
}
func (analyzer *myAnalyzer)Analyze(respParsers []ParsePesPonse, resp base.Response) ([]base.Data, []error)  {
	if respParsers == nil{
		err := errors.New("The response parser list is invalid")
		return nil,[]error{err}
	}
	httpResp := resp.HttpRes()
	if httpResp == nil{
		err := errors.New("The httpResq is invalid")
		return nil, []error{err}
	}
	//记录日志
	var reqUrl *url.URL = httpResp.Request.URL
	logger.Infof("parse the response (reqUrl)=%s \n",reqUrl)

	respDepth := resp.Depth()

	dataList := make([] base.Data,0)
	errList := make([]error,0)

	for i,respParse := range respParsers{
		if respParse == nil{
			err:=errors.New(fmt.Sprintf("The document parser [%d] is invalid!", i))
			errList = append(errList,err)
		}
		pDataList, pErrorList := respParse(httpResp, respDepth)
		if pDataList != nil {
			for _, pData := range pDataList {
				dataList = appendDataList(dataList, pData, respDepth)
			}
		}
		if pErrorList != nil {
			for _, pError := range pErrorList {
				errList = appendErrorList(errList, pError)
			}
		}
		continue
	}
	return dataList,errList
}

// 添加请求值或条目值到列表。
func appendDataList(dataList []base.Data, data base.Data, respDepth uint32) []base.Data {
	if data == nil {
		return dataList
	}
	req, ok := data.(*base.Request)  // 判断data的类型是不是请求
	if !ok {
		return append(dataList, data)
	}
	newDepth := respDepth + 1
	if req.Depth() != newDepth {
		req = base.NewRequest(req.HttpReq(), newDepth)
	}
	return append(dataList, req)
}


// 添加错误值到列表。
func appendErrorList(errorList []error, err error) []error {
	if err == nil {
		return errorList
	}
	return append(errorList, err)
}