package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)


type Page struct {
	Id int
	Website string
	Email string
}

func init()  {
	// 注册数据库
	orm.RegisterDataBase("default","mysql",
		"root@123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(Page))
}
func GetPage() Page {
	//rtn:=Page{Website: "hellobeego.com", Email: "model@beego.com"}
	//return rtn

	o:=orm.NewOrm()
  	//p:=Page{1}
  	p:=Page{Id:1}
	err:=o.Read(&p)
	fmt.Println(err)
	return p
}

func UpdatePage(){
	p:=Page{Id: 1, Email: "myemail2131",}
	o:=orm.NewOrm()
	o.Update(&p,"Email")

}
