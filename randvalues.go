package testutils

import (
	"math/rand"
	"runtime"
	"sync"
)

type RandValues []interface{}

func NewRandValues() *RandValues {
	rv := make([]interface{}, 0, 1)
	return (*RandValues)(&rv)
}

func (rv *RandValues) Len() int {
	return len(*rv)
}

func (rv *RandValues) Put(v ...interface{}) *RandValues {
	*rv = append(*rv, v...)
	return rv
}

func (rv *RandValues) GetAll() []interface{} {
	return []interface{}(*rv)
}

// end is non-inclusive
func fasttransfer(start, end int, transfer func(i int) bool) bool {
	var wg sync.WaitGroup
	var failed bool

	count := end - start
	numgoroutines := runtime.GOMAXPROCS(-1)
	chunksize := (count + numgoroutines - 1) / numgoroutines

	routine := func(s, e int) {
		for i := s; i < e; i++ {
			if !transfer(i) {
				failed = true
				break
			}
		}
		wg.Done()
	}

	for i := start; i < end; i += chunksize {
		s := i
		e := i + chunksize
		if e >= end {
			e = end
		}
		wg.Add(1)
		go routine(s, e)
	}
	wg.Wait()

	return !failed
}

func (rv *RandValues) Clone() *RandValues {
	newrv := make([]interface{}, len(*rv))

	transfer := func(i int) bool {
		newrv[i] = (*rv)[i]
		return true
	}

	fasttransfer(0, len(*rv), transfer)

	return (*RandValues)(&newrv)
}

func (rv *RandValues) GetAllInt32() []int32 {
	vals := rv.GetAll()
	arr := make([]int32, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(int32)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllInt64() []int64 {
	vals := rv.GetAll()
	arr := make([]int64, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(int64)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllUint32() []uint32 {
	vals := rv.GetAll()
	arr := make([]uint32, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(uint32)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllUint64() []uint64 {
	vals := rv.GetAll()
	arr := make([]uint64, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(uint64)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllFloat32() []float32 {
	vals := rv.GetAll()
	arr := make([]float32, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(float32)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllFloat64() []float64 {
	vals := rv.GetAll()
	arr := make([]float64, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(float64)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllStrings() []string {
	vals := rv.GetAll()
	arr := make([]string, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(string)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) GetAllBool() []bool {
	vals := rv.GetAll()
	arr := make([]bool, len(*rv))

	transfer := func(i int) bool {
		var ok bool
		arr[i], ok = vals[i].(bool)
		return ok
	}

	if !fasttransfer(0, len(vals), transfer) {
		return nil
	}

	return arr
}

func (rv *RandValues) Shuffle() *RandValues {
	rand.Shuffle(len(*rv), func(i, j int) {
		(*rv)[i], (*rv)[j] = (*rv)[j], (*rv)[i]
	})
	return rv
}

func (rv *RandValues) Clear() *RandValues {
	*rv = (*rv)[0:0]
	return rv
}

func (rv *RandValues) AddConsecutiveInt32(start int32, count int) *RandValues {
	for i := start; i < start+int32(count); i++ {
		*rv = append(*rv, i)
	}
	return rv
}

func (rv *RandValues) AddConsecutiveInt64(start int64, count int) *RandValues {
	for i := start; i < start+int64(count); i++ {
		*rv = append(*rv, i)
	}
	return rv
}

func (rv *RandValues) AddSparseInt32(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Int31())
	}
	return rv
}

func (rv *RandValues) AddSparseInt64(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Int63())
	}
	return rv
}

func (rv *RandValues) AddUniformInt32(count int, mod int32) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Int31n(mod))
	}
	return rv
}

func (rv *RandValues) AddUniformInt64(count int, mod int64) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Int63n(mod))
	}
	return rv
}

func (rv *RandValues) AddSparseUint32(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Uint32())
	}
	return rv
}

func (rv *RandValues) AddSparseUint64(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Uint64())
	}
	return rv
}

func (rv *RandValues) AddSparseFloat32(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Float32())
	}
	return rv
}

func (rv *RandValues) AddSparseFloat64(count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, rand.Float64())
	}
	return rv
}

func (rv *RandValues) AddIdenticalBool(value bool, count int) *RandValues {
	for i := 0; i < count; i++ {
		*rv = append(*rv, value)
	}
	return rv
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (rv *RandValues) AddStringOfLength(count, length int) *RandValues {
	strs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		b := make([]byte, length)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		strs[i] = string(b)
	}
	*rv = append(*rv, strs...)
	return rv
}
