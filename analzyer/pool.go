package analzyer

import (
	"WebSpider/middleware"
	"reflect"
	"fmt"
	"errors"
)

type AnalyzerPool interface {
	Take() (Analyzer,error)    		// 获取一个解析器
	Return(analyzer Analyzer) error	// 回收一个解析器
	Total() uint32					// 解析器个数
	Used() uint32					//这在使用的解析器个数
}


type myAnalyzerPool struct {
	pool  middleware.Pool //实体池
	etype reflect.Type    //实体池类型
}

type genAnalyzer func() Analyzer
func NewAnalyzerPool(total uint32,gen genAnalyzer) (AnalyzerPool,error)  {
	etype := reflect.TypeOf(gen())
	genEntity := func() middleware.Entity{
		return gen()
	}
	_pool,err := middleware.NewPool(total,etype,genEntity)
	if err != nil{
		return nil,err
	}
	dlpool := &myAnalyzerPool{pool: _pool, etype: etype}
	return dlpool, nil
}

func (pd *myAnalyzerPool)Take() (Analyzer, error){
	entity, err := pd.pool.Take()
	if err != nil {
		return nil, err
	}
	dl, ok := entity.(Analyzer)
	if !ok {
		errMsg := fmt.Sprintf("The type of entity is NOT %s!\n", pd.etype)
		panic(errors.New(errMsg))
	}
	return dl, nil
}
func (pd *myAnalyzerPool) Return( pdl Analyzer)  error{
	return pd.pool.Return(pdl)
}
func (pd *myAnalyzerPool) Total() uint32 {
	return pd.pool.Total()
}
func (pd *myAnalyzerPool)  Used() uint32 {
	return pd.pool.Used()
}