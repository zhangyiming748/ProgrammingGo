package crud_mysql

import (
	"ProgrammingGo/log"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Insert(name, passwd string) {
	//使用工具获取数据库连接
	db := InitDB()
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		//事务开启失败，直接panic
		panic(err)
	}
	//准备SQL语句
	sql := "insert into t_testuser (`name`, `password`) values (?, ?)"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(name, passwd)
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	//提交事务
	if err := tx.Commit(); err != nil {
		log.Error.Panicln("提交事务出错")
	}
	//返回插入记录的id
	if last, err := result.LastInsertId(); err == nil {
		log.Debug.Printf("插入记录的id:%d", last)
	}

}
func SelectAll() {
	//使用工具获取数据库连接
	db := InitDB()
	//准备SQL语句
	sql := "select * from t_testuser "
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	var users []User
	for rows.Next() {
		var id int
		var name, password string
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			//读取结果集失败
			panic(err)
		}
		var user User
		user.SetId(id)
		user.SetName(name)
		user.SetPassword(password)
		users = append(users, user)
	}
	fmt.Println(users)
}
func SelectByName(name string) {
	//使用工具获取数据库连接
	db := InitDB()
	//准备SQL语句
	sql := "select * from t_testuser where `name` = ?"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(name)
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	var users []User
	for rows.Next() {
		var id int
		var name, password string
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			//读取结果集失败
			panic(err)
		}
		var user User
		user.SetId(id)
		user.SetName(name)
		user.SetPassword(password)
		users = append(users, user)
	}
	fmt.Println(users)
}
func Update() {
	//使用工具获取数据库连接
	db := InitDB()
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		//事务开启失败，直接panic
		panic(err)
	}
	//准备SQL语句
	sql := "update t_testuser set `password` = ? where `id` = ?"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("789", 1)
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	//提交事务
	if err := tx.Commit(); err != nil {
		log.Error.Panicln("提交事务出错")
	}
}
func DeleteByName(name string) {
	//使用工具获取数据库连接
	db := InitDB()
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		//事务开启失败，直接panic
		panic(err)
	}
	//准备SQL语句
	sql := "delete from t_testuser where `name` = ?"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(name)
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	//提交事务
	if err := tx.Commit(); err != nil {
		log.Error.Panicln("提交事务出错")
	}
}
func isDuplicate(name string) bool {
	//使用工具获取数据库连接
	//准备SQL语句
	db := InitDB()
	sql := "select * from t_testuser where `name` = ?"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(name)
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	var users []User
	for rows.Next() {
		var id int
		var name, password string
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			//读取结果集失败
			panic(err)
		}
		var user User
		user.SetId(id)
		user.SetName(name)
		user.SetPassword(password)
		users = append(users, user)
	}
	if len(users) != 0 {
		return true
	}
	return false
}
