package controller

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"zhoukai/configure"
	md "zhoukai/middleware"
	"zhoukai/model"
	ModelSys "zhoukai/model/sys"
	"zhoukai/service"
	"zhoukai/utils"
)

// @Tags 登录
// @Descriptions 用户登录
// @Param account body string true "用户名"
// @Param pass_word body string true "密码"
// @Router /api/v1/login [post]
func UserLogin(c *gin.Context)  {
	var sysUser ModelSys.SysUserLogin
	err := utils.Verify(&sysUser, c)
	if err == nil {
		// 验证账户、密码
		data, result, status := service.Login(sysUser)
		if result.RowsAffected > 0 {
			GenerateToken(c, data)
		}else {
			resultM := "密码不正确"
			if status == 1 {
				resultM = "用户名不存在"
			}
			c.JSON(http.StatusOK, model.Response{
				Code: configure.RequestUserPassValidationError,
				Message:  resultM,
			})
		}
	}
}

// token生成
func GenerateToken(c *gin.Context, user ModelSys.SysUserLogin)  {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := md.NewJWT()
	// 构造用户claims信息(负荷)
	claims := md.CustomClaims{
		user.ID,
		user.Account,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600 * 24), // 签名过期时间为一天
			Issuer:    "cn.start6.cn",                    // 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: configure.RequestAuthorizedFailed,
			Message: "token生成失败 ->> ",
		})
	}
	zap.L().Info("生成的token: ->> " + token)
	c.JSON(http.StatusOK, model.Response{
		Code: configure.RequestSuccess,
		Message: "登录成功",
		Data: token,
	})
}