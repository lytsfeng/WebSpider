package middleware

import (
	"WebSpider/base"
	"errors"
	"sync"
	"fmt"
)

// 被用来表示通道管理器的状态的类型。
type ChannelManagerStatus uint8

const (
	CHANNEL_MANAGER_STATUS_UNINITIALIZED ChannelManagerStatus = 0 // 未初始化状态。
	CHANNEL_MANAGER_STATUS_INITIALIZED   ChannelManagerStatus = 1 // 已初始化状态。
	CHANNEL_MANAGER_STATUS_CLOSED        ChannelManagerStatus = 2 // 已关闭状态。
)

// 表示状态代码与状态名称之间的映射关系的字典。
var statusNameMap = map[ChannelManagerStatus]string{
	CHANNEL_MANAGER_STATUS_UNINITIALIZED: "uninitialized",
	CHANNEL_MANAGER_STATUS_INITIALIZED:   "initialized",
	CHANNEL_MANAGER_STATUS_CLOSED:        "closed",
}

//通道管理器接口
type ChannelManager interface {
	/**
	初始化通道管理器
	channelLen 各类通道的初始长度
	reset 是否重新初始化管道
	*/
	Init(channelLen uint, reset bool) bool
	//关闭通信管理器
	Close() bool
	//获取请求通道
	ReqChan() (chan base.Request, error)
	//获取响应通道
	RespChan() (chan base.Response, error)
	//过去条目通道
	ItemChan() (chan base.Item, error)
	//获取错误通道
	ErrorChan() (chan error, error)
	//获取通道长度
	ChannelLen() uint
	//获取通道状态
	Status() ChannelManagerStatus
	//获取摘要信息
	Summary() string
}

type myChannelManager struct {
	rwMutex sync.RWMutex			//读写锁
	channelLen uint                 //通道长度
	reqChan    chan base.Request    //请求通道
	respChan   chan base.Response   //效应通道
	itemChan   chan base.Item       //条目通道
	errChan    chan error           //错误通道
	status     ChannelManagerStatus //通道管理器状态
}

func (chaman *myChannelManager) Init(channelLen uint, reset bool) bool {

	if channelLen == 0 {
		panic(errors.New("The channle length is invalid!"))
	}
	chaman.rwMutex.Lock()			//加锁
	defer chaman.rwMutex.Unlock()	//解锁
	if chaman.status == CHANNEL_MANAGER_STATUS_INITIALIZED && !reset {
		return false
	}
	chaman.channelLen = channelLen
	chaman.reqChan = make(chan base.Request, channelLen)
	chaman.respChan = make(chan base.Response, channelLen)
	chaman.itemChan = make(chan base.Item, channelLen)
	chaman.errChan = make(chan error, channelLen)
	chaman.status = CHANNEL_MANAGER_STATUS_INITIALIZED
	return true
}
/**
	检查状态，在获取通道的时候，通道管理器应处于已经初始化的状态
	如果未处于初始化状态讲放回 非nil
 */
func (chaman *myChannelManager)checkStatus()  error{
	if chaman.status == CHANNEL_MANAGER_STATUS_INITIALIZED{
		return  nil
	}
	statusName,ok := statusNameMap[chaman.status]
	if !ok{
		statusName = fmt.Sprintf("%d",chaman.status)
	}
	errMsg :=  fmt.Sprintf("The ubdesirable status od channel manager:%s\n", statusName)
	return errors.New(errMsg)
}

//获取请求通道
func (chaman *myChannelManager) ReqChan() (chan base.Request, error){
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	if err:= chaman.checkStatus();err != nil{
		return nil,err
	}
	return chaman.reqChan,nil
}
//获取响应通道
func (chaman *myChannelManager) RespChan() (chan base.Response, error){
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	if err:= chaman.checkStatus();err != nil{
		return nil,err
	}
	return chaman.respChan,nil
}
//过去条目通道
func (chaman *myChannelManager) ItemChan() (chan base.Item, error){
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	if err:= chaman.checkStatus();err != nil{
		return nil,err
	}
	return chaman.itemChan,nil
}
//获取错误通道
func (chaman *myChannelManager) ErrorChan() (chan error, error){
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	if err:= chaman.checkStatus();err != nil{
		return nil,err
	}
	return chaman.errChan,nil
}
//获取通道长度
func (chaman *myChannelManager) ChannelLen() uint{
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	return chaman.channelLen
}

func (chaman * myChannelManager)Close() bool {
	chaman.rwMutex.Lock()
	defer chaman.rwMutex.Unlock()
	if chaman.status != CHANNEL_MANAGER_STATUS_INITIALIZED{
		return false
	}
	close(chaman.reqChan)
	close(chaman.respChan)
	close(chaman.errChan)
	close(chaman.itemChan)
	chaman.status = CHANNEL_MANAGER_STATUS_UNINITIALIZED
	return true
}

func (chanman *myChannelManager) Status() ChannelManagerStatus {
	return chanman.status
}

var chanmanSummaryTemplate = "status: %s, " +
		"requestChannel: %d/%d, " +
		"responseChannel: %d/%d, " +
		"itemChannel: %d/%d, " +
		"errorChannel: %d/%d"

func (chanman *myChannelManager) Summary() string {
	summary := fmt.Sprintf(chanmanSummaryTemplate,
		statusNameMap[chanman.status],
		len(chanman.reqChan), cap(chanman.reqChan),
		len(chanman.respChan), cap(chanman.respChan),
		len(chanman.itemChan), cap(chanman.itemChan),
		len(chanman.errChan), cap(chanman.errChan))
	return summary
}

func NewChannelManager(channelLen uint) ChannelManager {
	if channelLen == 0 {
		channelLen = 16
	}
	chanmam := &myChannelManager{}
	chanmam.Init(channelLen, true)
	return chanmam
}
