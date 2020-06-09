package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

// 使用一个工厂函数,每次传个type,新建一个对应类型的实例并返回
// 就是把创建实例这件事统一在一个函数中去做,方便管理吧.
// 且创建的是有相关性的实例,比如仅仅是type不一样等特点
func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ )
	case DiskStorage:
		return newDiskStorage( /*...*/ )
	default:
		return newTempStorage( /*...*/ )
	}
}
