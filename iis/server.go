package iis

import (
	"adminMonitor/utils"
	"fmt"
	"github.com/kataras/iris"
	"log"
	"os/exec"
)

func Start(ctx iris.Context) {
	iisreset(ctx, "start")
}

func Stop(ctx iris.Context) {
	iisreset(ctx, "stop")
}

func Restart(ctx iris.Context) {
	iisreset(ctx, "restart")
}

//iisreset 命令，操作iis状态
func iisreset(ctx iris.Context, setStatus string) {
	ctx.ContentType("text/html")
	ctx.Header("Transfer-Encoding", "chunked")
	cmd := exec.Command("iisreset.exe", "/"+setStatus)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return
	}
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}

	go utils.SyncLog(ctx, stdout)
	err = cmd.Wait()

}
