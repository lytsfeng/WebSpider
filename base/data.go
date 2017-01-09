package base

import "net/http"



//数据的接口
type Data interface {
	Valid() bool  //验证
}

//请求
type Request struct{
	httpReq *http.Request   //http请求指针
	depth  uint32			//请求深度
}

//创建新的请求
func NewRequess(httpReq *http.Request,depth uint32) *Request{
	return &Request{httpReq:httpReq,depth:depth}
}

//获取http请求
func (req *Request) HttpReq() *http.Request  {
	return req.httpReq
}
//获取深度
func (req *Request) Depth() uint32  {
	return req.depth
}
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

//响应
type Response struct {
	httpRes *http.Response  // http 响应
	depth uint32
}
func NewResponse(httpRes http.Response,depth uint32) *Response {
	return &Response{httpRes:httpRes,depth:depth}
}
func (res *Response) HttpRes() *http.Response  {
	return res.httpRes
}
func (res *Response) Depth() uint32  {
	return res.depth
}
func (res *Response) Valid() bool {
	return res.httpRes != nil && res.httpRes.Body != nil
}

//条目
type Item map[string]interface{}

func (item Item)Valid() bool  {
	return item != nil
}


