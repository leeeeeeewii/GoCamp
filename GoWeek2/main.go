package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

func main() {
	query, err := dao()
  if err != nil{
    fmt.Printf("dao error : %+v",err)
    return
  }
  fmt.Println("query : %s",query)
}

func dao() (query, error) {
  sql = "select * from DB"
  query, err := queryDB(sql) //fake function to simulate interacting with database
  if err != nil{
	  return nil, errors.Wrap(sql.ErrNoRows, "sql=%s, err=%+v", sql, err)
    }
  return query, nil
}
