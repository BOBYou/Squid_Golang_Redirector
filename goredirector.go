package main

/*
--------参考文档
golang终端输出进度更新的代码	http://golanghome.com/post/607

--------360常用前端公共库CDN服务
http://libs.useso.com/
*/

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"regexp"
	"runtime"
)

func main() {
	//fmt.Println(runtime.GOOS)
	runtime.GOMAXPROCS(runtime.NumCPU())

	running := true

	rwing := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := rwing.ReadLine()
		command := string(data)

		if len(command) == 0 {
			running = false
		} else {

			yn, _ := regexp.MatchString("http://ajax.googleapis.com/ajax/libs/jquery/.+/jquery.min.js.*", command)
			if yn {
				os.Stdout.WriteString(reg_str(command))
				os.Stdout.Sync()
			} else {
				yn, _ := regexp.MatchString("http://fonts.googleapis.com/css.*", command)
				if yn {
					os.Stdout.WriteString(reg_str(command))
					os.Stdout.Sync()
				} else {
					os.Stdout.WriteString("\n")
					os.Stdout.Sync()
				}
			}
		}
	}

}

func reg_str(uri string) string {
	greg := regexp.MustCompile("ajax.googleapis.com/ajax/libs/jquery/")
	return fmt.Sprintf("302:%s\n", greg.ReplaceAllString(uri, "ajax.useso.com/ajax/libs/jquery/"))
}

func reg_str2(uri string) string {
	greg := regexp.MustCompile("fonts.googleapis.com/css")
	return fmt.Sprintf("302:%s\n", greg.ReplaceAllString(uri, "fonts.useso.com/css"))
}
