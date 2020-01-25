package main

import "github.com/gin-gonic/gin"

//设计一个通用的session服务
//1、
//session接口设计：
//SET()、GET()、DEL()
//SAVE():session存储，redis的实现延迟加载
//2、
//sessionMgr接口设计
//INIT():初始化，加载redis地址
//createSession（）：创建一个新的session
//getSession（）：通过sessionId获取对应的session对象
//3、
//MemorySession设计：
//定义memorySession对象（字段：sessionId、存kv的map，读写锁）
//构造函数，为了获取对象
//SET()、GET()、DEL()、SAVE()
//4、
//定义RedisSession对象（字段：sessionId，存kv的map，读写锁，redis连接池，记录内存中map是否被修改的标记）
//SET()、GET()、DEL()、SAVE()
//5、redisSessionMgr设计
//定义redisSessionMgr对象（字段：redis地址、密码、连接池、读写锁、大Map）

func main() {
	r :=gin.Default()

	r.Run(":8000")
}
