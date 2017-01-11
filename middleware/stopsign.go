package middleware

import (
	"sync"
	"fmt"
)

type StopSign interface {
	/**
	设置停止新信号，相当与发送停止信号
	如果已经发送过停止信号 返回false
	 */
	Sign() bool
	//判断停止型号时候已经被发出
	Signed() bool
	//重置信号  相当于回收信号，并清除所有停止信号的处理记录，相当于初始化
	Reset()
	/**
		处理停止信号
		code代表停止信号处理的代号，该代号出现在信号处理的记录中
	 */
	Deal(code string)
	//获取某一个停止信号处理方式的处理计数。该处理计数会从相应的停止信号处理记录中获得
	DealCount(code string) uint32
	////获取停止信号被处理的总数
	DealTotal() uint32
	//概要
	Summary() string
}

type myStopSign struct {
	signed 			bool 				//表示信号是否已经发出的标志位
	dealCountMap	map[string]uint32 	//处理计数的字典
	rwmutex 		sync.RWMutex		//读写锁
}

func NewStopSign() StopSign {
	_StopSign := &myStopSign{dealCountMap:make(map[string]uint32)}
	return _StopSign
}

func (ss *myStopSign) Sign()  bool {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()
	if ss.signed{
		return false
	}
	ss.signed = false
	return true
}

func (ss *myStopSign) Signed() bool  {
	return ss.signed
}

func (ss *myStopSign) Deal(code string){
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()
	if !ss.signed {
		return
	}

	if _,ok := ss.dealCountMap[code]; !ok{
		ss.dealCountMap[code] = 1
	} else {
		ss.dealCountMap[code] += 1
	}
}
func (ss *myStopSign) Reset()  {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()
	ss.signed = false
	ss.dealCountMap = make(map[string]uint32)
}

//获取某一个停止信号处理方式的处理计数。该处理计数会从相应的停止信号处理记录中获得
func (ss *myStopSign) DealCount(code string) uint32 {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()
	return ss.dealCountMap[code]
}
////获取停止信号被处理的总数
func (ss *myStopSign) DealTotal() uint32 {
	ss.rwmutex.Lock()
	defer ss.rwmutex.Unlock()
	var total uint32
	for _,v := range ss.dealCountMap{
		total += v
	}
	return total
}
//概要
func (ss *myStopSign)  Summary() string {
	if ss.signed {
		return fmt.Sprintf("signed: true, dealCount: %v", ss.dealCountMap)
	} else {
		return "signed: false"
	}
}

