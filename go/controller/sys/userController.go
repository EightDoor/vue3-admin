package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)

func UserList(c *gin.Context)  {
	var list []ModelSys.SysUser
	dates, result := ServiceSys.UserQuery(list, c)
	utils.R(dates, result, c)
}
func UserSinge(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data := ModelSys.SysUser{
		Id: id,
	}
	singData, result := ServiceSys.UserSinge(data)
	utils.R(singData, result, c)
}
func UserCreate(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	utils.Verify(&sysUser, c)
	//fmt.Println(sysUser)
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(data))
	//var result map[string]interface{}
	//err := json.Unmarshal(data, result)
	//if err != nil {
	//	fmt.Println(err.Error(), "err")
	//}
	//var sysUser ModelSys.SysUser
	//error := mapstructure.Decode(result, &sysUser)
	//if error != nil {
	//	fmt.Println(err.Error(), "errr")
	//}
	//fmt.Println(sysUser, "sysU")
	//model := &ModelSys.SysUser{}
	//utils.Verify(model, c)
}