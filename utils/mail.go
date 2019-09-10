package utils

import (
	"github.com/go-gomail/gomail"
	"log"
	"strings"
)

var smtpServer = struct {
	Host     string
	port     int
	Username string
	Passwd   string
}{
	Host:     "smtp.exmail.qq.com",
	port:     465,
	Username: "monitor@7net.cc",
	Passwd:   "Mon123",
}

func FormatAddress(addressStr string, m *gomail.Message) []string {
	addresses := strings.Split(addressStr, ",")
	addr := []string{}

	for _, value := range addresses {
		//log.Println("value:",value)
		thisAdd := strings.Split(value, "<")
		if len(thisAdd) > 1 {
			//log.Println("addr: ", strings.TrimRight(thisAdd[1], ">"))
			addr = append(addr, m.FormatAddress(strings.TrimRight(thisAdd[1], ">"), thisAdd[0]))
		} else {
			addr = append(addr, value)
		}
	}
	return addr
}

func SendMail(toUsers string, ccUsers string, subject string, htmlBody string, fileAttach string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", smtpServer.Username /*"发件人地址"*/, "发件人") // 发件人

	m.SetHeader("To", FormatAddress(toUsers, m)...) // 收件人

	//m.SetHeader("Cc",
	//	m.FormatAddress("xxxx@7net.cc", "收件人")) //抄送
	if len(ccUsers) > 0 {
		m.SetHeader("Cc", FormatAddress(ccUsers, m)...) //抄送
	}

	//m.SetHeader("Bcc",
	//	m.FormatAddress("xxxx@7net.cc", "收件人")) // 暗送

	m.SetHeader("Subject", subject) // 主题

	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/html", htmlBody) // 正文

	m.Attach(fileAttach) //添加附件

	d := gomail.NewDialer(smtpServer.Host, smtpServer.port, smtpServer.Username, smtpServer.Passwd) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return err
	}

	log.Println("done.发送成功")
	return nil

}

func main() {
	SendMail("天才周<zhouyang@7net.cc>,研发部<dev@7net.cc>,zhoumin@7net.cc", "", "测试邮件", "测试发送", "/Volumes/data/DEV/go/admin/src/myWinMonitor/utils/mail.go")

}
