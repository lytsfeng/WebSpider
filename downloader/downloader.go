package downloader

import "webcrawler/base"

type PageDownloader interface {
	Id() uint32
	//根据请求下载网页并返回响应
	Download(req base.Request) ( *base.Response,error)
}









