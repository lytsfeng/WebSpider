package scheduler

import (
	"net/http"
	"WebSpider/itempipeline"
	"WebSpider/analzyer"
)


// 被用来生成HTTP客户端的函数类型。
type GenHttpClient func() *http.Client

type Scheduler interface {
	/**
		启动调度器
		该方法会使调度器创建和初始化各个组件。再此之后调度器会激活爬取流程的执行

		channelLen				数据传输通道的长度
		poolSize				下载器池和分析器池的容量
		crawlDepth				爬取的最大深度，超过就忽略
		httpClientGennerator	生成HTTP客户端的函数
		respParsers				分析器所需的被用来解析HTTP响应的函数的序列。
		itemProcessors			需要被置入条目处理管道中的条目处理器的序列。
		firstHttpReq			代表首次请求。调度器会以此为起始点开始执行爬取流程。
	*/
	Start(channelLen uint,poolSize uint32,crawDepth uint32,
			httpClientGennerator GenHttpClient,
			respParsers [] analzyer.ParsePesPonse,
			itemProcessors [] itempipeline.ProcessorItem,
			firstHttpReq *http.Request) (err error)
	Stop() bool 		// 停止调度器运行
 	Running() bool 		//判断调度器是否正在运行
	/**
		获取错误通道
		调度器以及各个处理模块在运行过程中出现所有的错误都会被送到该通道
		如果放回nil 说明错误通道不可用或调度器已经停止
	 */
	ErrorChan() <- chan error
	//判断调度器时候处于空闲状态
	Idle() bool
	//获取摘要信息
	Summary(prefix string) SchedSummary
}
