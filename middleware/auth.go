package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/commen/utils"
	"our_blog/model/dao"
)

//

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/user/login" || c.Request.URL.Path == "/user/register" {
			return
		}

		//获取前端传来的 acc token
		accesstoken := c.Request.Header.Get("Authorization")
		log.Println("acctoken: ", accesstoken)
		//若没有token，则为游客状态
		if accesstoken == "" {
			result.Error(c, result.UnLoginStatus)
			c.Abort()
			return
		}

		//判断acc token是否过期
		timeOut, err := utils.ValidToken(accesstoken)
		if err == utils.TokenStatusErr {
			result.Error(c, result.ParseTokenErrStatus)
			c.Abort()
			return
		}
		if err != nil {
			result.Error(c, result.ServerErrStatus)
			c.Abort()
			return
		}

		//acc token过期，寻找对应的ref token
		if timeOut {
			log.Println("acc token expired")
			refreshtoken, err := dao.NewTokenDao().GetKey(accesstoken)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}

			timeOut, err = utils.ValidToken(refreshtoken)
			if err != nil {
				result.Error(c, result.ParseTokenErrStatus)
				c.Abort()
				return
			}
			//ref token 过期，需要重新登录
			if timeOut {
				log.Println("ref token expired")
				err = dao.NewTokenDao().DelKey(accesstoken)
				if err != nil {
					result.Error(c, result.ServerErrStatus)
					c.Abort()
					return
				}
				result.Error(c, result.TokenExiredStatus)
				c.Abort()
				return
			}

			//ref token没有过期，重新设置acc token，ref token
			log.Println("ref token valid")
			err = dao.NewTokenDao().DelKey(accesstoken)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}
			userid, err := utils.GetUserIdFromToken(refreshtoken)
			if err != nil {
				result.Error(c, result.ParseTokenErrStatus)
				c.Abort()
				return
			}
			accesstoken, err := utils.CreateAccessToken(userid)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}

			//更新acc token,并把userid放入上下文
			err = dao.NewTokenDao().SetKey(accesstoken, refreshtoken)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}
			c.Header("newtoken", accesstoken)
			c.Set("userid", userid)
			c.Next()
			return
		}

		//acctoken没有过期
		log.Println("acc token valid")
		userid, err := utils.GetUserIdFromToken(accesstoken)
		if err != nil {
			result.Error(c, result.ParseTokenErrStatus)
			c.Abort()
			return
		}
		c.Set("userid", userid)
		c.Next()
		return

	}
}
