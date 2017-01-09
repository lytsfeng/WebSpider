package scheduler

type SchedSummary interface {
	String() string 				//获取摘要的一般表示
	Detail() string 				//获取摘要的详细信息
	Same(other SchedSummary) bool	//判断是否与另一份信息相同
}
