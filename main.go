package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/iamxuwenjin/blog/pkg/logging"
	"github.com/iamxuwenjin/blog/pkg/setting"
	"github.com/iamxuwenjin/blog/routers"
	"syscall"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.Error("server err: %v", err)
	}
}
