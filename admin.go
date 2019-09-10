package main

import (
	"adminMonitor/adminPrivileges"
	"adminMonitor/config"
	"adminMonitor/iis"
	"adminMonitor/vfs"
	"github.com/kardianos/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"log"
	"os"
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

	api.PartyFunc("/admin", func(admin router.Party) {
		admin.Post("/add", adminPrivileges.AddAdmin)
		admin.Post("/del", adminPrivileges.DelAdmin)
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
	addr := config.Conf.Get("app.addr").(string)
	app.Run(iris.Addr(addr))
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func init() {
	file := config.GetCurrentDirectory() + "/admin_log.txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	//log.SetPrefix("[qSkipTool]")
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {

	svcConfig := &service.Config{
		Name:        "septnetAdmin",        //服务显示名称
		DisplayName: "septnetAdmin",        //服务名称
		Description: "七天网络开发环境调试权限,IIS管理等", //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Println(err)
	}

	if err != nil {

		log.Println(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			if err := s.Install(); err != nil {
				log.Println("服务安装成功")
				return
			}

		}

		if os.Args[1] == "remove" {
			if err := s.Uninstall(); err != nil {
				log.Println("服务卸载成功")
				return
			}

		}
	}

	err = s.Run()
	if err != nil {
		log.Println(err)
	}

}
