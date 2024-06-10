package test

import (
	"fmt"
	"strings"
)
import "strconv"

var name = "2233"
var age = 1234

func Render() string {
	out := &strings.Builder{}
	out.WriteString("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Test</title>\n</head>\n<body>\n<!-- 导包指令 单独收集 没必要按顺序 -->\n")
	out.WriteString("\n<!-- 会放在主函数外面的 单独收集 没必要按顺序 -->\n")
	out.WriteString("\n<!-- 放在主函数里面的 out 隐式对象 必须按顺序来 -->\n")
	out.WriteString("你好")
	out.WriteString("\n<!-- 打印输出的 必须按顺序来 -->\n")
	out.WriteString(fmt.Sprintf("%v", name))
	out.WriteString("\n")
	out.WriteString(fmt.Sprintf("%v", strconv.Itoa(2233)))
	out.WriteString("\n")
	var names = []string{"2233", "4455"}
	for _, item := range names {
		out.WriteString("\n<h1>")
		out.WriteString(fmt.Sprintf("%v", item))
		out.WriteString("</h1>\n")
	}
	out.WriteString("\n</body>\n</html>")
	return out.String()
}
