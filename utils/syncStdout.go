package utils

import (
	"fmt"
	"github.com/kataras/iris"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"strings"
)

//通过管道同步获取日志的函数
func SyncLog(ctx iris.Context, reader io.ReadCloser) {

	buf := make([]byte, 1024, 1024)
	for {
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			//fmt.Print(string(outputByte))
			//ctx.Write(outputByte)
			b, _ := simplifiedchinese.GBK.NewDecoder().Bytes(outputByte)
			if _, err := ctx.WriteString(strings.ReplaceAll(string(b), "\n", "<br>")); err != nil {
				fmt.Println(err)
			}

			ctx.ResponseWriter().Flush()
			//ctx.Text(string(outputByte))

		}
		if err != nil {
			//读到结尾
			if err == io.EOF {
				err = nil
			}
		}
	}
}
