package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/WangYiwei-oss/jdnotes-casbin-enforcer/src/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var grpcClient *grpc.ClientConn
var grpcClientDoOnce sync.Once

func init() {
	var err error
	grpcClientDoOnce.Do(func() {
		grpcClient, err = grpc.DialContext(context.Background(), "localhost:8083", grpc.WithInsecure())
		if err != nil {
			log.Fatalln("LoginRegisterService: GRPC客户端初始化失败", err)
		}
	})
}

type UserInfo struct {
	Exp      int    `json:"exp"`
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func DoEnforce(ctx *gin.Context) {
	base64Info := ctx.GetHeader("Userinfo")
	if base64Info == "" {
		ctx.AbortWithError(400, fmt.Errorf("未解析到用户信息"))
	}
	userInfoJson, err := base64.StdEncoding.DecodeString(base64Info)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	userInfo := UserInfo{}
	err = json.Unmarshal(userInfoJson, &userInfo)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	req := &services.EnforceRequest{
		UserName: userInfo.UserName,
		Route:    ctx.Request.RequestURI,
		Act:      ctx.Request.Method,
	}
	rsp := &services.EnforceResponse{}
	err = grpcClient.Invoke(ctx, "", req, rsp)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	if rsp.Pass == false {
		ctx.AbortWithError(500, fmt.Errorf("用户没有权限!"))
	} else {
		ctx.Next()
	}
}
