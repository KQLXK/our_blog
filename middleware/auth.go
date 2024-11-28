package middleware

import (
	"github.com/gin-gonic/gin"
	"our_blog/commen/result"
	"our_blog/dao"
	"our_blog/utils"
)

//

func Auth(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		accesstoken := c.Request.Header.Get("Authorization")
		//若没有token，则为游客状态
		if accesstoken == "" {
			c.Set("userid", "")
			c.Next()
			return
		}

		//判断acc token是否过期
		timeOut, err := utils.ValidToken(accesstoken)
		if err != nil {
			result.Error(c, result.ServerErrStatus)
			c.Abort()
			return
		}

		//acc token过期，寻找对应的ref token
		if timeOut {
			refershtoken, err := dao.TokenDao.GetKey(accesstoken)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}

			timeOut, err = utils.ValidToken(refershtoken)

			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}
			//ref token 过期，需要重新登录
			if timeOut {
				dao.TokenDao.DelKey(accesstoken)
				result.Error(c, result.TokenExiredStatus)
			}

			//ref token没有过期，重新设置acc token，ref token
			dao.TokenDao.DelKey(accesstoken)
			userid, err := utils.GetUserIdFromToken(refershtoken)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}
			accesstoken, err := utils.CreateAccessToken(userid)
			if err != nil {
				result.Error(c, result.ServerErrStatus)
				c.Abort()
				return
			}

			//更新acc token
			dao.TokenDao.SetKey(accesstoken, refershtoken)

		}
	}
}
