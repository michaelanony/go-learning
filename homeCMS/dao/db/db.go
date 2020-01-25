package db

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var(
	MysqlDB *sqlx.DB
	RedisPool *redis.Pool
)

func Init(mysqlDns,RedisDns string) (err error) {
	MysqlDB, err = sqlx.Open("mysql", mysqlDns)
	if err!=nil{
		return err
	}
	if err=MysqlDB.Ping();err!=nil{
		return nil
	}
	MysqlDB.SetMaxIdleConns(100)
	MysqlDB.SetMaxOpenConns(16)

	return nil


}