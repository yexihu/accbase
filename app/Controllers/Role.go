package Controllers

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Adapter = gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", true)
var Enforcer = casbin.NewEnforcer("./casbin/model.conf", Adapter);

func GetAllRoles(c *gin.Context){
	allRoles := Enforcer.GetAllRoles()
	c.JSON(200, gin.H{
		"status": true,
		"data": allRoles,
		"message":"查询成功",
	})
}

type PoliceUserRoleObjectAction struct {
	Police string `json:"police"`
	User string `json:"user"`
	Role string `json:"role"`
	Object string `json:"object"`
	Action string `json:"action"`
}

func AddRoleForUser(c *gin.Context){
	var msg string
	var json PoliceUserRoleObjectAction
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}
	fmt.Println(json.User, json.Role)
	added := Enforcer.AddRoleForUser(json.User, json.Role)

	if added {
		msg = "添加成功！"
	} else {
		msg = "添加失败！"
	}

	c.JSON(200, gin.H{
		"status": true,
		"message":msg,
	})
}


func AddNamedPolicy(c *gin.Context){
	var msg string
	var json PoliceUserRoleObjectAction
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return
	}

	fmt.Println(json.Police, json.User, json.Object, json.Action)

	added := Enforcer.AddNamedPolicy(json.Police, json.User, json.Object, json.Action)

	if !added {
		msg = "添加成功！"
	} else {
		msg = "添加失败！"
	}

	c.JSON(200, gin.H{
		"status": true,
		"message":msg,
	})
}
