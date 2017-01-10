package middleware

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