package analzyer

type AnalyzerPool interface {
	Take() (Analyzer,error)    		// 获取一个解析器
	Return(analyzer Analyzer) error	// 回收一个解析器
	Total() uint32					// 解析器个数
	Used() uint32					//这在使用的解析器个数
}


