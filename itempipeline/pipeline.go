package itempipeline

import "WebSpider/base"

//条目处理管道的接口类型
type ItemPipeline interface {
	//发送条目
	Send(item base.Item) []error
	/**
	FailFast返回一个bool值，该值表示当前条目处理管道是否是快速失败
	这里的快速失败指：只要对每个条目的处理流程在某一个步骤上出错，那么
	条目管理通道就会忽略掉后续的所有处理步骤并报告错误
	*/
	FailFase() bool
	//设置是否快速失败
	SetFailFast(failfast bool)
	/**
	获取已发送、已接受和已处理的条目计数值
	作为结果值的切片总会有三个元素，这三个值分别代表前面说的三个计数
	*/
	Count() []uint64
	//获取正在被处理的条目的数量
	ProcessingNumber() uint64
	//获取摘要信息
	Summary() string
}
