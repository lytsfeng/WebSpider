package downloader

type PageDownloaderPool interface {
	Take() (PageDownloader,error) 	//获取一个网页下载器
	Return(pdl PageDownloader) 		//回收一个网页下载器
	Total() uint32 					//获取池的总容量
	Used() uint32					//获取正在使用的网页下载器数量
}
