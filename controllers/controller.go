package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"myblog/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

// Post 处理注册业务逻辑
func (c *RegisterController) Post() {
	//username: test
	//password: 123456
	//repassword: 123456

	//思路
	//1.从前端获取表单信息
	//获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)

	//2.对数据进行判断
	//注册之前先判断该用户是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "用户名已经存在",
		}
		c.ServeJSON()
		return
	}

	//3.将密码进行加密
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后:", password)

	//4.将数据插入数据库中
	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
	}
	_, err := models.InsertUser(user)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "注册失败",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "注册成功",
		}
	}
	c.ServeJSON()
}
