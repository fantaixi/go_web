package model

import (
	"fmt"
	"go_web/03_crud_01/utils"
)

type User struct {
	ID int
	Username string
	Password string
	Email string
}

//添加
//方法1
func (user *User) AddUser()  error{
	//1、sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2、预编译 Prepare 好处
	//避免通过引号组装拼接sql语句。避免sql注入带来的安全风险
	//可以多次执行的sql语句。
	inStmt,err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常：",err)
		return err
	}
	//3、执行
	_,err2 := inStmt.Exec("admin","123456","122@qq.com")
	if err2 != nil {
		fmt.Println("执行异常：",err)
		return err2
	}
	return nil
}

//添加
//方法2
func (user *User) AddUser2()  error{
	//1、sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2、执行
	_,err := utils.Db.Exec(sqlStr,"admin222","123456","122@qq.com")
	if err != nil {
		fmt.Println("执行异常：",err)
		return err
	}
	return nil
}

//根据ID查询
func (user *User) GetUserById() (*User, error) {
	//1、sql
	sqlStr := "select * from users where id = ?"
	//直接执行
	row := utils.Db.QueryRow(sqlStr,user.ID)
	//声明
	var id int
	var username string
	var password string
	var email string
	//调用row里面的Scan方法将结果映射到对应的字段
	err := row.Scan(&id,&username,&password,&email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID: id,
		Username: username,
		Password: password,
		Email: email,
	}
	return u, err
}

//获取所有数据
func (user *User) GetUsers() ([]*User, error) {
	//1、sql
	sqlStr := "select * from users"
	//2、执行
	rows,err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建User切片
	var users []*User

	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id,&username,&password,&email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID: id,
			Username: username,
			Password: password,
			Email: email,
		}
		users = append(users, u)
	}
	return users, nil
}