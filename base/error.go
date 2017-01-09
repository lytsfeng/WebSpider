package base

import (
	"bytes"
	"fmt"
)

type ErrorType string

const (
	DOWNLOADER_ERROR     ErrorType = "Downloader Error"
	ANALYZER_ERROR       ErrorType = "Analyzer Error"
	ITEM_PROCESSOR_ERROR ErrorType = "Item Processoro Error"
)
// 爬虫错误的接口。
type CrawlerError interface {
	Type() 	ErrorType
	Error() string
}
// 爬虫错误的实现。
type myCrawlerError struct {
	errType    ErrorType //错误类型
	errMsg     string    //错误提示信息
	fullErrMsg string    //完整错误信息
}
// 创建一个新的爬虫错误。
func NewCrawlerError(errType ErrorType,errMsg string)  CrawlerError {
	return &myCrawlerError{errType:errType,errMsg:errMsg}
}

func (ce *myCrawlerError) Type() ErrorType  {
	return ce.errType
}

func (ce *myCrawlerError)Error() string{
	if ce.fullErrMsg == ""{
		ce.getFullErrMsg()
	}
	return ce.fullErrMsg
}

func (ce *myCrawlerError) getFullErrMsg(){
	var buffer bytes.Buffer
	buffer.WriteString("Error WebSpider")
	if ce.errType != ""{
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprint("%s\n",buffer.String())
	return
}