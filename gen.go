/*
@author: sk
@date: 2024/6/10
*/
package main

import (
	"os"
	"strings"
)

func GenCode(file string) {
	bs, err := os.ReadFile(file)
	HandleErr(err)
	content := string(bs)

	packages := make([]string, 0) // 所有的导包语句
	packages = append(packages, DefaultImport)
	public := make([]string, 0) // 所有的函数外定义
	writer := NewFuncWriter()   // 所有函数内容
	start := 0
	for { // 每次循环处理一个指令
		end := strings.Index(content[start:], "<%") // 寻找开头
		if end < 0 {
			break
		}
		writer.WriteOut(content[start : start+end])
		type0 := content[start+end+2] // @ ! = ' ' 4种
		start = start + end + 3
		end = strings.Index(content[start:], "%>") // 寻找结尾
		if end < 0 {
			panic("lost end scope")
		}
		temp := strings.TrimSpace(content[start : start+end])
		switch type0 {
		case '@':
			packages = append(packages, temp)
		case '!':
			public = append(public, temp)
		case '=':
			writer.WriteVar(temp)
		default:
			writer.WriteCode(temp)
		}
		start = start + end + 2
	}
	writer.WriteOut(content[start:])

	start = strings.LastIndexByte(file, '/') + 1
	end := strings.LastIndexByte(file, '.')
	name := file[start:end]
	codes := &strings.Builder{}
	codes.WriteString("package ")
	codes.WriteString(name)
	codes.WriteString("\n\n")
	for _, pkg := range packages {
		codes.WriteString(pkg)
		codes.WriteString("\n")
	}
	codes.WriteString("\n")
	for _, pbc := range public {
		codes.WriteString(pbc)
		codes.WriteString("\n\n")
	}
	codes.WriteString("func Render() string {\n\tout := &strings.Builder{}\n")
	codes.WriteString(writer.String())
	codes.WriteString("\treturn out.String()\n}")

	if !Exists(BasePath + name) {
		err = os.Mkdir(BasePath+name, os.ModePerm)
		HandleErr(err)
	}
	out, err := os.Create(BasePath + name + "/main.go")
	HandleErr(err)
	_, err = out.WriteString(codes.String())
	HandleErr(err)
	err = out.Close()
	HandleErr(err)
}
