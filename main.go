package main

import (
	"hasdsd.cn/op-dl-server/router"
)

func main() {
	//todo:下面这个函数还是不跑
	//schedule.SwagInit()
	r := router.Router()
	r.Run()
}
