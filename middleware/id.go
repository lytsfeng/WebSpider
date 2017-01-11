package middleware

import (
	"sync"
	"math"
)

type IdGenerator interface {
	GetUint32() uint32 // 获取一个uint32类型的id
}
type IdGenerator2 interface {
	GetUint64() uint64
}

type cyclicIdGenerator struct {
	sn    uint32     // 当前的ＩＤ
	ended bool       //是否已经是其类型所能表示的最大值
	mutex sync.Mutex // 互斥锁
}

type cyclicIdGeberator2 struct{
	base cyclicIdGenerator
	cycleCount uint64
}


func NewIdGenerator() IdGenerator {
	return &cyclicIdGenerator{}
}
func NoewIdGenerator2() IdGenerator2 {
	return &cyclicIdGeberator2{}
}

func (gen *cyclicIdGenerator) GetUint32() uint32  {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.ended{
		gen.sn = 0;
		gen.ended = false
		return gen.sn
	}
	id := gen.sn
	if id < math.MaxUint32{
		gen.sn ++
	}else {
		gen.ended = true
	}
	return id
}

func (gen *cyclicIdGeberator2) GetUint64() uint64{
	var id64 uint64
	if gen.cycleCount % 2 == 1{     // uint64  　是　uint32的两倍
		id64 += math.MaxUint32
	}
	id32 := gen.base.GetUint32()

	if id32 == math.MaxUint32{
		gen.cycleCount ++
	}
	id64 += uint64(id32)
	return id64
}