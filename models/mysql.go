package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	//注册model
	orm.RegisterModel(&User{})
	//自动建表
	orm.RunSyncdb("default",false,true)
}
// 定义User模型，绑定users表结构, 其实就是用来保存sql查询结果。
type User struct {
	Id int
	Username string
	Password string
}
//一对一 模型
type Age struct {
	Id int
	Age int
}
type Users struct {
	Id int
	Username string
	Age *Age
}

type UserFrom struct {
	Id int `json:"-"`
	Name string `json:"username"`
	Phone string
}


// 定义User 模型绑定那个表？
func (u *User) TableName() string {
	// 返回mysql表名
	return "user"
}

//初始化函数，可以用来向orm注册model


// 根据id查询用户信息
func GetUserById(id int) *User {
	if  id == 0 {
		return  nil
	}

	// 创建orm对象, 后面都是通过orm对象操作数据库
	o := orm.NewOrm()

	// 初始化一个User模型对象
	user := User{}
	// 设置查询参数
	user.Id = id

	// 调用Read方法，根据user设置的参数，查询一条记录，结果保存到user结构体变量中
	// 默认是根据主键进行查询
	// 等价sql： SELECT `id`, `username`, `password` FROM `users` WHERE `id` = 1
	err := o.Read(&user)
	fmt.Println("o.read:",o.Read(&user))

	// 检测查询结果，
	if err == orm.ErrNoRows {
		// 找不到记录
		return nil
	} else if err == orm.ErrMissPK {
		// 找不到住建
		return nil
	}

	return &user
}
