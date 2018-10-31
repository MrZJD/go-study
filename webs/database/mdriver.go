package main

import (
	"database/sql"
)

// database/sql/driver接触上定义了更高阶的方法 并实现了一个建议的conn pool
// type DB struct {
// 	driver   driver.Driver
// 	dsn      string
// 	mu       sync.Mutex
// 	freeConn []driver.Conn
// 	closed   boll
// }

// driver.Driver
// type Driver interface { // db驱动
// 	Open(name string) (Conn, error) // 返回一个数据连接 Conn只能用来进行一次go routine操作
// }

// driver.Conn
// type Conn inerface { // db连接
// 	Prepare(query string) (Stmt, error) // 执行sql的准备状态
// 	Close() error
// 	Begin() (Tx, error) // Tx代表事务处理
// }

// driver.Execer // Conn 可选择实现的接口
// type Execer interface {
// 	Exec(query string, args []Value) (Result, error)
// }

// driver.Stmt
// type Stmt interface {
// 	Close() error
// 	NumInput() int // 返回当前预留参数的个数
// 	Exec(args []Value) (Result, error) // 执行Prepare准备好的sql，传入参数执行update/insert等操作
// 	Query(args []Value) (Rows, error) // 执行Prepare准备好的sql，传入需要的参数执行select操作
// }

// driver.Tx
// type Tx interface {
// 	Commit() error // 递交事务
// 	Rollback() error // 回滚事务
// }

// driver.Result
// type Result interface {
// 	LastInsertId() (int64, error) // 执行insert操作得到的自增id号
// 	RowsAffected() (int64, error) // query操作影响的数目条数
// }

// driver.Rows
// type Rows interface {
// 	Columns() []string // 查询表的字段信息
// 	Close() error
// 	Next(dest []Value) error // 返回下一条数据 将值赋给dest dest里面的元素必须是driver.Value中的值(string外) // io.EOF (end of Rows)
// }

// driver.RowsAffected
// type RowsAffected int64

// func (RowsAffected) LastInsertId() (int64, error)
// func (v RowsAffected) RowsAffected() (int64, error)

// driver.Value
// type Value interface{} // int64 float64 bool []byte string time.Time nil

// driver.ValueConverter
// type ValueConverter interface { // 将普通的值转为driver.Value的接口
// 	ConvertValue (v interface{}) (Value, error)
// }

// driver.Valuer
// type Valuer interface {
// 	Value() (Value, error)
// }

func init() { // 1. 注册数据库驱动
	sql.Register("msql", driver)
}
