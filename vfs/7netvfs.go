package vfs

import (
	"adminMonitor/utils"
	"fmt"
	"github.com/kataras/iris"
	"os/exec"
)

func Start(ctx iris.Context) {
	vfsStstus(ctx, "start")
}

func Stop(ctx iris.Context) {
	vfsStstus(ctx, "stop")
}

func Restart(ctx iris.Context) {
	vfsStstus(ctx, "restart")
}

func vfsStstus(ctx iris.Context, status string) {
	ctx.ContentType("text/html")
	ctx.Header("Transfer-Encoding", "chunked")
	cmd := exec.Command("net", status, "7netVFS")
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return
	}
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	go utils.SyncLog(ctx, stdout)
	err = cmd.Wait()
}
