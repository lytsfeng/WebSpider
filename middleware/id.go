package middleware

type IdGenerator interface {
	GetUint32() uint32   // 获取一个uint32类型的id
}

