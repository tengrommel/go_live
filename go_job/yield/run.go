package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type resource struct {
	url 				string
	target 			string
	start 			int
	end 				int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		url: "http://localhost:8888/",
		target: "",
		start: 0,
		end: 0,
	}
	r2 := resource{
		url: "http://localhost:8888/list/{$id}.html",
		target: "{$id}",
		start: 1,
		end: 21,
	}
	r3 := resource{
		url: "http://localhost:8888/movie/{$id}.html",
		target: "{$id}",
		start: 1,
		end: 12924,
	}
	res = append(res, r1)
	res = append(res, r2)
	res = append(res, r3)
	return res
}

func buildUrl(res []resource) []string{
	var list []string
	for _, resItem := range res{
		if len(resItem.target)==0{
			list = append(list, resItem.url)
		}else {
			for i :=resItem.start;i<=resItem.end;i++{
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

func main() {
	total := flag.Int("total",100, "how many rows by created")
	filePath := flag.String("filePath", "/home/teng/java_error_in_GOLAND_3228.log", "file path")

	flag.Parse()
	res := ruleResource()
	// 需要构造出真实的网站url集合
	list := buildUrl(res)
	// 按照要求，生成$total 行日志内容，院子这里的集合

	for i:=0;i<=*total;i++{
		logStr := list[i]
		ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}

	fmt.Println(*total, *filePath)
}
