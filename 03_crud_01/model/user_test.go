package model

import (
	"fmt"
	"testing"
)

// TestMain 函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始............")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("111111")
	//通过t.Run() 来执行子测试函数
	//t.Run("测试添加",testAddUser)
	t.Run("测试查询",testUser_GetUserById)
	t.Run("测试查询所有",testGetUsers)
}

//函数名不是以Test开头，默认不执行，但是可以设置成一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	//user := &User{}
	//user.AddUser()
	//user.AddUser2()
}

func testUser_GetUserById(t *testing.T) {
	fmt.Println("查询到的数据为：")
	user := User{
		ID: 1,
	}
	//调用获取的方法
	u,_ := user.GetUserById()
	fmt.Println(*u)
}

//测试获取所有
func testGetUsers(t *testing.T) {
	fmt.Println("所有记录：")
	user := &User{}
	us,_ := user.GetUsers()
	//遍历切片
	for k,v := range us{
		fmt.Printf("第%v个用户是%v\n",k+1,v)
	}
}

