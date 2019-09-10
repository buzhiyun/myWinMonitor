package adminPrivileges

import (
	"adminMonitor/config"
	"adminMonitor/modules"
	"adminMonitor/utils"
	"bytes"
	"github.com/kataras/iris"
	"github.com/pelletier/go-toml"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os/exec"
	"time"
	//"encoding/json"
)

var users = config.Conf.Get("user").(*toml.Tree)

//添加管理员
func AddAdmin(ctx iris.Context) {
	adminUser := new(modules.AdminUser)
	if err := ctx.ReadJSON(&adminUser); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(modules.ApiResource(false, nil, "请求参数错误"))
	} else {
		//检查用户名是否合法
		userInfo := users.Get(adminUser.Username)
		if userInfo != nil {
			setAdmin(ctx, adminUser.Username, true)

			//发送通知邮件

		} else {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.JSON(modules.ApiResource(false, nil, "请确认用户名"))
			go delAdmin(adminUser.Username, adminUser.EnableHours)
			//发送邮件通知
			sendEmail(userInfo.(*toml.Tree).Get("name").(string),
				userInfo.(*toml.Tree).Get("email").(string),
				"cc环境权限授权", "您的cc环境帐号 admin 权限已经授予，请小心使用。")
		}
	}
}

func DelAdmin(ctx iris.Context) {
	adminUser := new(modules.AdminUser)
	if err := ctx.ReadJSON(&adminUser); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(modules.ApiResource(false, nil, "请求参数错误"))
	} else {
		//检查用户名是否合法
		userInfo := users.Get(adminUser.Username)
		if userInfo != nil {
			setAdmin(ctx, adminUser.Username, false)
		} else {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.JSON(modules.ApiResource(false, nil, "请确认用户名"))
		}
	}
}

func sendEmail(userEmail string, username string, subject string, message string) {

	mailstr1 := `<div><div style="font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px;"><span style="font-size: 10.5pt; line-height: 1.5; background-color: window;">`

	mailstr2 := `，您好：</span></div><div style="font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px;"><span style="font-size: 10.5pt; line-height: 1.5; background-color: window;"><br></span></div><div style="font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px;"><span style="font-family: ''; font-size: 10.5pt; line-height: 1.5; background-color: window;">&nbsp; &nbsp;&nbsp;</span><span style="font-family: ''; font-size: 10.5pt; line-height: 1.5; background-color: window;">&nbsp; &nbsp;&nbsp;</span><span style="background-color: window; font-size: 10.5pt; line-height: 1.5;">`

	mailstr3 := `</span></div><div style="font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px;"><br></div><hr color="#b5c4df" size="1" align="left" style="box-sizing: border-box; font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px; width: 210px; height: 1px;"><div style="font-family: &quot;Microsoft YaHei UI&quot;; line-height: 21px;"><div style="position: static !important; margin: 10px; font-size: 10pt;"><div class="MsoNormal" align="left" style="font-size: 10.5pt; line-height: normal; font-family: Calibri, sans-serif; text-align: justify; margin: 0cm 0cm 0.0001pt;"><span style="font-size: 10.5pt; line-height: 1.5; background-color: window;">Best regards,</span></div><p class="MsoNormal" style="margin: 0cm 0cm 0.0001pt; font-size: 10.5pt; line-height: normal; font-family: Calibri, sans-serif; text-align: justify;"><span lang="EN-US" style="color: rgb(31, 73, 125);">&nbsp;</span><span lang="EN-US"><o:p></o:p></span></p><p class="MsoNormal" style="margin: 0cm 0cm 0.0001pt; font-size: 10.5pt; line-height: normal; font-family: Calibri, sans-serif; text-align: justify;"><b><span style="font-size: 14pt; font-family: 华文新魏; color: rgb(64, 64, 64);">监控中心</span></b><span lang="EN-US"><o:p></o:p></span></p><p class="MsoNormal" style="margin: 0cm 0cm 0.0001pt; font-family: Verdana; font-size: 14px; line-height: normal; text-align: justify;"><font color="#ee9a1e" face="微软雅黑, sans-serif" size="2"><span style="line-height: 19px;">七天网络</span></font></p></div></div></div><div><sign signid="0"><div style="font-size:14px;font-family:Verdana;color:#000;">
	</div></sign></div><div>&nbsp;</div><div><includetail><!--<![endif]--></includetail></div>`

	htmlmailstr := mailstr1 + username + mailstr2 + message + mailstr3
	utils.SendMail(userEmail, "周杨<zhouyang@7net.cc>", subject, htmlmailstr, "")

}

//异步调度到移除用户
func delAdmin(username string, hours float32) {
	time.Sleep(time.Duration((int)(hours*3600)) * time.Second)

	cmd := exec.Command("net", "localgroup", "administrators", username, "/delete")
	if err := cmd.Start(); err != nil {
		return
	}
	if err := cmd.Wait(); err != nil {
		return
	}
}

//设置用户到administrator组或者移出该组
func setAdmin(ctx iris.Context, userName string, enable bool) {
	//cmd := "net  administrators %s /delete"

	oprStr := "/delete"

	if enable {
		oprStr = "/add"
	}

	cmd := exec.Command("net", "localgroup", "administrators", userName, oprStr)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		return
	}
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	err = cmd.Wait()

	b, _ := simplifiedchinese.GBK.NewDecoder().Bytes(out.Bytes())
	ctx.JSON(modules.ApiResource(true, nil, string(b)))

}
