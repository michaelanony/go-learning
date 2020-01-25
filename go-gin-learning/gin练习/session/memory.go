package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	//save kv
	data map[string] interface{}
	rwLock sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{},16),
	}
	return s
}

func (m *MemorySession)Set(key string,value interface{}) (err error) {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//Set value
	m.data[key] = value
	return
}

func (m *MemorySession)Get(key string) (value interface{},err error) {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//Get value
	value,ok := m.data[key]
	if !ok{
		err =errors.New("Key not exists")
	}
	return
}

func (m *MemorySession)Del(key string)(err error)  {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data,key)
	return
}

func (m *MemorySession)Save(key string)(err error)  {
	return
}