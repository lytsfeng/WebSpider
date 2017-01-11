package downloader

import (
	"WebSpider/base"
	"net/http"
	"WebSpider/middleware"
)

//id生成器
var downloaderIdGenerator middleware.IdGenerator
//生成ＩＤ
func genDownloaderId() uint32  {
	return downloaderIdGenerator.GetUint32()
}


type PageDownloader interface {
	Id() uint32 // 通过ID获取实例。
	//根据请求下载网页并返回响应
	Download(req base.Request) (*base.Response, error)
}

type myPageDownloader struct {
	httpClient http.Client		//Http客户端
	id         uint32			//ID
}

func NewPageDownloader(client *http.Client) PageDownloader  {
	id := genDownloaderId()
	if client == nil{
		client = &http.Client{}
	}
	return &myPageDownloader{id:id,httpClient:*client}
}


func (pdl * myPageDownloader) Id() uint32{
	return pdl.id
}
//根据请求下载网页并返回响应
func (pdl * myPageDownloader) Download(req base.Request) (*base.Response, error) {

	httpReq := req.HttpReq()
	resp, err := pdl.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return base.NewResponse(resp,req.Depth()), nil
}
