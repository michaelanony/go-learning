package session

type Session interface {
	Set(key string,value interface{}) error
	Get(Key string)(interface{},error)
	Del(key string)error
	Save()error
}
