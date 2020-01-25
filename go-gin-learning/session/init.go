package session

import "fmt"

var(
	sessionMgr SessionMgr
)
func Init(provider string,addr string,options ...string) (err error) {
	switch provider {
	case"memory":
		sessionMgr = NewMemorySessionMgr()
	case"redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Println("Unknown")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}
