/*
@author: sk
@date: 2024/6/10
*/
package main

import (
	"strings"
)

type FuncWriter struct {
	buff *strings.Builder
}

// 	out.WriteString("")
func (f *FuncWriter) WriteOut(val string) {
	if len(val) == 0 {
		return
	}
	f.buff.WriteString("\tout.WriteString(\"")
	val = strings.ReplaceAll(val, "\n", "\\n")  // 标准回车
	val = strings.ReplaceAll(val, "\"", "\\\"") // 标准引号
	f.buff.WriteString(val)
	f.buff.WriteString("\")\n")
}

// 	out.WriteString(fmt.Sprintf("%v", name))
func (f *FuncWriter) WriteVar(val string) {
	f.buff.WriteString("\tout.WriteString(fmt.Sprintf(\"%v\",")
	f.buff.WriteString(val)
	f.buff.WriteString("))\n")
}

func (f *FuncWriter) WriteCode(code string) {
	f.buff.WriteString("\t")
	f.buff.WriteString(code)
	f.buff.WriteString("\n")
}

func (f *FuncWriter) String() string {
	return f.buff.String()
}

func NewFuncWriter() *FuncWriter {
	return &FuncWriter{buff: &strings.Builder{}}
}
