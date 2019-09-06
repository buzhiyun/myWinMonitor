package main

import (
	"adminMonitor/iis"
	"adminMonitor/vfs"
	"github.com/kardianos/service"
	"github.com/kataras/iris"
	//"github.com/jander/golog/logger"
	"github.com/kataras/iris/core/router"
	//"os"
)

/**
 * 初始化 iris app
 * @method NewApp
 * @return  {[type]}  api      *iris.Application  [iris app]
 */

func newApp() (api *iris.Application) {
	api = iris.New()

	api.PartyFunc("/iis", func(server router.Party) {
		server.Get("/start", iis.Start)
		server.Get("/stop", iis.Stop)
		server.Get("/restart", iis.Restart)
	})
	api.PartyFunc("/7netVFS", func(vfsServer router.Party) {
		vfsServer.Get("/start", vfs.Start)
		vfsServer.Get("/stop", vfs.Stop)
		vfsServer.Get("/restart", vfs.Restart)
	})

	return
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 启动代码
	app := newApp()

	addr := ":8080"
	app.Run(iris.Addr(addr))
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {

	//svcConfig := &service.Config{
	//	Name:        "7netAdmin", //服务显示名称
	//	DisplayName: "开发环境权限管理服务", //服务名称
	//	Description: "七天网络开发环境调试权限,IIS管理等", //服务描述
	//}
	//
	//prg := &program{}
	//s, err := service.New(prg, svcConfig)
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//if err != nil {
	//
	//	logger.Fatal(err)
	//}
	//
	//if len(os.Args) > 1 {
	//	if os.Args[1] == "install" {
	//		if err := s.Install(); err != nil{
	//			logger.Println("服务安装成功")
	//			return
	//		}
	//
	//	}
	//
	//	if os.Args[1] == "remove" {
	//		if err := s.Uninstall() ;err != nil{
	//			logger.Println("服务卸载成功")
	//			return
	//		}
	//
	//	}
	//}
	//
	//err = s.Run()
	//if err != nil {
	//	logger.Error(err)
	//}

	app := newApp()

	addr := "127.0.0.1:12010"
	app.Run(iris.Addr(addr))
}
