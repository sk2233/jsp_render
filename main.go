/*
@author: sk
@date: 2024/6/10
*/
package main

import (
	"fmt"
	"template_engine/res/test"
)

// https://aosabook.org/en/500L/a-template-engine.html

func main() {
	GenCode("res/test.html")
	fmt.Println(test.Render())
}
