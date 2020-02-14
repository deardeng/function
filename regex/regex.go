package main

import (
	"fmt"
	"regexp"
)

const text = `My email is dengxin05@baidu.com
email1 is kk@def.org
email2 is kkk@qq.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}
