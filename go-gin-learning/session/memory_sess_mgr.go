package session

import (
	"errors"
	"github.com/satori/go.uuid"
	"sync"
)

//定义memorySession对象（字段：sessionId、存kv的map，读写锁）
//构造函数，为了获取对象
//SET()、GET()、DEL()、SAVE()

type MemorySessionMgr struct {
	rwLock sync.RWMutex
	sessionMap map[string]Session
}

func NewMemorySessionMgr() *MemorySessionMgr  {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session,1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(add string,options ...string) (err error){
	return
}
func (s *MemorySessionMgr) CreateSession()(session Session,err error)  {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	//用uuid作为sessionid
	id := uuid.NewV4()
	//转string
	sessionId := id.String()
	//创建session
	session = NewMemorySession(sessionId)
	//加入到大map
	s.sessionMap[sessionId] = session
	return
}
func (s *MemorySessionMgr) Get(sessionId string)(session Session,err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	session,ok := s.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exists")
		return
	}
	return
}