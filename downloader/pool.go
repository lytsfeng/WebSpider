package downloader

import (
	"WebSpider/middleware"
	"reflect"
	"fmt"
	"errors"
)

type PageDownloaderPool interface {
	Take() (PageDownloader, error) //获取一个网页下载器
	Return(pdl PageDownloader)   error  //回收一个网页下载器
	Total() uint32                 //获取池的总容量
	Used() uint32                  //获取正在使用的网页下载器数量
}

type myPageDownloaderPool struct {
	pool  middleware.Pool //实体池
	etype reflect.Type    //实体池类型
}

type GenPageDownloader func() PageDownloader

func NewPageDownloaderPool(total uint32, gen GenPageDownloader) (PageDownloaderPool, error) {
	etype := reflect.TypeOf(gen)
	genEntity := func() middleware.Entity{
		return gen()
	}
	_pool,err := middleware.NewPool(total,etype,genEntity)
	if err != nil{
		return nil,err
	}
	return &myPageDownloaderPool{pool:_pool,etype:etype}, nil
}

func (pd *myPageDownloaderPool)Take() (PageDownloader, error){
	entity, err := pd.pool.Take()
	if err != nil {
		return nil, err
	}
	dl, ok := entity.(PageDownloader)
	if !ok {
		errMsg := fmt.Sprintf("The type of entity is NOT %s!\n", pd.etype)
		panic(errors.New(errMsg))
	}
	return dl, nil
}
func (pd *myPageDownloaderPool) Return( pdl PageDownloader)  error{
	return pd.pool.Return(pdl)
}
func (pd *myPageDownloaderPool) Total() uint32 {
	return pd.pool.Total()
}
func (pd *myPageDownloaderPool)  Used() uint32 {
	return pd.pool.Used()
}


