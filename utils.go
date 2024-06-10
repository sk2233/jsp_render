/*
@author: sk
@date: 2024/6/10
*/
package main

import "os"

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil { //os.Stat获取文件信息
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
