package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"go-blog/util"
	"time"
)

func JwtAuth(ctx *context.Context) {
	token := ctx.GetCookie("token")
	if len(token) == 0 {
		logs.Info("token is empty")
		generateToken, err := util.GenerateToken("admin", "121233")
		if err == nil {
			ctx.SetCookie("token", generateToken, 5, "/")
		}
	} else {
		logs.Info("token is :", token)

		parseToken, err := util.ParseToken(token)
		if nil != err {
			logs.Info("parseToken error")
		} else if time.Now().Unix() > parseToken.ExpiresAt {
			logs.Info("parseToken Expired", parseToken)
		}
	}
}
